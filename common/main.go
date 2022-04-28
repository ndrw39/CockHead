package common

import (
	"io/ioutil"
	"log"
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
