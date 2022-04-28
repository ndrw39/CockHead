package telegram

import (
	"cockHead/common"
	"log"
	"strings"
)

func Start(userID int) {
	pathText := "files/mainText.txt"
	mainText := common.GetText(pathText)
	if mainText == "" {
		log.Println("path not exist " + pathText)
		return
	}

	s := strings.Split(mainText, "<br>")
	for k, a := range s {
		if k == 0 {
			SendPhoto(userID, "files/head.jpg")
		}
		SendMessage(userID, a)
	}
}
