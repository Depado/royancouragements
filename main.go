package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Depado/royancouragements/conf"
	"github.com/Depado/royancouragements/views"
)

func main() {
	var err error

	if err = conf.Load("conf.yml"); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Static("/static", "./assets")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", views.Index)
	r.POST("/", views.PostMessage)

	r.Run(fmt.Sprintf("%v:%d", conf.C.Host, conf.C.Port))
}
