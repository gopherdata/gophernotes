package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strings"

	"github.com/cosmos72/gomacro/base"

	"github.com/cosmos72/gomacro/xreflect"
)

// Support an interface similar - but not identical - to the IPython (canonical Jupyter kernel).
// See http://ipython.readthedocs.io/en/stable/api/generated/IPython.display.html#IPython.display.display
// for a good overview of the support types.

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
	MIMETypeText       = "text/plain"
)

/**
 * general interface, allows libraries to fully specify
 * how their data is displayed by Jupyter.
 * Supports multiple MIME formats.
 *
 * Note that Data defined above is an alias:
 * libraries can implement Renderer without importing gophernotes
 */
type Renderer = interface {
	Render() Data
}

/**
 * simplified interface, allows libraries to specify
 * how their data is displayed by Jupyter.
 * Supports multiple MIME formats.
 *
 * Note that MIMEMap defined above is an alias:
 * libraries can implement SimpleRenderer without importing gophernotes
 */
type SimpleRenderer = interface {
	SimpleRender() MIMEMap
}

/**
 * specialized interfaces, each is dedicated to a specific MIME type.
 *
 * They are type aliases to emphasize that method signatures
 * are the only important thing, not the interface names.
 * Thus libraries can implement them without importing gophernotes
 */
type HTMLer = interface {
	HTML() string
}
type JavaScripter = interface {
	JavaScript() string
}
type JPEGer = interface {
	JPEG() []byte
}
type JSONer = interface {
	JSON() map[string]interface{}
}
type Latexer = interface {
	Latex() string
}
type Markdowner = interface {
	Markdown() string
}
type PNGer = interface {
	PNG() []byte
}
type PDFer = interface {
	PDF() []byte
}
type SVGer = interface {
	SVG() string
}

// injected as placeholder in the interpreter, it's then replaced at runtime
// by a closure that knows how to talk with Jupyter
func stubDisplay(Data) error {
	return errors.New("cannot display: connection with Jupyter not available")
}

// fill kernel.renderer map used to convert interpreted types
// to known rendering interfaces
func (kernel *Kernel) initRenderers() {
	type Pair = struct {
		name string
		typ  xreflect.Type
	}
	var pairs []Pair
	for name, typ := range kernel.display.Types {
		if typ.Kind() == reflect.Interface {
			pairs = append(pairs, Pair{name, typ})
		}
	}
	// for deterministic behaviour, sort alphabetically by name
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].name < pairs[j].name
	})
	kernel.render = make([]xreflect.Type, len(pairs))
	for i, pair := range pairs {
		kernel.render[i] = pair.typ
	}
}

// if vals[] contain a single non-nil value which is auto-renderable,
// convert it to Data and return it.
// otherwise return MakeData("text/plain", fmt.Sprint(vals...))
func (kernel *Kernel) autoRenderResults(vals []interface{}, types []xreflect.Type) Data {
	var nilcount int
	var obj interface{}
	var typ xreflect.Type
	for i, val := range vals {
		if kernel.canAutoRender(val, types[i]) {
			obj = val
			typ = types[i]
		} else if val == nil {
			nilcount++
		}
	}
	if obj != nil && nilcount == len(vals)-1 {
		return kernel.autoRender("", obj, typ)
	}
	if nilcount == len(vals) {
		// if all values are nil, return empty Data
		return Data{}
	}
	return MakeData(MIMETypeText, fmt.Sprint(vals...))
}

// return true if data type should be auto-rendered graphically
func (kernel *Kernel) canAutoRender(data interface{}, typ xreflect.Type) bool {
	switch data.(type) {
	case Data, Renderer, SimpleRenderer, HTMLer, JavaScripter, JPEGer, JSONer,
		Latexer, Markdowner, PNGer, PDFer, SVGer, image.Image:
		return true
	}
	if kernel == nil || typ == nil {
		return false
	}
	// in gomacro, methods of interpreted types are emulated,
	// thus type-asserting them to interface types as done above cannot succeed.
	// Manually check if emulated type "pretends" to implement
	// one of the interfaces above
	for _, xtyp := range kernel.render {
		if typ.Implements(xtyp) {
			return true
		}
	}
	return false
}

