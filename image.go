package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
)

// Image converts an image.Image to DisplayData containing PNG []byte,
// or to DisplayData containing error if the conversion fails
func Image(img image.Image) Data {
	bytes, mimeType, err := encodePng(img)
	if err != nil {
		return makeDataErr(err)
	}
	return Data{
		Data: MIMEMap{
			mimeType: bytes,
		},
		Metadata: MIMEMap{
			mimeType: imageMetadata(img),
		},
	}
}

// encodePng converts an image.Image to PNG []byte
func encodePng(img image.Image) (data []byte, mimeType string, err error) {
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), MIMETypePNG, nil
}

// imageMetadata returns image size, represented as MIMEMap{"width": width, "height": height}
func imageMetadata(img image.Image) MIMEMap {
	rect := img.Bounds()
	return MIMEMap{
		"width":  rect.Dx(),
		"height": rect.Dy(),
	}
}

// if vals[] contain a single non-nil value which is auto-renderable,
// convert it to Data and return it.
// otherwise return MakeData("text/plain", fmt.Sprint(vals...))
func autoRenderResults(vals []interface{}) Data {
	var nilcount int
	var obj interface{}
	for _, val := range vals {
		if canAutoRender(val) {
			obj = val
		} else if val == nil {
			nilcount++
		}
	}
	if obj != nil && nilcount == len(vals)-1 {
		return autoRender("", obj)
	}
	if nilcount == len(vals) {
		// if all values are nil, return empty Data
		return Data{}
	}
	return MakeData(MIMETypeText, fmt.Sprint(vals...))
}
