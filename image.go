package main

import (
	"bytes"
	"image"
	"image/png"
	"log"
)

// Image converts an image.Image to DisplayData containing PNG []byte,
// or to DisplayData containing error if the conversion fails
func Image(img image.Image) DisplayData {
	data, err := image0(img)
	if err != nil {
		return DisplayData{
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

// image0 converts an image.Image to DisplayData containing PNG []byte,
// or error if the conversion fails
func image0(img image.Image) (DisplayData, error) {
	bytes, mime, err := encodePng(img)
	if err != nil {
		return DisplayData{}, err
	}
	return DisplayData{
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

// publishImages sends a "display_data" broadcast message for given image.Image.
func publishImage(img image.Image, receipt *msgReceipt) error {
	data, err := image0(img)
	if err != nil {
		return err
	}
	return receipt.PublishDisplayData(data)
}

// publishImagesAndDisplayData sends a "display_data" broadcast message for each
// image.Image and DisplayData found in vals, then replaces it with nil
// to avoid overloading the front-end with huge amounts of output text:
// fmt.Sprint(val) is often very large for an image and other multimedia data.
func publishImagesAndDisplayData(vals []interface{}, receipt *msgReceipt) []interface{} {
	for i, val := range vals {
		switch obj := val.(type) {
		case image.Image:
			err := publishImage(obj, receipt)
			if err != nil {
				log.Printf("Error publishing image.Image: %v\n", err)
			} else {
				vals[i] = nil
			}
		case DisplayData:
			err := receipt.PublishDisplayData(obj)
			if err != nil {
				log.Printf("Error publishing DisplayData: %v\n", err)
			} else {
				vals[i] = nil
			}
		}
	}
	// if all values are nil, return empty slice
	for _, val := range vals {
		if val != nil {
			return vals
		}
	}
	return nil
}
