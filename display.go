package main

import (
	"errors"
	"fmt"
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

func MakeData(mimeType string, data interface{}) Data {
	return Data{
		Data: BundledMIMEData{
			"text/plain": fmt.Sprint(data),
			mimeType:     data,
		},
	}
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
	return MakeData3(mimeType, mimeType, bytes)
}

func HTML(html string) Data {
	return MakeData(MIMETypeHTML, html)
}

func JSON(json map[string]interface{}) Data {
	return MakeData(MIMETypeJSON, json)
}

func JavaScript(javascript string) Data {
	return MakeData(MIMETypeJavaScript, javascript)
}

func JPEG(jpeg []byte) Data {
	return MakeData3(MIMETypeJPEG, "jpeg image", jpeg) // []byte are encoded as base64 by the marshaller
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
	return MakeData3(MIMETypePDF, "pdf document", pdf) // []byte are encoded as base64 by the marshaller
}

func PNG(png []byte) Data {
	return MakeData3(MIMETypePNG, "png image", png) // []byte are encoded as base64 by the marshaller
}

func String(mimeType string, s string) Data {
	return MakeData(mimeType, s)
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
func MIME(data, metadata map[string]interface{}) Data {
	return Data{data, metadata, nil}
}

// prepare imports.Package for interpreted code
var display = imports.Package{
	Binds: map[string]r.Value{
		"Bytes":              r.ValueOf(Bytes),
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
		"String":             r.ValueOf(String),
		"SVG":                r.ValueOf(SVG),
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
