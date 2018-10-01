package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	r "reflect"
	"strings"

	"github.com/cosmos72/gomacro/imports"
)

// Support an interface similar - but not identical - to the IPython (canonical Jupyter kernel).
// See http://ipython.readthedocs.io/en/stable/api/generated/IPython.display.html#IPython.display.display
// for a good overview of the support types. Note: This is missing _repr_markdown_ and _repr_javascript_.

const (
	MIMETypeHTML       = "text/html"
	MIMETypeJavaScript = "application/javascript"
	MIMETypeJPEG       = "image/jpeg"
	MIMETypeJSON       = "application/json"
	MIMETypeLatex      = "text/latex"
	MIMETypeMarkdown   = "text/markdown"
	MIMETypePNG        = "image/png"
	MIMETypePDF        = "application/pdf"
	MIMETypeSVG        = "image/svg+xml"
)

// injected as placeholder in the interpreter, it's then replaced at runtime
// by a closure that knows how to talk with Jupyter
func stubDisplay(Data) error {
	return errors.New("cannot display: connection with Jupyter not available")
}

// TODO handle the metadata

func read(data interface{}) ([]byte, string) {
	var b []byte
	var s string
	switch x := data.(type) {
	case string:
		s = x
	case []byte:
		b = x
	case io.Reader:
		bb, err := ioutil.ReadAll(x)
		if err != nil {
			panic(err)
		}
		b = bb
	case io.WriterTo:
		var buf bytes.Buffer
		x.WriteTo(&buf)
		b = buf.Bytes()
	default:
		panic(errors.New(fmt.Sprintf("unsupported type, cannot display: expecting string, []byte, io.Reader or io.WriterTo, found %T", data)))
	}
	if len(s) == 0 {
		s = fmt.Sprint(data)
	}
	return b, s
}

func Any(mimeType string, data interface{}) Data {
	b, s := read(data)
	if len(mimeType) == 0 {
		mimeType = http.DetectContentType(b)
	}
	d := Data{
		Data: BundledMIMEData{
			"text/plain": s,
		},
	}
	if mimeType != "text/plain" {
		d.Data[mimeType] = b
	}
	return d
}

func MakeData(mimeType string, data interface{}) Data {
	d := Data{
		Data: BundledMIMEData{
			mimeType: data,
		},
	}
	if mimeType != "text/plain" {
		d.Data["text/plain"] = fmt.Sprint(data)
	}
	return d
}

func MakeData3(mimeType string, plaintext string, data interface{}) Data {
	return Data{
		Data: BundledMIMEData{
			"text/plain": plaintext,
			mimeType:     data,
		},
	}
}

func Bytes(mimeType string, bytes []byte) Data {
	if len(mimeType) == 0 {
		mimeType = http.DetectContentType(bytes)
	}
	return MakeData3(mimeType, mimeType, bytes)
}

func File(mimeType string, path string) Data {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return Bytes(mimeType, bytes)
}

func HTML(html string) Data {
	return String(MIMETypeHTML, html)
}

func JSON(json map[string]interface{}) Data {
	return MakeData(MIMETypeJSON, json)
}

func JavaScript(javascript string) Data {
	return String(MIMETypeJavaScript, javascript)
}

func JPEG(jpeg []byte) Data {
	return Bytes(MIMETypeJPEG, jpeg)
}

func Latex(latex string) Data {
	return MakeData3(MIMETypeLatex, latex, "$"+strings.Trim(latex, "$")+"$")
}

func Markdown(markdown string) Data {
	return String(MIMETypeMarkdown, markdown)
}

func Math(latex string) Data {
	return MakeData3(MIMETypeLatex, latex, "$$"+strings.Trim(latex, "$")+"$$")
}

func PDF(pdf []byte) Data {
	return Bytes(MIMETypePDF, pdf)
}

func PNG(png []byte) Data {
	return Bytes(MIMETypePNG, png)
}

func Reader(mimeType string, r io.Reader) Data {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return Bytes(mimeType, b)
}

func String(mimeType string, s string) Data {
	if len(mimeType) == 0 {
		mimeType = http.DetectContentType([]byte(s))
	}
	return MakeData3(mimeType, s, s)
}

func SVG(svg string) Data {
	return String(MIMETypeSVG, svg)
}

func WriterTo(mimeType string, to io.WriterTo) Data {
	var buf bytes.Buffer
	_, err := to.WriteTo(&buf)
	if err != nil {
		panic(err)
	}
	return Bytes(mimeType, buf.Bytes())
}

// MIME encapsulates the data and metadata into a Data.
// The 'data' map is expected to contain at least one {key,value} pair,
// with value being a string, []byte or some other JSON serializable representation,
// and key equal to the MIME type of such value.
// The exact structure of value is determined by what the frontend expects.
// Some easier-to-use functions for common formats supported by the Jupyter frontend
// are provided by the various functions above.
func MIME(data, metadata map[string]interface{}) Data {
	return Data{data, metadata, nil}
}

// prepare imports.Package for interpreted code
var display = imports.Package{
	Binds: map[string]r.Value{
		"Any":                r.ValueOf(Any),
		"Bytes":              r.ValueOf(Bytes),
		"File":               r.ValueOf(File),
		"HTML":               r.ValueOf(HTML),
		"Image":              r.ValueOf(Image),
		"JPEG":               r.ValueOf(JPEG),
		"JSON":               r.ValueOf(JSON),
		"JavaScript":         r.ValueOf(JavaScript),
		"Latex":              r.ValueOf(Latex),
		"MakeData":           r.ValueOf(MakeData),
		"MakeData3":          r.ValueOf(MakeData3),
		"Markdown":           r.ValueOf(Markdown),
		"Math":               r.ValueOf(Math),
		"MIME":               r.ValueOf(MIME),
		"MIMETypeHTML":       r.ValueOf(MIMETypeHTML),
		"MIMETypeJavaScript": r.ValueOf(MIMETypeJavaScript),
		"MIMETypeJPEG":       r.ValueOf(MIMETypeJPEG),
		"MIMETypeJSON":       r.ValueOf(MIMETypeJSON),
		"MIMETypeLatex":      r.ValueOf(MIMETypeLatex),
		"MIMETypeMarkdown":   r.ValueOf(MIMETypeMarkdown),
		"MIMETypePDF":        r.ValueOf(MIMETypePDF),
		"MIMETypePNG":        r.ValueOf(MIMETypePNG),
		"MIMETypeSVG":        r.ValueOf(MIMETypeSVG),
		"PDF":                r.ValueOf(PDF),
		"PNG":                r.ValueOf(PNG),
		"Reader":             r.ValueOf(Reader),
		"String":             r.ValueOf(String),
		"SVG":                r.ValueOf(SVG),
		"WriterTo":           r.ValueOf(WriterTo),
	},
	Types: map[string]r.Type{
		"BundledMIMEData": r.TypeOf((*BundledMIMEData)(nil)).Elem(),
		"Data":            r.TypeOf((*Data)(nil)).Elem(),
	},
}

// allow importing "display" and "github.com/gopherdata/gophernotes" packages
func init() {
	imports.Packages["display"] = display
	imports.Packages["github.com/gopherdata/gophernotes"] = display
}
