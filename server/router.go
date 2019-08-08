package server

import (
	"log"
	"net/http"
	conf "webserver/server/common"

	c "webserver/server/controller"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var router errgroup.Group

// RouterSetting : api routing setting
func RouterSetting(config *conf.ConfigTemp) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	// router.Use(c.RecoverMW())
	// router.Use(DummyMiddleware)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/course/list", c.GetCourseList)
		v1.GET("/course/l", c.GetCourseList)
		v1.GET("/course/get/:id", c.GetCourse)
		v1.GET("/course/g/:id", c.GetCourse)
		v1.POST("/course/create", c.CreateCourse)
		v1.POST("/course/c", c.CreateCourse)
		v1.POST("/course/delete", c.DeleteCourse)
		v1.POST("/course/d", c.DeleteCourse)
		v1.POST("/course/update", c.UpdateCourse)
		v1.POST("/course/u", c.UpdateCourse)

		// v1.GET("/tags/search", Product.GetTagList)
		// v1.GET("/usergroup", UserGroup.GetGroupList)
		// v1.GET("/usergroup/:id", UserGroup.GetGroup)

		// v1.GET("/users/search", User.GetUserList)
		// v1.GET("/user/:id", User.GetUser)
		// v1.GET("/user/:id/contents", User.GetUserContentList)

		// v1.POST("/user/signup", User.CreateUser)
		// v1.POST("/user/login", User.Login)
		log.Println("v1")
		log.Printf("Group:  %+v", v1)
		log.Printf("Handler:  %+v", v1.Handlers)
	}

	return router
}
