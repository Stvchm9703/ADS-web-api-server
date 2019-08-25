package controller

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"webserver/server/common"

	"github.com/gin-gonic/gin"
)

func MdLoader() gin.HandlerFunc {
	fmt.Println()
	fmt.Println(common.PathInRun)
	return func(c *gin.Context) {
		posty := c.Request.URL.String()
		if strings.Contains(posty, "/md/") && strings.Index(posty, ".md") == len(posty)-3 {
			mdfile, err := ioutil.ReadFile(path.Join(common.PathInRun, posty))
			if err != nil {
				fmt.Println(err)
				c.Abort()
			}
			c.Data(200, "text/html; charset=utf-8", mdfile)
		}
		c.Next()

	}
}

var Errpage []byte

func init() {
	Errpage, _ = ioutil.ReadFile(path.Join(common.PathInRun, "template", "error.html"))
}
func ErrorPage(c *gin.Context) {
	c.Data(200, "text/html; charset=utf-8", Errpage)
}
