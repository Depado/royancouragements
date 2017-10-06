package views

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/Depado/royancouragements/conf"
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
	var a Athlete
	if !conf.C.Accept {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	var form SendForm
	if err = c.Bind(&form); err != nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"athletes": athletes, "posted": "errors", "accept": conf.C.Accept})
		return
	}
	if a, err = NewFromString(form.Athlete); err != nil {
		c.HTML(http.StatusBadRequest, "index.tmpl", gin.H{"athletes": athletes, "posted": "errors", "accept": conf.C.Accept})
	}
	logrus.WithFields(logrus.Fields{
		"athlete": a,
		"from":    fmt.Sprintf("%s <%s>", form.Name, form.Mail),
		"message": form.Message,
	}).Info("Got message")

	var resp *http.Response
	v := url.Values{}
	v.Add("epr", "20")
	v.Add("bib", form.Athlete)
	v.Add("nam", form.Name)
	v.Add("mel", form.Mail)
	v.Add("msg", form.Message)
	if resp, err = http.Get(fmt.Sprintf("http://francechrono.fr/lesdirects/doencourage.php?%v", v.Encode())); err != nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"athletes": athletes, "posted": "remote_error", "accept": conf.C.Accept})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"athletes": athletes, "posted": "remote_error", "accept": conf.C.Accept})
		return
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"athletes": athletes, "posted": "success", "accept": conf.C.Accept})
}
