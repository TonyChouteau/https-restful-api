package main

import {
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := http.ListenAndServeTLS(":8081", "/etc/letsencrypt/live/vps.tonychouteau.fr/fullchain.pem", "/etc/letsencrypt/live/vps.tonychouteau.fr/privkey.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
