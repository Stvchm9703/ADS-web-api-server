package controller

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"webserver/server/common"

	"github.com/gin-gonic/gin"
)

// var fileMount map[string][]byte
func MdLoader() gin.HandlerFunc {
	fmt.Println()
	fmt.Println("load file dir : default ./template/")
	// var files []string
	// root := path.Join(common.PathInRun, "server", "template")
	// fmt.Println("template root: " , root)
	// err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	// 	// fmt.Println(path)
	// 	if strings.Contains(path, ".") {
	// 		files = append(files, path)
	// 		// fileMount[]
	// 	}
	// 	return nil
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// for _, file := range files {
	// 	fmt.Println(file)
	// }
	return func(c *gin.Context) {
		posty := c.Request.URL.String()
		if strings.Contains(posty, "/md/") && strings.Index(posty, ".md") == len(posty)-3 {
			mdfile, err := ioutil.ReadFile(path.Join(common.PathInRun, "server", posty))
			if err != nil {
				fmt.Println(err)
				c.Abort()
			}
			c.Data(200, "text/html; charset=utf-8", mdfile)
		}
		// } else if strings.Contains(posty, "/api/v1/") {
		c.Next()
		// } else {

		// c.HTML(200, posty, posty+"index.html")
		// }
	}
}

// var fileDir  map[string][string]

// func HtmlLoader() gin.HandlerFunc {
// 	fmt.Println()
// 	fmt.Println("load file dir : default ./template/")
// 	var files []string
// 	root := path.Join(common.PathInRun, "server", "template")
// 	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
// 		files = append(files, path)
// 		return nil
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	return func(c *gin.Context) {
// 		posty := c.Request.URL.String()
// 		if strings.Contains(posty, "/api/v1/") {
// 			c.Next()
// 		} else {

// 		}
// 	}
// }
