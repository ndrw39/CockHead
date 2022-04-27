package main

import (
	"cockHead/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	fmt.Print(conf.API_TOKEN)
	router := gin.Default()
	router.TrustedPlatform = gin.PlatformGoogleAppEngine
	router.TrustedPlatform = "X-CDN-IP"

	router.POST("/"+conf.API_TOKEN, func(c *gin.Context) {
		// If you set TrustedPlatform, ClientIP() will resolve the
		// corresponding header and return IP directly
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
	})

	err := router.Run()
	if err != nil {
		return
	}
}
