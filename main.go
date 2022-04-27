package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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
		var requestBody error
		err := c.BindJSON(&requestBody)
		errorHandler(err, false)
		c.IndentedJSON(http.StatusOK, requestBody)
		str := fmt.Sprintf("%v", requestBody)
		sendMessage(834117686, str)
	})
	err := router.Run()
	errorHandler(err, true)
}

func sendMessage(userId int, message string) {
	values := map[string]string{
		"chat_id": strconv.Itoa(userId),
		"text":    message,
	}
	jsonValue, _ := json.Marshal(values)
	buffer := bytes.NewBuffer(jsonValue)
	resp, err := http.Post(ApiUrl+"/sendMessage", "application/json", buffer)

	errorHandler(err, false)
	errorHandler(resp.Body.Close(), true)
}

func errorHandler(err error, fatal bool) {
	if err != nil {
		if fatal == true {
			log.Println(err)
		} else {
			log.Fatalln(err)
		}
	}
}
