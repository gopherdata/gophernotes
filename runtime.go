package main

import (
	"fmt"
	"strings"
)

// Aim to support an interface as similar to the IPython (canonical Jupyter kernel) as possible.
// See http://ipython.readthedocs.io/en/stable/api/generated/IPython.display.html#IPython.display.display
// for a good overview of the support types. Note: This is missing _repr_markdown_ and _repr_javascript_.

// The following *Renderable interfaces are used by user classes to provide render functions for various
// representations.

type HTMLRenderable interface {
	// RenderAsHTML renders the data as an HTML string. The representaiton should **not** include the
	// `<html>` or `<body>` tags.
	RenderAsHTML() string
}

type MarkdownRenderable interface {
	// RenderAsMarkdown renders the data as a markdown string.
	RenderAsMarkdown() string
}

type SVGRenderable interface {
	// RenderAsSVG renders the data as an svg string (xml) **without the `<svg></svg>` tag.
	RenderAsSVG() string
}

type PNGRenderable interface {
	// RenderAsPNG renders the data as an image in the PNG format.
	RenderAsPNG() []byte
}

type JPEGRenderable interface {
	// RenderAsJPEG renders the data as an image in the JPEG format.
	RenderAsJPEG() []byte
}

type LatexRenderable interface {
	// RenderAsLatex renders the data as a latex string **surrounded in an opening and closing `$`**.
	RenderAsLatex() string
}

type JSONRenderable interface {
	// RenderAsJSON renders the data as a JSON dictionary.
	RenderAsJSON() map[string]interface{}
}

type JavaScriptRenderable interface {
	// RenderAsJavaScript renders the data as a JavaScript string.
	RenderAsJavaScript() string
}

type MIMEBundleRenderable interface {
	// RenderAsMIMEBundle renders the data as a plain MIME bundle. The keys of the map are the MIME type of the
	// data (value) associated with that key. The data will be some JSON serializable object but the structure is
	// determined by what the frontend expects. Some common formats supported by the Jupyter frontend are provided
	// by the various `*Renderable` interfaces.
	RenderAsMIMEBundle() (data, metadata map[string]interface{})
}

type DisplayData struct {
	Data      map[string]interface{}
	Metadata  map[string]interface{}
	Transient map[string]interface{}
}

func (data DisplayData) WithData(mimeType string, rendered interface{}) DisplayData {
	data.Data[mimeType] = rendered
	return data
}

func (data DisplayData) WithMetadata(key string, value interface{}) DisplayData {
	data.Metadata[key] = value
	return data
}

func (data DisplayData) WithTransientData(key string, value interface{}) DisplayData {
	data.Transient[key] = value
	return data
}

const (
	MIMETypeHTML       = "text/html"
	MIMETypeMarkdown   = "text/markdown"
	MIMETypeLatex      = "text/latex"
	MIMETypeSVG        = "image/svg+xml"
	MIMETypePNG        = "image/png"
	MIMETypeJPEG       = "image/jpeg"
	MIMETypeJSON       = "application/json"
	MIMETypeJavaScript = "application/javascript"
)

func Text(data interface{}) DisplayData {
	return DisplayData{
		Data: newTextBundledMIMEData(fmt.Sprint(data)),
	}
}

func HTML(html HTMLRenderable) DisplayData {
	return Text(html).WithData(
		MIMETypeHTML,
		html.RenderAsHTML(),
	)
}

func Markdown(markdown MarkdownRenderable) DisplayData {
	return Text(markdown).WithData(
		MIMETypeMarkdown,
		markdown.RenderAsMarkdown(),
	)
}

func SVG(svg SVGRenderable) DisplayData {
	return Text(svg).WithData(
		MIMETypeSVG,
		svg.RenderAsSVG(),
	)
}

func PNG(png PNGRenderable) DisplayData {
	return Text(png).WithData(
		MIMETypePNG,
		png.RenderAsPNG(), // []byte are encoded as base64 by the marshaller
	)
}

func JPEG(jpeg JPEGRenderable) DisplayData {
	return Text(jpeg).WithData(
		MIMETypeJPEG,
		jpeg.RenderAsJPEG(), // []byte are encoded as base64 by the marshaller
	)
}

func Math(latex LatexRenderable) DisplayData {
	return Text(latex).WithData(
		MIMETypeLatex,
		"$$"+strings.Trim(latex.RenderAsLatex(), "$")+"$$",
	)
}

func Latex(latex LatexRenderable) DisplayData {
	return Text(latex).WithData(
		MIMETypeLatex,
		"$"+strings.Trim(latex.RenderAsLatex(), "$")+"$",
	)
}

func JSON(json JSONRenderable) DisplayData {
	return Text(json).WithData(
		MIMETypeJSON,
		json.RenderAsJSON(),
	)
}

func JavaScript(javascript JavaScriptRenderable) DisplayData {
	return Text(javascript).WithData(
		MIMETypeJavaScript,
		javascript.RenderAsJavaScript(),
	)
}

//TODO the above functions need to handle the metadata

func (receipt *msgReceipt) display(data interface{}) error {
	if dispData, ok := data.(DisplayData); ok {
		return receipt.PublishDisplayData(dispData.Data, dispData.Metadata, dispData.Transient)
	} else if bundleRenderer, ok := data.(MIMEBundleRenderable); ok {
		data, metadata := bundleRenderer.RenderAsMIMEBundle()
		return receipt.PublishDisplayData(data, metadata, map[string]interface{}{})
	} else {
		dispData := Text(data)
		return receipt.PublishDisplayData(dispData.Data, dispData.Metadata, dispData.Transient)
	}
}
