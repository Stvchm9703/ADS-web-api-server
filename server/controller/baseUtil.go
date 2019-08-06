package controller

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"webserver/server/model"

	Ch "github.com/ahl5esoft/golang-underscore"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
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
	// if pageSet != nil {
	res.Count = pageSet.Count
	res.PageLimit = pageSet.PageLimit
	res.PageNum = pageSet.PageNum
	// }
	res.Data = payload
	w.JSON(status, res)
}

func RespondJSONWithError(c *gin.Context, code int, message interface{}) {
	resp := map[string]interface{}{"error": message}
	c.AbortWithStatusJSON(code, resp)
}

func RequestJson(c *gin.Context, mod interface{}, ps *model.PageMeta) error {
	var tem RequestData
	err := c.BindJSON(&tem)
	fmt.Println(tem)
	if err == nil {
		ps.PageLimit = tem.PageLimit
		ps.PageNum = tem.PageNum
		mod = tem.Param
		return nil
	} else {
		fmt.Println(err)
		return err
	}
}

func BindingErr(c *gin.Context, exp interface{}) {
	temp := map[string]interface{}{}
	t := reflect.TypeOf(exp)
	for i := 0; i < t.NumField(); i++ {
		temp[t.Field(i).Name] = t.Field(i).Type
	}
	fmt.Println(temp)
	resp := map[string]interface{}{
		"error": "expected Model cannot pare",
		"model": temp,
		"input": c.Params,
	}
	RespondJSONWithError(c, 510, resp)
}

// rume
var queryRume = []string{"eq", "gt", "gte", "in", "lt", "lte", "ne", "nin"}

// BindQuery
func BindQuery(q map[string][]string, lookup interface{}) *bson.M {
	// fmt.Println(q)
	// fmt.Println(reflect.TypeOf(lookup))
	bsonQ := bson.M{}
	Qbson := []bson.M{}
	v := reflect.ValueOf(lookup)
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		tg := strings.Split(fi.Tag.Get("json"), ",")[0]
		// fmt.Println(fi.Name, fi.Type, tg)
		if q[tg] != nil {
			fmt.Println("q:", q[tg][0])
			cond := buildBsonM(q[tg][0])
			Qbson = append(Qbson, bson.M{tg: cond})
		}
	}
	// fmt.Println("len : ", len(Qbson))
	if len(Qbson) > 0 {
		bsonQ["$or"] = Qbson
	}
	return &bsonQ
}

func buildBsonM(q string) []bson.M {
	cond := []bson.M{}
	var t map[string]interface{}
	json.Unmarshal([]byte(q), &t)
	fmt.Println("t:", t)
	for k, v := range t {
		isExist := Ch.Chain(queryRume).FindIndex(func(e string, _ int) bool {
			return k == e
		})
		if isExist > 0 {
			cond = append(cond, bson.M{"$" + k: v})
		}
	}
	return cond
}