// detect and render data types that should be auto-rendered graphically
func (kernel *Kernel) autoRender(mimeType string, arg interface{}, typ xreflect.Type) Data {
	var s string
	var b []byte
	var err error
	var ret Data
	datain := arg
again:
	switch data := datain.(type) {
	case Data:
		ret = data
	case Renderer:
		ret = data.Render()
	case SimpleRenderer:
		ret.Data = data.SimpleRender()
	case HTMLer:
		mimeType = MIMETypeHTML
		s = data.HTML()
	case JavaScripter:
		mimeType = MIMETypeJavaScript
		s = data.JavaScript()
	case JPEGer:
		mimeType = MIMETypeJPEG
		b = data.JPEG()
	case JSONer:
		ret.Data = MIMEMap{MIMETypeJSON: data.JSON()}
	case Latexer:
		mimeType = MIMETypeLatex
		s = data.Latex()
	case Markdowner:
		mimeType = MIMETypeMarkdown
		s = data.Markdown()
	case PNGer:
		mimeType = MIMETypePNG
		b = data.PNG()
	case PDFer:
		mimeType = MIMETypePDF
		b = data.PDF()
	case SVGer:
		mimeType = MIMETypeSVG
		s = data.SVG()
	case image.Image:
		b, mimeType, err = encodePng(data)
		if err == nil {
			ret.Metadata = imageMetadata(data)
		}
	default:
		if kernel != nil && typ != nil {
			// in gomacro, methods of interpreted types are emulated.
			// Thus type-asserting them to interface types as done above cannot succeed.
			// Manually check if emulated type "pretends" to implement one of the above interfaces
			// and, in case, tell the interpreter to convert to them
			for _, xtyp := range kernel.render {
				if typ.Implements(xtyp) {
					fun := kernel.ir.Comp.Converter(typ, xtyp)
					data = base.ValueInterface(fun(reflect.ValueOf(datain)))
					if data != nil {
						s = fmt.Sprint(data)
						datain = data
						// avoid infinite recursion
						kernel = nil
						typ = nil
						goto again
					}
				}
			}
		}
		panic(fmt.Errorf("internal error, autoRender invoked on unexpected type %T", data))
	}
	return fillDefaults(ret, arg, s, b, mimeType, err)
}

func fillDefaults(ret Data, data interface{}, s string, b []byte, mimeType string, err error) Data {
	if err != nil {
		return makeDataErr(err)
	}
	if ret.Data == nil {
		ret.Data = make(MIMEMap)
	}
	// cannot autodetect the mime type of a string
	if len(s) != 0 && len(mimeType) != 0 {
		ret.Data[mimeType] = s
	}
	// ensure plain text is set
	if ret.Data[MIMETypeText] == "" {
		if len(s) == 0 {
			s = fmt.Sprint(data)
		}
		ret.Data[MIMETypeText] = s
	}
	// if []byte is available, use it
	if len(b) != 0 {
		if len(mimeType) == 0 {
			mimeType = http.DetectContentType(b)
		}
		if len(mimeType) != 0 && mimeType != MIMETypeText {
			ret.Data[mimeType] = b
		}
	}
	return ret
}

// do our best to render data graphically
func render(mimeType string, data interface{}) Data {
	var kernel *Kernel // intentionally nil
	if kernel.canAutoRender(data, nil) {
		return kernel.autoRender(mimeType, data, nil)
	}
	var s string
	var b []byte
	var err error
	switch data := data.(type) {
	case string:
		s = data
	case []byte:
		b = data
	case io.Reader:
		b, err = ioutil.ReadAll(data)
	case io.WriterTo:
		var buf bytes.Buffer
		data.WriteTo(&buf)
		b = buf.Bytes()
	default:
		panic(fmt.Errorf("unsupported type, cannot render: %T", data))
	}
	return fillDefaults(Data{}, data, s, b, mimeType, err)
}

func makeDataErr(err error) Data {
	return Data{
		Data: MIMEMap{
			"ename":     "ERROR",
			"evalue":    err.Error(),
			"traceback": nil,
			"status":    "error",
		},
	}
}

func Any(mimeType string, data interface{}) Data {
	return render(mimeType, data)
}

// same as Any("", data), autodetects MIME type
func Auto(data interface{}) Data {
	return render("", data)
}

func MakeData(mimeType string, data interface{}) Data {
	d := Data{
		Data: MIMEMap{
			mimeType: data,
		},
	}
	if mimeType != MIMETypeText {
		d.Data[MIMETypeText] = fmt.Sprint(data)
	}
	return d
}

func MakeData3(mimeType string, plaintext string, data interface{}) Data {
	return Data{
		Data: MIMEMap{
			MIMETypeText: plaintext,
			mimeType:     data,
		},
	}
}

func File(mimeType string, path string) Data {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return Any(mimeType, bytes)
}

func HTML(html string) Data {
	return MakeData(MIMETypeHTML, html)
}

func JavaScript(javascript string) Data {
	return MakeData(MIMETypeJavaScript, javascript)
}

func JPEG(jpeg []byte) Data {
	return MakeData(MIMETypeJPEG, jpeg)
}

func JSON(json map[string]interface{}) Data {
	return MakeData(MIMETypeJSON, json)
}

func Latex(latex string) Data {
	return MakeData3(MIMETypeLatex, latex, "$"+strings.Trim(latex, "$")+"$")
}

func Markdown(markdown string) Data {
	return MakeData(MIMETypeMarkdown, markdown)
}

func Math(latex string) Data {
	return MakeData3(MIMETypeLatex, latex, "$$"+strings.Trim(latex, "$")+"$$")
}

func PDF(pdf []byte) Data {
	return MakeData(MIMETypePDF, pdf)
}

func PNG(png []byte) Data {
	return MakeData(MIMETypePNG, png)
}

func SVG(svg string) Data {
	return MakeData(MIMETypeSVG, svg)
}

// MIME encapsulates the data and metadata into a Data.
// The 'data' map is expected to contain at least one {key,value} pair,
// with value being a string, []byte or some other JSON serializable representation,
// and key equal to the MIME type of such value.
// The exact structure of value is determined by what the frontend expects.
// Some easier-to-use functions for common formats supported by the Jupyter frontend
// are provided by the various functions above.
func MIME(data, metadata MIMEMap) Data {
	return Data{data, metadata, nil}
}
