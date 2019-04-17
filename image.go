package main

import (
	"bytes"
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
