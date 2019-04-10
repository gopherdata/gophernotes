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
)

/**
 * general interface, allows libraries to fully control
 * how their data is displayed by gophernotes.
 * Supports multiple MIME formats.
 *
 * Note that Data is an alias:
 * type Data = struct { Data, Metadata, Transient map[string]interface{} }
 * thus libraries can implement Renderer without importing gophernotes
 */
type Renderer = interface {
	Render() Data
}

/**
 * simplified interface, allows libraries to control
 * how their data is displayed by gophernotes.
 * It only supports a single MIME format.
 *
 * Note that MIMEMap is an alias
 * type MIMEMap = map[string]interface{}
 * thus libraries can implement SimpleRenderer without importing gophernotes
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
type Javascripter = interface {
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
	if img, ok := data.(image.Image); ok {
		return Image(img)
	}
	b, s := read(data)
	if len(mimeType) == 0 {
		mimeType = http.DetectContentType(b)
	}
	d := Data{
		Data: MIMEMap{
			"text/plain": s,
		},
	}
	if mimeType != "text/plain" {
		d.Data[mimeType] = b
	}
	return d
}

// same as Any("", data), autodetects MIME type
func Auto(data interface{}) Data {
	return Any("", data)
}

func MakeData(mimeType string, data interface{}) Data {
	d := Data{
		Data: MIMEMap{
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
		Data: MIMEMap{
			"text/plain": plaintext,
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
