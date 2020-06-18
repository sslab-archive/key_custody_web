package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sslab-archive/key_custody_web/controller"
	)


func main(){
	//repository.LoadKeyList()
	startClientWebServer()
}
func startClientWebServer() {
	defaultRouter := gin.Default()
	defaultRouter.LoadHTMLGlob("templates/*")
	defaultRouter.Use(CORSMiddleware())
	controller.RegisterUserController(defaultRouter)
	defaultRouter.Run()
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Next()
	}
}