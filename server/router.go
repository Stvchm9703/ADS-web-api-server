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
		// Course list
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

		// department
		v1.GET("/dept/list", c.GetDepartmentList)
		v1.GET("/dept/l", c.GetDepartmentList)
		v1.GET("/dept/get/:id", c.GetDepartment)
		v1.GET("/dept/g/:id", c.GetDepartment)
		v1.POST("/dept/create", c.CreateDepartment)
		v1.POST("/dept/c", c.CreateDepartment)
		v1.POST("/dept/delete", c.DeleteDepartment)
		v1.POST("/dept/d", c.DeleteDepartment)
		v1.POST("/dept/update", c.UpdateDepartment)
		v1.POST("/dept/u", c.UpdateDepartment)

		// enroll
		v1.GET("/enroll/list", c.GetEnrollList)
		v1.GET("/enroll/l", c.GetEnrollList)
		v1.GET("/enroll/get/:id", c.GetEnroll)
		v1.GET("/enroll/g/:id", c.GetEnroll)
		v1.POST("/enroll/create", c.CreateEnroll)
		v1.POST("/enroll/c", c.CreateEnroll)
		v1.POST("/enroll/delete", c.DeleteEnroll)
		v1.POST("/enroll/d", c.DeleteEnroll)
		v1.POST("/enroll/update", c.UpdateEnroll)
		v1.POST("/enroll/u", c.UpdateEnroll)

		// offer
		v1.GET("/offer/list", c.GetOfferList)
		v1.GET("/offer/l", c.GetOfferList)
		v1.GET("/offer/get/:id", c.GetOffer)
		v1.GET("/offer/g/:id", c.GetOffer)
		v1.POST("/offer/create", c.CreateOffer)
		v1.POST("/offer/c", c.CreateOffer)
		v1.POST("/offer/delete", c.DeleteOffer)
		v1.POST("/offer/d", c.DeleteOffer)
		v1.POST("/offer/update", c.UpdateOffer)
		v1.POST("/offer/u", c.UpdateOffer)

		// department
		v1.GET("/student/list", c.GetStudentList)
		v1.GET("/student/l", c.GetStudentList)
		v1.GET("/student/get/:id", c.GetStudent)
		v1.GET("/student/g/:id", c.GetStudent)
		v1.POST("/student/create", c.CreateStudent)
		v1.POST("/student/c", c.CreateStudent)
		v1.POST("/student/delete", c.DeleteStudent)
		v1.POST("/student/d", c.DeleteStudent)
		v1.POST("/student/update", c.UpdateStudent)
		v1.POST("/student/u", c.UpdateStudent)

		log.Println("v1")
		log.Printf("Group:  %+v", v1)
		log.Printf("Handler:  %+v", v1.Handlers)
	}

	return router
}
