package main

import (
	"cockHead/conf"
	"fmt"
	//"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	fmt.Print(conf.API_TOKEN)
	//router := gin.Default()
	//router.TrustedPlatform = gin.PlatformGoogleAppEngine
	//// Or set your own trusted request header for another trusted proxy service
	//// Don't set it to any suspect request header, it's unsafe
	//router.TrustedPlatform = "X-CDN-IP"
	//
	//router.GET("/", func(c *gin.Context) {
	//	// If you set TrustedPlatform, ClientIP() will resolve the
	//	// corresponding header and return IP directly
	//	fmt.Printf("ClientIP: %s\n", c.ClientIP())
	//})
	//router.GET("/albums", getAlbums)
	//router.POST("/albums", postAlbums)
	//
	//err := router.Run()
	//if err != nil {
	//	return
	//}
}
