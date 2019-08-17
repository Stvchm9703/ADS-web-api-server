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
		// department
		v1.GET("/l/dept", c.GetDepartmentList)
		v1.GET("/g/dept/:dept_id", c.GetDepartment)
		v1.POST("/c/dept", c.CreateDepartment)
		v1.POST("/d/dept", c.DeleteDepartment)
		v1.POST("/u/dept", c.UpdateDepartment)

		//  Course
		v1.GET("/l/dept/:dept_id/course", c.GetCourseList)
		v1.GET("/g/dept/:dept_id/course/:course_id", c.GetCourse)
		v1.POST("/c/dept/:dept_id/course", c.CreateCourse)
		v1.POST("/d/dept/:dept_id/course/:course_id", c.DeleteCourse)
		v1.POST("/u/dept/:dept_id/course/:course_id", c.UpdateCourse)

		// //  Offer
		// v1.GET("/l/dept/:dept_id/course/:course_id/offer", c.GetOfferList)
		// v1.GET("/g/dept/:dept_id/offer/:offer_id", c.GetOffer)
		// v1.POST("/c/dept/:dept_id/course/:course_id/offer", c.CreateOffer)
		// v1.POST("/d/dept/:dept_id/offer/:offer_id", c.DeleteOffer)
		// v1.POST("/u/dept/:dept_id/offer/:offer_id", c.UpdateOffer)

		// // Student
		// v1.GET("/l/student", c.GetStudentList)
		// v1.GET("/g/student/:stud_id", c.GetStudent)
		// v1.POST("/c/student", c.CreateStudent)
		// v1.POST("/d/student", c.DeleteStudent)
		// v1.POST("/u/student", c.UpdateStudent)

		// // Enroll
		// v1.GET("/l/student/:stud_id/enrolled", c.GetEnrollList)
		// v1.POST("/c/student/:stud_id/enrolled", c.CreateEnroll)
		// v1.POST("/d/student/:stud_id/enrolled", c.DeleteEnroll)
		// v1.POST("/u/student/:stud_id/enrolled", c.UpdateEnroll)

		// Enroll Course Detail
		v1.GET("/l/student/:stud_id/enrolled/:course_id", c.GetEnrollList)

		log.Println("v1")
		log.Printf("Group:  %+v", v1)
		log.Printf("Handler:  %+v", v1.Handlers)
	}

	return router
}

// func RouterSetting(config *conf.ConfigTemp) http.Handler {
// 	router := gin.New()
// 	router.Use(gin.Recovery())
// 	// router.Use(c.RecoverMW())
// 	// router.Use(DummyMiddleware)
// 	v1 := router.Group("/api/v1")
// 	{
// 		// Course list
// 		v1.GET("/course/list", c.GetCourseList)
// 		v1.GET("/course/l", c.GetCourseList)
// 		v1.GET("/course/get/:id", c.GetCourse)
// 		v1.GET("/course/g/:id", c.GetCourse)
// 		v1.POST("/course/create", c.CreateCourse)
// 		v1.POST("/course/c", c.CreateCourse)
// 		v1.POST("/course/delete", c.DeleteCourse)
// 		v1.POST("/course/d", c.DeleteCourse)
// 		v1.POST("/course/update", c.UpdateCourse)
// 		v1.POST("/course/u", c.UpdateCourse)

// 		// department
// 		v1.GET("/dept/list", c.GetDepartmentList)
// 		v1.GET("/dept/l", c.GetDepartmentList)
// 		v1.GET("/dept/get/:id", c.GetDepartment)
// 		v1.GET("/dept/g/:id", c.GetDepartment)
// 		v1.POST("/dept/create", c.CreateDepartment)
// 		v1.POST("/dept/c", c.CreateDepartment)
// 		v1.POST("/dept/delete", c.DeleteDepartment)
// 		v1.POST("/dept/d", c.DeleteDepartment)
// 		v1.POST("/dept/update", c.UpdateDepartment)
// 		v1.POST("/dept/u", c.UpdateDepartment)

// 		// enroll
// 		v1.GET("/enroll/list", c.GetEnrollList)
// 		v1.GET("/enroll/l", c.GetEnrollList)
// 		v1.GET("/enroll/get/:id", c.GetEnroll)
// 		v1.GET("/enroll/g/:id", c.GetEnroll)
// 		v1.POST("/enroll/create", c.CreateEnroll)
// 		v1.POST("/enroll/c", c.CreateEnroll)
// 		v1.POST("/enroll/delete", c.DeleteEnroll)
// 		v1.POST("/enroll/d", c.DeleteEnroll)
// 		v1.POST("/enroll/update", c.UpdateEnroll)
// 		v1.POST("/enroll/u", c.UpdateEnroll)

// 		// offer
// 		v1.GET("/offer/list", c.GetOfferList)
// 		v1.GET("/offer/l", c.GetOfferList)
// 		v1.GET("/offer/get/:id", c.GetOffer)
// 		v1.GET("/offer/g/:id", c.GetOffer)
// 		v1.POST("/offer/create", c.CreateOffer)
// 		v1.POST("/offer/c", c.CreateOffer)
// 		v1.POST("/offer/delete", c.DeleteOffer)
// 		v1.POST("/offer/d", c.DeleteOffer)
// 		v1.POST("/offer/update", c.UpdateOffer)
// 		v1.POST("/offer/u", c.UpdateOffer)

// 		// student
// 		v1.GET("/student/list", c.GetStudentList)
// 		v1.GET("/student/l", c.GetStudentList)
// 		v1.GET("/student/get/:id", c.GetStudent)
// 		v1.GET("/student/g/:id", c.GetStudent)
// 		v1.POST("/student/create", c.CreateStudent)
// 		v1.POST("/student/c", c.CreateStudent)
// 		v1.POST("/student/delete", c.DeleteStudent)
// 		v1.POST("/student/d", c.DeleteStudent)
// 		v1.POST("/student/update", c.UpdateStudent)
// 		v1.POST("/student/u", c.UpdateStudent)

// 		//view model : course-offer
// 		v1.GET("/v/course/list", v.GetCourseList)
// 		v1.GET("/v/course/l", v.GetCourseList)
// 		v1.GET("/v/course/get/:id", v.GetCourse)
// 		v1.GET("/v/course/g/:id", v.GetCourse)

// 		// view model : Student-Enrolled
// 		v1.GET("/v/std_enroll/list", v.GetEnrollList)
// 		v1.GET("/v/std_enroll/l", v.GetEnrollList)
// 		v1.GET("/v/std_enroll/get/:id", v.GetEnroll)
// 		v1.GET("/v/std_enroll/g/:id", v.GetEnroll)

// 		log.Println("v1")
// 		log.Printf("Group:  %+v", v1)
// 		log.Printf("Handler:  %+v", v1.Handlers)
// 	}

// 	return router
// }
