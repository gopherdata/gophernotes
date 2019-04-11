package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"net/http"
	r "reflect"
	"strings"

	"github.com/cosmos72/gomacro/imports"
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
 * It only supports a single MIME format.
 *
 * Note that MIMEMap defined above is an alias:
 * libraries can implement SimpleRenderer without importing gophernotes
 */
type SimpleRenderer = interface {
	Render() MIMEMap
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

// return true if data type should be auto-rendered graphically
func canAutoRender(data interface{}) bool {
	switch data.(type) {
	case Data, Renderer, SimpleRenderer, HTMLer, JavaScripter, JPEGer, JSONer,
		Latexer, Markdowner, PNGer, PDFer, SVGer, image.Image:
		return true
	default:
		return false
	}
}

// detect and render data types that should be auto-rendered graphically
func autoRender(mimeType string, data interface{}) Data {
	var s string
	var b []byte
	var err error
	var ret Data
	switch data := data.(type) {
	case Data:
		ret = data
	case Renderer:
		ret = data.Render()
	case SimpleRenderer:
		ret.Data = data.Render()
	case HTMLer:
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
		panic(fmt.Errorf("internal error, autoRender invoked on unexpected type %T", data))
	}
	return fillDefaults(ret, data, s, b, mimeType, err)
}

func fillDefaults(ret Data, data interface{}, s string, b []byte, mimeType string, err error) Data {
	if err != nil {
		return makeDataErr(err)
	}
	if ret.Data == nil {
		ret.Data = make(MIMEMap)
	}
	if ret.Data[MIMETypeText] == "" {
		if len(s) == 0 {
			s = fmt.Sprint(data)
		}
		ret.Data[MIMETypeText] = s
	}
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
	if canAutoRender(data) {
		return autoRender(mimeType, data)
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

// prepare imports.Package for interpreted code
var display = imports.Package{
	Binds: map[string]r.Value{
		"Any":                r.ValueOf(Any),
		"Auto":               r.ValueOf(Auto),
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
		"SVG":                r.ValueOf(SVG),
	},
	Types: map[string]r.Type{
		"Data":    r.TypeOf((*Data)(nil)).Elem(),
		"MIMEMap": r.TypeOf((*MIMEMap)(nil)).Elem(),
	},
}

// allow importing "display" and "github.com/gopherdata/gophernotes" packages
func init() {
	imports.Packages["display"] = display
	imports.Packages["github.com/gopherdata/gophernotes"] = display
}
