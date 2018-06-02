package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
)

// publishImages sends a "display_data" broadcast message for each image.Image found in vals,
// then replaces it with a placeholder string equal to reflect.TypeOf(val).String()
// to avoid overloading the front-end with huge amounts of output text:
// fmt.Sprint(val) is often very large for an image
func publishImages(vals []interface{}, receipt *msgReceipt) []interface{} {
	for i, val := range vals {
		if img, ok := val.(image.Image); ok {
			err := publishImage(img, receipt)
			if err != nil {
				log.Printf("Error publishing image: %v\n", err)
			} else {
				vals[i] = fmt.Sprintf("%T", val)
			}
		}
	}
	return vals
}

// publishImages sends a "display_data" broadcast message for given image.
func publishImage(img image.Image, receipt *msgReceipt) error {
	bytes, mime, err := encodePng(img)
	if err != nil {
		return err
	}
	data := bundledMIMEData{
		mime: bytes,
	}
	metadata := bundledMIMEData{
		mime: imageMetadata(img),
	}
	return receipt.PublishDisplayData(data, metadata)
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

// imageMetadata returns image size, represented as bundledMIMEData{"width": width, "height": height}
func imageMetadata(img image.Image) bundledMIMEData {
	rect := img.Bounds()
	return bundledMIMEData{
		"width":  rect.Dx(),
		"height": rect.Dy(),
	}
}
