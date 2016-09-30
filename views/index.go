package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index is the main view
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"athletes": athletes, "posted": "nope"})
}
