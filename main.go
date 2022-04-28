package main

import (
	"cockHead/common"
	"cockHead/conf"
	"cockHead/telegram"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.TrustedPlatform = gin.PlatformGoogleAppEngine
	router.TrustedPlatform = "X-CDN-IP"

	router.POST("/"+conf.ApiToken, func(c *gin.Context) {
		var request telegram.Request
		err := c.BindJSON(&request)
		if common.ErrorHandler(err, false) != false {
			return
		}

		switch request.Message.Text {
		case "/start":
			telegram.Start(request.Message.Chat.ID)
		}
	})

	err := router.Run()
	common.ErrorHandler(err, true)
}
