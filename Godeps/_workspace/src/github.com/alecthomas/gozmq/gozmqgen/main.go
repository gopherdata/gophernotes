package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
)

type Args struct {
	zversion      string
	pages         string
	comment_width uint
	templsource   string
	headersource  string
}

var args = Args{
	zversion:      "2.1,2.2,3.2,4.0",
	pages:         "getsockopt,setsockopt",
	comment_width: 72,
	templsource:   "./gozmqgen/template.txt",
	headersource:  "./gozmqgen/header.txt",
}

func main() {
	flag.StringVar(&args.zversion, "zversion", args.zversion, "version of zmq")
	flag.StringVar(&args.pages, "pages", args.pages, "comma-delimited man page names")
	flag.UintVar(&args.comment_width, "comment-width", args.comment_width, "maximum width of comment text")
	flag.StringVar(&args.templsource, "template", args.templsource, "Name of template file or - to read STDIN.")
	flag.StringVar(&args.headersource, "header", args.headersource, "Name of header file.")
	flag.Parse()

	if len(args.templsource) == 0 {
		panic("'template' is required.")
	}

	data := map[string]interface{}{
		"copyright": string(mustRead(args.headersource)),
	}
	for _, version := range strings.Split(args.zversion, ",") {
		data["build"] = buildtags[version]
		data["version"] = version
		for _, page := range strings.Split(args.pages, ",") {
			manual, err := LoadManual(version, "zmq_"+page)
			if err != nil {
				panic(err.Error())
			}
			defer manual.Close()
			cite := "http://api.zeromq.org/" + version + ":zmq-" + page
			var options []map[string]interface{}
			optlist, err := ParseOptions(version, manual)
			if err != nil {
				panic(err.Error())
			}
			for _, o := range optlist {
				o.SetCitation(cite + "#" + o.anchor)
				if !ignore[""][o.shortname] && !ignore[version][o.shortname] {
					options = append(options, o.Pod())
				}
			}
			data[page] = map[string]interface{}{
				"cite":    cite,
				"options": options,
			}
		}
		raw_template := mustRead(args.templsource)
		t, err := template.New("main").Parse(string(raw_template))
		if err != nil {
			panic(err.Error())
		}
		out, err := os.Create("zmqgen_" + strings.Replace(version, ".", "_", -1) + ".go")
		if err != nil {
			panic(err.Error())
		} else if err = t.Execute(out, data); err != nil {
			panic(err.Error())
		}
	}
}

var (
	gotypes = map[string]map[string]string{
		"binary data": map[string]string{
			"": "string",
		},
		"binary data or Z85 text string": map[string]string{
			"": "string",
		},
		"character string": map[string]string{
			"": "string",
		},
		"int": map[string]string{
			"":             "int",
			"boolean":      "bool",
			"milliseconds": "time.Duration",
		},
		"int on POSIX systems, SOCKET on Windows": map[string]string{
			"": "int",
		},
		"int64_t": map[string]string{
			"":             "int64",
			"boolean":      "bool",
			"milliseconds": "time.Duration",
		},
		"NULL-terminated character string": map[string]string{
			"": "string",
		},
		"uint32_t": map[string]string{
			"": "uint32",
		},
		"uint64_t": map[string]string{
			"":        "uint64",
			"boolean": "bool",
		},
	}

	ztypes = map[string]map[string]string{
		"binary data":                    map[string]string{"": "String"},
		"binary data or Z85 text string": map[string]string{"": "String"},
		"character string":               map[string]string{"": "String"},
		"int":                            map[string]string{"": "Int"},
		"int on POSIX systems, SOCKET on Windows": map[string]string{"": "Int"},
		"int64_t":                          map[string]string{"": "Int64"},
		"NULL-terminated character string": map[string]string{"": "String"},
		"uint32_t":                         map[string]string{"": "UInt32"},
		"uint64_t":                         map[string]string{"": "UInt64"},
	}

	lowtypes = map[string]string{
		"int32_t":  "int32",
		"int64_t":  "int64",
		"uint32_t": "uint32",
		"uint64_t": "uint64",
	}

	replacements = map[string]string{
		"buf":      "Buf",
		"Hwm":      "HWM",
		"hwm":      "HWM",
		"Ipv4only": "IPv4Only",
		"more":     "More",
		"msg":      "Msg",
		"pub":      "PUB",
		"Router":   "ROUTER",
		"size":     "Size",
		"Tcp":      "TCP",
		"timeo":    "Timeout",
	}

	cachedir = path.Join(".cache", "codegen")

	// version : shortname : C type
	fixedtypes = map[string]map[string]string{
		"": map[string]string{
			"EVENTS":       "uint64_t",
			"FD":           "int",
			"RATE":         "int64_t",
			"RCVBUF":       "uint64_t",
			"RECOVERY_IVL": "int64_t",
			"SNDBUF":       "uint64_t",
			"TYPE":         "uint64_t",
		},
		"2.1": map[string]string{
			"RCVMORE": "uint64_t",
		},
		"2.2": map[string]string{
			"RCVMORE": "uint64_t",
		},
	}

	fixedgotypes = map[string]string{
		"TYPE": "SocketType",
	}

	// shortname : unit
	fixedunits = map[string]string{
		"DELAY_ATTACH_ON_CONNECT": "boolean",
		"ROUTER_MANDATORY":        "boolean",
		"XPUB_VERBOSE":            "boolean",
	}

	// version : shortname
	ignore = map[string]map[string]bool{
		"": map[string]bool{
			"FD":             true,
			"LAST_ENDPOINT":  true,
			"MULTICAST_HOPS": true,
		},
		"2.1": map[string]bool{
			"RECOVERY_IVL": true,
		},
		"2.2": map[string]bool{
			"RECOVERY_IVL": true,
		},
		"3.2": map[string]bool{},
	}

	// shortname : shortname
	rename = map[string]string{
		"RECOVERY_IVL_MSEC": "RECOVERY_IVL",
	}

	buildtags = map[string]string{
		"2.1": "zmq_2_1",
		"2.2": "!zmq_2_1,!zmq_3_x,!zmq_4_x",
		"3.2": "zmq_3_x",
		"4.0": "zmq_4_x",
	}
)

