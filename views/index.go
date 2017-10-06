package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Depado/royancouragements/conf"
)

// Index is the main view
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"athletes": athletes,
		"posted":   "nope",
		"accept":   conf.C.Accept,
	})
}
