package main

import (
	"bytes"
	"image"
	"image/png"
	"log"
)

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

// if vals[] contain a single non-nil value which is an image.Image
// or a display.Data, publishImageOrDisplayData sends a "display_data"
// broadcast message for such value, then returns nil to avoid overloading
// the front-end with huge amounts of output text:
// fmt.Sprint(val) is often very large for an image and other multimedia data.
func (receipt *msgReceipt) PublishImageOrDisplayData(vals []interface{}) []interface{} {
	var nilcount int
	var data interface{}
	for _, val := range vals {
		switch obj := val.(type) {
		case image.Image, Data:
			data = obj
		case nil:
			nilcount++
		}
	}
	if data != nil && nilcount == len(vals)-1 {
		switch obj := data.(type) {
		case image.Image:
			err := receipt.PublishImage(obj)
			if err != nil {
				log.Printf("Error publishing image.Image: %v\n", err)
			} else {
				nilcount++
			}
		case Data:
			err := receipt.PublishDisplayData(obj)
			if err != nil {
				log.Printf("Error publishing Data: %v\n", err)
			} else {
				nilcount++
			}
		}
	}
	if nilcount == len(vals) {
		// if all values are nil, return empty slice
		return nil
	}
	return vals
}
