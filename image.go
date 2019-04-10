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
	data, err := image0(img)
	if err != nil {
		return Data{
			Data: MIMEMap{
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
		Data: MIMEMap{
			mime: bytes,
		},
		Metadata: MIMEMap{
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
func imageMetadata(img image.Image) MIMEMap {
	rect := img.Bounds()
	return MIMEMap{
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

// if vals[] contain a single non-nil value which is an image.Image,
// convert it to Data and return it.
// if instead the single non-nil value is a Data, return it.
// otherwise return MakeData("text/plain", fmt.Sprint(vals...))
func renderResults(vals []interface{}) Data {
	var nilcount int
	var obj interface{}
	for _, val := range vals {
		switch val.(type) {
		case image.Image, Data:
			obj = val
		case nil:
			nilcount++
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
	return MakeData("text/plain", fmt.Sprint(vals...))
}
