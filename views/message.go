package views

import (
	"fmt"
	"net/http"
	"net/url"

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
	var resp *http.Response
	v := url.Values{}
	v.Add("epr", "187")
	v.Add("bib", form.Athlete)
	v.Add("nam", form.Name)
	v.Add("mel", form.Mail)
	v.Add("msg", form.Message)
	if resp, err = http.Get(fmt.Sprintf("http://www.marchons.com/directlive/doencourage.php?%v", v.Encode())); err != nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"athletes": athletes, "posted": "remote_error"})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"athletes": athletes, "posted": "remote_error"})
		return
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"athletes": athletes, "posted": "success"})
}
