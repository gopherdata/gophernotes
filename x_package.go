package main

import (
	"image"
	"image/color"
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

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
		"Data":           r.TypeOf((*Data)(nil)).Elem(),
		"HTMLer":         r.TypeOf((*HTMLer)(nil)).Elem(),
		"JavaScripter":   r.TypeOf((*JavaScripter)(nil)).Elem(),
		"Image":          r.TypeOf((*image.Image)(nil)).Elem(),
		"JPEGer":         r.TypeOf((*JPEGer)(nil)).Elem(),
		"JSONer":         r.TypeOf((*JSONer)(nil)).Elem(),
		"Latexer":        r.TypeOf((*Latexer)(nil)).Elem(),
		"Markdowner":     r.TypeOf((*Markdowner)(nil)).Elem(),
		"MIMEMap":        r.TypeOf((*MIMEMap)(nil)).Elem(),
		"PNGer":          r.TypeOf((*PNGer)(nil)).Elem(),
		"PDFer":          r.TypeOf((*PDFer)(nil)).Elem(),
		"Renderer":       r.TypeOf((*Renderer)(nil)).Elem(),
		"SimpleRenderer": r.TypeOf((*SimpleRenderer)(nil)).Elem(),
		"SVGer":          r.TypeOf((*SVGer)(nil)).Elem(),
	}, Proxies: map[string]r.Type{
		// these are needed to allow interpreted types
		// to implement the corresponding interfaces
		"HTMLer":         r.TypeOf((*proxy_HTMLer)(nil)).Elem(),
		"Image":          r.TypeOf((*proxy_image_Image)(nil)).Elem(),
		"JPEGer":         r.TypeOf((*proxy_JPEGer)(nil)).Elem(),
		"JSONer":         r.TypeOf((*proxy_JSONer)(nil)).Elem(),
		"Latexer":        r.TypeOf((*proxy_Latexer)(nil)).Elem(),
		"Markdowner":     r.TypeOf((*proxy_Markdowner)(nil)).Elem(),
		"PNGer":          r.TypeOf((*proxy_PNGer)(nil)).Elem(),
		"PDFer":          r.TypeOf((*proxy_PDFer)(nil)).Elem(),
		"Renderer":       r.TypeOf((*proxy_Renderer)(nil)).Elem(),
		"SimpleRenderer": r.TypeOf((*proxy_SimpleRenderer)(nil)).Elem(),
		"SVGer":          r.TypeOf((*proxy_SVGer)(nil)).Elem(),
	},
}

// --------------- proxy for display.HTMLer ---------------
type proxy_HTMLer struct {
	Object interface{}
	HTML_  func(interface{}) string
}

func (P *proxy_HTMLer) HTML() string {
	return P.HTML_(P.Object)
}

// --------------- proxy for display.JPEGer ---------------
type proxy_JPEGer struct {
	Object interface{}
	JPEG_  func(interface{}) []byte
}

func (P *proxy_JPEGer) JPEG() []byte {
	return P.JPEG_(P.Object)
}

// --------------- proxy for display.JSONer ---------------
type proxy_JSONer struct {
	Object interface{}
	JSON_  func(interface{}) string
}

func (P *proxy_JSONer) JSON() string {
	return P.JSON_(P.Object)
}

// --------------- proxy for display.Latexer ---------------
type proxy_Latexer struct {
	Object interface{}
	Latex_ func(interface{}) string
}

func (P *proxy_Latexer) Latex() string {
	return P.Latex_(P.Object)
}

// --------------- proxy for display.Markdowner ---------------
type proxy_Markdowner struct {
	Object    interface{}
	Markdown_ func(interface{}) string
}

func (P *proxy_Markdowner) Markdown() string {
	return P.Markdown_(P.Object)
}

// --------------- proxy for display.PNGer ---------------
type proxy_PNGer struct {
	Object interface{}
	PNG_   func(interface{}) []byte
}

func (P *proxy_PNGer) PNG() []byte {
	return P.PNG_(P.Object)
}

// --------------- proxy for display.PDFer ---------------
type proxy_PDFer struct {
	Object interface{}
	PDF_   func(interface{}) []byte
}

func (P *proxy_PDFer) PDF() []byte {
	return P.PDF_(P.Object)
}

// --------------- proxy for display.Renderer ---------------
type proxy_Renderer struct {
	Object  interface{}
	Render_ func(interface{}) Data
}

func (P *proxy_Renderer) Render() Data {
	return P.Render_(P.Object)
}

// --------------- proxy for display.SimpleRenderer ---------------
type proxy_SimpleRenderer struct {
	Object        interface{}
	SimpleRender_ func(interface{}) MIMEMap
}

func (P *proxy_SimpleRenderer) SimpleRender() MIMEMap {
	return P.SimpleRender_(P.Object)
}

// --------------- proxy for display.SVGer ---------------
type proxy_SVGer struct {
	Object interface{}
	SVG_   func(interface{}) string
}

func (P *proxy_SVGer) SVG() string {
	return P.SVG_(P.Object)
}

// --------------- proxy for image.Image ---------------
type proxy_image_Image struct {
	Object      interface{}
	At_         func(_proxy_obj_ interface{}, x int, y int) color.Color
	Bounds_     func(interface{}) image.Rectangle
	ColorModel_ func(interface{}) color.Model
}

func (P *proxy_image_Image) At(x int, y int) color.Color {
	return P.At_(P.Object, x, y)
}
func (P *proxy_image_Image) Bounds() image.Rectangle {
	return P.Bounds_(P.Object)
}
func (P *proxy_image_Image) ColorModel() color.Model {
	return P.ColorModel_(P.Object)
}

// --------------------------------------------------------
// allow importing "display" and "github.com/gopherdata/gophernotes" packages
func init() {
	imports.Packages["display"] = display
	imports.Packages["github.com/gopherdata/gophernotes"] = display
}
