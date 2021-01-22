package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sizzflyer/go-vanilla-web/src/rest"
)

func main() {
	r := gin.Default()
	h, err := rest.NewHandler()
	if err != nil {
		log.Println(err)
	}
	r.GET("/", h.Main)
	r.Run("127.0.0.1:5000")
}