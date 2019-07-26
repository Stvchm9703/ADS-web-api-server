package wildbase

import (
	"log"
	"net/http"
	conf "webserver/common"
	"wildbase/pkg/wildbase/controller/Product"
	"wildbase/pkg/wildbase/controller/User"
	"wildbase/pkg/wildbase/controller/UserGroup"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var router errgroup.Group

func RouterSetting(config *conf.ConfigTemp) http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	// router.Use(DummyMiddleware)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/products/search", Product.GetProductList)
		v1.GET("/product/:id", Product.GetProduct)
		v1.GET("/product/:id/comment", Product.GetCommentList)
		v1.GET("/product/:id/comment/:CommentID", Product.GetComment)

		v1.GET("/tags/search", Product.GetTagList)

		v1.GET("/usergroup", UserGroup.GetGroupList)
		v1.GET("/usergroup/:id", UserGroup.GetGroup)

		v1.GET("/users/search", User.GetUserList)
		v1.GET("/user/:id", User.GetUser)
		v1.GET("/user/:id/contents", User.GetUserContentList)

		v1.POST("/user/signup", User.CreateUser)
		// v1.POST("/user/login", User.Login)
		log.Println("v1")
		log.Printf("Group:  %+v", v1)
		log.Printf("Handler:  %+v", v1.Handlers)
	}

	return router
}
