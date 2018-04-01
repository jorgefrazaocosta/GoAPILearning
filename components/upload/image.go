package upload

import (
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"strings"
)

func Image(data string, to string) (bool, string) {

	log.Println("image upload")

	result, commaPosition := isValid(data)
	if !result {
		return false, "image is not valid"
	}

	format := extractFormat(data, commaPosition)
	imageName := "novaimage." + format

	dst, err := os.Create(imageName)
	if err != nil {
		return false, "erro"
	}
	defer dst.Close()

	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[commaPosition+1:]))
	if _, err = io.Copy(dst, dec); err != nil {
		return false, "erro"
	}

	return true, ""

}

func isValid(data string) (bool, int) {

	i := strings.Index(data, ",")
	if i < 0 {
		return false, 0
	}

	return true, i

}

func extractFormat(data string, i int) string {

	switch strings.TrimSuffix(data[5:i], ";base64") {
	case "image/png":
		return "png"

	case "image/jpeg":
		return "jpeg"
	}

	return ""

}

func extractImage(data string, i int) image.Image {

	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[i+1:]))

	switch strings.TrimSuffix(data[5:i], ";base64") {
	case "image/png":
		pngI, err := png.Decode(dec)

		if err != nil {
			return pngI
		}

	case "image/jpeg":
		jpgI, err := jpeg.Decode(dec)

		if err != nil {
			return jpgI
		}

	}

	return nil

}
