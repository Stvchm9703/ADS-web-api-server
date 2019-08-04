package controller

import (
	"fmt"
	"reflect"
	"webserver/server/model"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	model.PageMeta
	Status int
	Data   interface{}
}

type RequestData struct {
	model.PageMeta
	Param interface{}
}

/**
 * Json response function
 * */
func RespondJSON(w *gin.Context, status int, payload interface{}, pageSet *model.PageMeta) {
	fmt.Println("status ", status)
	var res ResponseData
	w.Status(status)
	res.Count = pageSet.Count
	res.PageLimit = pageSet.PageLimit
	res.PageNum = pageSet.PageNum
	res.Data = payload
	w.JSON(status, res)
}

func RespondJSONWithError(c *gin.Context, code int, message interface{}) {
	resp := map[string]interface{}{"error": message}
	c.AbortWithStatusJSON(code, resp)
}

func RequestJson(c *gin.Context, mod *interface{}, ps *model.PageMeta) error {
	var tem RequestData
	err := c.BindJSON(&tem)
	if err == nil {
		ps.PageLimit = tem.PageLimit
		ps.PageNum = tem.PageNum
		mod = &tem.Param
		return nil
	} else {
		return err
	}
}

func BindingErr(c *gin.Context, exp interface{}) {
	temp := map[string]interface{}{}
	t := reflect.TypeOf(exp)
	for i := 0; i < t.NumField(); i++ {
		temp[t.Field(i).Name] = t.Field(i).Type
	}
	fmt.Println("temp")
	fmt.Println(temp)

	resp := map[string]interface{}{
		"error": "expected Model cannot pare",
		"model": temp,
		"input": c.Params,
	}
	RespondJSONWithError(c, 510, resp)
}
