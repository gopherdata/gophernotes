package main

import (
	"errors"
	"fmt"
	"image"
	r "reflect"
	"strings"

	"github.com/cosmos72/gomacro/imports"
)

// Support an interface similar - but not identical - to the IPython (canonical Jupyter kernel).
// See http://ipython.readthedocs.io/en/stable/api/generated/IPython.display.html#IPython.display.display
// for a good overview of the support types. Note: This is missing _repr_markdown_ and _repr_javascript_.

var globalReceipt *msgReceipt // ugly global variable. any alternative?

const (
	mimeTypeHTML       = "text/html"
	mimeTypeMarkdown   = "text/markdown"
	mimeTypeLatex      = "text/latex"
	mimeTypeSVG        = "image/svg+xml"
	mimeTypePNG        = "image/png"
	mimeTypeJPEG       = "image/jpeg"
	mimeTypeJSON       = "application/json"
	mimeTypeJavaScript = "application/javascript"
)

// TODO handle the metadata

func render2(mimeType string, data interface{}) error {
	return render3(mimeType, fmt.Sprint(data), data)
}

func render3(mimeType string, text string, data interface{}) error {
	receipt := globalReceipt
	if receipt == nil {
		return errors.New("msgReceipt is nil, cannot send display_data message")
	}
	return receipt.PublishDisplayData(
		bundledMIMEData{
			"text/plain": text,
			mimeType:     data,
		}, make(bundledMIMEData))
}

func HTML(html string) error {
	return render2(mimeTypeHTML, html)
}

func Markdown(markdown string) error {
	return render2(mimeTypeMarkdown, markdown)
}

func SVG(svg string) error {
	return render2(mimeTypeSVG, svg)
}

func PNG(png []byte) error {
	return render3(mimeTypePNG, "{png-image}", png) // []byte are encoded as base64 by the marshaller
}

func JPEG(jpeg []byte) error {
	return render3(mimeTypeJPEG, "{jpeg-image}", jpeg) // []byte are encoded as base64 by the marshaller
}

func Image(img image.Image) error {
	return publishImage(img, globalReceipt)
}

func Math(latex string) error {
	return render3(mimeTypeLatex, latex, "$$"+strings.Trim(latex, "$")+"$$")
}

func Latex(latex string) error {
	return render3(mimeTypeLatex, latex, "$"+strings.Trim(latex, "$")+"$")
}

func JSON(json map[string]interface{}) error {
	return render2(mimeTypeJSON, json)
}

func JavaScript(javascript string) error {
	return render2(mimeTypeJavaScript, javascript)
}

// MIME renders the data as a plain MIME bundle. The keys of the map are the MIME type of the
// data (value) associated with that key. The data will be some JSON serializable object but the structure is
// determined by what the frontend expects. Some easier-to-use formats supported by the Jupyter frontend
// are provided by the various functions above.
func MIME(data, metadata map[string]interface{}) error {
	return globalReceipt.PublishDisplayData(data, metadata)
}

// prepare imports.Package for interpreted code
var display = imports.Package{
	Binds: map[string]r.Value{
		"HTML":       r.ValueOf(HTML),
		"JPEG":       r.ValueOf(JPEG),
		"JSON":       r.ValueOf(JSON),
		"JavaScript": r.ValueOf(JavaScript),
		"Latex":      r.ValueOf(Latex),
		"Markdown":   r.ValueOf(Markdown),
		"Math":       r.ValueOf(Math),
		"MIME":       r.ValueOf(MIME),
		"PNG":        r.ValueOf(PNG),
		"SVG":        r.ValueOf(SVG),
	},
}

// allow import of "display" and "github.com/gopherdata/gophernotes" packages
func init() {
	imports.Packages["display"] = display
	imports.Packages["github.com/gopherdata/gophernotes"] = display
}