func fix(s string) string {
	for key, value := range replacements {
		s = strings.Replace(s, key, value, -1)
	}
	return s
}

func LoadManual(version string, funcname string) (io.ReadCloser, error) {
	pagename := version + ":" + funcname
	pagename = strings.Replace(pagename, ".", "-", -1)
	pagename = strings.Replace(pagename, "_", "-", -1)
	cachepath := path.Join(cachedir, pagename)
	cachefile, err := os.Open(cachepath)
	if err == nil {
		return cachefile, nil
	}
	os.MkdirAll(cachedir, 0755)
	url := "http://api.zeromq.org/" + pagename
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s -> %s", url, resp.Status)
	}
	if cachefile, err = os.Create(cachepath); err != nil {
		return nil, err
	}
	if _, err := io.Copy(cachefile, resp.Body); err != nil {
		return nil, err
	}
	return os.Open(cachepath)
}

type Option struct {
	typ       string   // type as listed in source documentation
	unit      string   // unit as listed source documentation
	fullname  string   // e.g. "ZMQ_RCVMORE"
	shortname string   // e.g. "RCVMORE"
	desc      []string // one or more paragraphs
	cite      string   // URL to source documentation
	ctype     string   // C type e.g. "uint64_t"
	gotype    string   // Go type e.g. "SocketType"
	ztype     string   // option method suffix e.g. "Int64"
	lowtype   string   // go equivalent of C type e.g. "uint64"
	anchor    string   // id in source documentation
	duration  bool     // true if gotype=="time.Duration"
	gounit    string   // e.g. "time.Millisecond"
	cast      bool     // true if gotype needs casting
}

func NewOption(fullname string) *Option {
	return &Option{
		fullname:  fullname,
		shortname: fullname[4:],
	}
}

func (o *Option) Name() (name string) {
	shortname, ok := rename[o.shortname]
	if !ok {
		shortname = o.shortname
	}
	for _, part := range strings.Split(shortname, "_") {
		name = name + string(part[0]) + strings.ToLower(part[1:])
	}
	return fix(name)
}

func (o *Option) AppendDescription(line string) {
	line = strings.TrimSpace(line)
	if len(line) > 0 {
		o.desc = append(o.desc, line)
	}
}

func (o *Option) SetCitation(cite string) {
	o.cite = cite
}

func (o *Option) Comment() (comment string) {
	desc := o.desc[1:]
	if len(o.cite) > 0 {
		desc = append([]string{o.cite, ""}, desc...)
	}
	for _, line := range desc {
		if len(line) > 0 && line[0] == ' ' {
			comment = comment + strings.TrimRight(line, "\n") + "\n"
		} else {
			// TODO: wrap line to width chars
			//wrapped = textwrap.wrap(line, width)
			//"\n".join(wrapped) + "\n%s\n" % wrapped
			comment = comment + line
		}
	}
	return
}

func (o *Option) Summary() string {
	return o.fullname + ": " + o.desc[0] + "."
}

func (o *Option) Pod() map[string]interface{} {
	return map[string]interface{}{
		"fullname":    o.fullname,
		"shortname":   o.shortname,
		"nicename":    o.Name(),
		"summary":     o.Summary(),
		"description": strings.Split(o.Comment(), "\n"),
		"ctype":       o.typ,
		"gotype":      o.gotype,
		"ztype":       o.ztype,
		"lowtype":     o.lowtype,
		"anchor":      o.anchor,
		"duration":    o.duration,
		"citation":    o.cite,
		"boolean":     o.gotype == "bool",
		"gounit":      o.gounit,
		"cast":        o.cast,
	}
}

