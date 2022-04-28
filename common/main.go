package common

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ErrorHandler(err error, fatal bool) bool {
	if err != nil {
		if fatal == true {
			log.Println(err)
			return true
		}

		panic(err)
		return true
	}
	return false
}

func GetText(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ""
	}

	fContent, err := ioutil.ReadFile("files/mainText.txt")
	if ErrorHandler(err, false) != false {
		return ""
	}
	return string(fContent)
}

func Base64ImageEncoding(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ""
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	default:
		return ""
	}
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)

	return base64Encoding
}
