package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
		router := gin.Default()
		
		router.StaticFS("/client/dev", "./client/dev")
		
		router.GET("/wtf", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello %s", "wtf")
		})
		
		if err := router.RunTLS(":3333", "crt/server.crt", "crt/server.key"); err != nil {
			fmt.Println(err)
		}
}