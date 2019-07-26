package controller

import (
	_ "webserver"
	"webserver/model"

	"github.com/gin-gonic/gin"
)

func InsertCourse(w *gin.Context) {
	_, err := DB.create(model.CoursesMod)
}
