package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendForm is the form required to send a message
type SendForm struct {
	Athlete string `form:"athlete" binding:"required"`
	Name    string `form:"name" binding:"required"`
	Mail    string `form:"mail" binding:"required"`
	Message string `form:"message" binding:"required"`
}

// PostMessage is the handler for the form post
func PostMessage(c *gin.Context) {
	var err error
	var form SendForm
	if err = c.Bind(&form); err != nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"athletes": athletes, "posted": "errors"})
		return
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"athletes": athletes, "posted": "success"})
}
