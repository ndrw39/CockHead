package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	ApiToken string = "5342559232:AAEB9q73QHm9EaNVoZMa3lS2o0iwL2fg9Po"
	ApiUrl   string = "https://api.telegram.org/bot5342559232:AAEB9q73QHm9EaNVoZMa3lS2o0iwL2fg9Po"
)

func main() {
	router := gin.Default()
	router.TrustedPlatform = gin.PlatformGoogleAppEngine
	router.TrustedPlatform = "X-CDN-IP"

	router.POST("/"+ApiToken, func(c *gin.Context) {
		// If you set TrustedPlatform, ClientIP() will resolve the
		// corresponding header and return IP directly
		sendMessage(834117686, "Hello Dev")
	})

	err := router.Run()
	if err != nil {
		return
	}
}

func sendMessage(userId int, message string) {
	values := map[string]string{
		"chat_id": strconv.Itoa(userId),
		"text":    message,
	}
	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post(ApiUrl+"/sendMessage", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)
}