func (o *Option) String() string { return o.Name() }

type OptionsBuilder struct {
	options []*Option
	version string
}

func (b *OptionsBuilder) Add(name string, info string) bool {
	if !strings.HasPrefix(name, "ZMQ_") {
		return false
	}
	option := NewOption(name)
	b.options = append(b.options, option)
	option.AppendDescription(info)
	return true
}

func (b *OptionsBuilder) Describe(info string) {
	if len(b.options) > 0 {
		b.options[len(b.options)-1].AppendDescription(info)
	}
}

func (b *OptionsBuilder) SetAnchor(anchor string) {
	if len(b.options) > 0 {
		b.options[len(b.options)-1].anchor = anchor
	}
}

func (b *OptionsBuilder) SetProperty(name string, value string) {
	if len(b.options) > 0 {
		option := b.options[len(b.options)-1]
		name = strings.TrimSpace(name)
		switch name {
		case "Option value type":
			option.typ = value
			break
		case "Option value unit":
			option.unit = value
			break
		case "Option value size":
			option.unit = "Z85"
			break
		}
		if len(option.typ) > 0 && len(option.unit) > 0 {
			if val, ok := fixedtypes[""][option.shortname]; ok {
				option.typ = val
			} else if val, ok := fixedtypes[b.version][option.shortname]; ok {
				option.typ = val
			}
			if val, ok := lowtypes[option.typ]; ok {
				option.lowtype = val
			} else {
				option.lowtype = option.typ
			}
			if val, ok := fixedunits[option.shortname]; ok {
				option.unit = val
			}
			gomap := gotypes[option.typ]
			if val, ok := gomap[option.unit]; ok {
				option.gotype = val
			} else {
				option.gotype = gomap[""]
			}
			if val, ok := fixedgotypes[option.shortname]; ok {
				option.gotype = val
				option.cast = true
			}
			zmap := ztypes[option.typ]
			if val, ok := zmap[option.unit]; ok {
				option.ztype = val
			} else {
				option.ztype = zmap[""]
			}
			if option.gotype == "time.Duration" {
				option.duration = true
				switch option.unit {
				case "milliseconds":
					option.gounit = "time.Millisecond"
				case "seconds":
					option.gounit = "time.Second"
				}
			}
		}
		option.AppendDescription(fmt.Sprintf(" %-25s %s\n", name, value))
	}
}

func ParseOptions(version string, r io.Reader) ([]*Option, error) {
	d := xml.NewDecoder(r)
	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity
	b := &OptionsBuilder{version: version}
	var state, text, text2 string
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		} else {
			switch token := t.(type) {
			case xml.StartElement:
				switch token.Name.Local {
				case "h3":
					state = "title"
					text = ""
					text2 = "unknown"
					for _, attr := range token.Attr {
						switch attr.Name.Local {
						case "id":
							text2 = attr.Value
						}
					}
				case "table":
					switch state {
					case "describing":
						for _, attr := range token.Attr {
							switch attr.Name.Local {
							case "class":
								if attr.Value == "wiki-content-table" {
									b.Describe(text)
									text = ""
									state = "properties"
								}
							}
						}
					}
				case "td":
					switch state {
					case "properties":
						state = "property-name"
					case "property-name":
						state = "property-value"
					}
				}
			case xml.EndElement:
				switch state {
				case "title":
					switch token.Name.Local {
					case "h3":
						parts := strings.SplitN(text, ": ", 2)
						if b.Add(parts[0], parts[1]) {
							state = "describing"
							b.SetAnchor(text2)
						} else {
							state = ""
						}
						text = ""
						text2 = ""
					}
				case "describing":
					switch token.Name.Local {
					case "p", "li":
						b.Describe(text)
						text = ""
					}
				case "properties":
					switch token.Name.Local {
					case "table":
						state = ""
					}
				case "property-value":
					switch token.Name.Local {
					case "td":
						b.SetProperty(text, text2)
						state = "properties"
						text = ""
						text2 = ""
					}
				}
			case xml.CharData:
				switch state {
				case "title":
					text += string(token)
				case "describing":
					text += string(token)
				case "property-name":
					text += string(token)
				case "property-value":
					text2 += string(token)
				}
			}
		}
	}
	return b.options, nil
}

func mustRead(name string) []byte {
	var raw []byte
	var err error
	if name == "-" {
		raw, err = ioutil.ReadAll(os.Stdin)
	} else {
		raw, err = ioutil.ReadFile(name)
	}
	if err != nil {
		panic(err.Error())
	}
	return raw
}
