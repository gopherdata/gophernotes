package main

import (
	"bytes"
	"fmt"
	"html"
	"image"
	"image/png"
	"reflect"
	"strings"
)

type HTMLer interface {
	HTML() string
}

// Image converts an image.Image to DisplayData containing PNG []byte,
// or to DisplayData containing error if the conversion fails
func Image(img image.Image) Data {
	data, err := image0(img)
	if err != nil {
		return Data{
			Data: BundledMIMEData{
				"ename":     "ERROR",
				"evalue":    err.Error(),
				"traceback": nil,
				"status":    "error",
			},
		}
	}
	return data
}

// Image converts an image.Image to Data containing PNG []byte,
// or error if the conversion fails
func image0(img image.Image) (Data, error) {
	bytes, mime, err := encodePng(img)
	if err != nil {
		return Data{}, err
	}
	return Data{
		Data: BundledMIMEData{
			mime: bytes,
		},
		Metadata: BundledMIMEData{
			mime: imageMetadata(img),
		},
	}, nil
}

// encodePng converts an image.Image to PNG []byte
func encodePng(img image.Image) (data []byte, mime string, err error) {
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), "image/png", nil
}

// imageMetadata returns image size, represented as BundledMIMEData{"width": width, "height": height}
func imageMetadata(img image.Image) BundledMIMEData {
	rect := img.Bounds()
	return BundledMIMEData{
		"width":  rect.Dx(),
		"height": rect.Dy(),
	}
}

// PublishImage sends a "display_data" broadcast message for given image.Image.
func (receipt *msgReceipt) PublishImage(img image.Image) error {
	data, err := image0(img)
	if err != nil {
		return err
	}
	return receipt.PublishDisplayData(data)
}

// makeHTML converts all arguments to HTML formatted versions.
func makeHTML(vals []interface{}) Data {
	out := make([]string, len(vals), len(vals))
	for i, item := range vals {
		if item == nil || reflect.ValueOf(item).IsNil() {
			continue
		}
		if v, ok := item.(HTMLer); ok {
			out[i] = v.HTML()
		} else if v, ok := item.(Data); ok {
			if t, ok := v.Data["text/html"]; ok {
				out[i] = fmt.Sprint(t)
			} else if t, ok := v.Data["text/plain"]; ok {
				out[i] = html.EscapeString(fmt.Sprint(t))
			}
		} else {
			out[i] = fmt.Sprintf("<pre>%s</pre>", html.EscapeString(fmt.Sprint(item)))
		}
	}
	return Data{
		Data: BundledMIMEData{
			"text/html": strings.Join(out, " "),
		},
	}
}

// if vals[] contain a single non-nil value which is an image.Image,
// convert it to Data and return it.
// if instead the single non-nil value is a Data, return it.
// otherwise return MakeData("text/plain", fmt.Sprint(vals...))
func renderResults(vals []interface{}) Data {
	hasHTML := false
	var nilcount int
	var obj interface{}
	for _, val := range vals {
		switch val.(type) {
		case image.Image, Data:
			obj = val
		case nil:
			nilcount++
		}
		if _, ok := val.(HTMLer); ok {
			hasHTML = true
		}
	}
	if obj != nil && nilcount == len(vals)-1 {
		switch val := obj.(type) {
		case image.Image:
			data, err := image0(val)
			if err == nil {
				return data
			}
		case Data:
			return val
		}
	}
	if nilcount == len(vals) {
		// if all values are nil, return empty Data
		return Data{}
	}
	out := MakeData("text/plain", fmt.Sprint(vals...))
	if hasHTML {
		h := makeHTML(vals)
		out.Data["text/html"] = h.Data["text/html"]
	}
	return out
}
