package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
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
	body := url.Values{
		"chat_id": {string(userId)},
		"text":    {message},
	}

	resp, err := http.PostForm(ApiUrl+"/sendMessage", body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
}
