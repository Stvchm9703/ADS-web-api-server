package controller

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"webserver/server/model"

	Ch "github.com/ahl5esoft/golang-underscore"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type ResponseData struct {
	model.PageMeta
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
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
	if pageSet != nil {
		res.Count = pageSet.Count
		res.PageLimit = pageSet.PageLimit
		res.PageNum = pageSet.PageNum
	}
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

func RecoverMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.Errors.Last(); err != nil {
			RespondJSONWithError(c, 500, err)
		}
		// c.Next()
	}
}

// --------------------------------

func BindingErr(c *gin.Context, exp interface{}) {
	temp := map[string]interface{}{}
	t := reflect.TypeOf(exp)
	for i := 0; i < t.NumField(); i++ {
		temp[t.Field(i).Name] = t.Field(i).Type.String()
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
			// fmt.Println("q:", q[tg][0])
			cond := buildBsonM(q[tg][0], &fi)
			Qbson = append(Qbson, bson.M{tg: cond})
		}
	}
	// fmt.Println("len : ", len(Qbson))
	// fmt.Println("Qbson", Qbson)
	if len(Qbson) > 1 {
		bsonQ["$or"] = Qbson
	} else if len(Qbson) == 1 {
		bsonQ = Qbson[0]
	}
	return &bsonQ
}

func buildBsonM(q string, ref *reflect.StructField) bson.M {
	// cond := []bson.M{}
	cond := bson.M{}
	var t map[string]interface{}
	if err := json.Unmarshal([]byte(q), &t); err != nil {
		fmt.Println(err)
		cond = buildFmBson(q, ref)
	} else {
		cond = buildFmJson(t, ref)
	}
	return cond
}

func maxVal(arr []interface{}) float64 {
	var m float64
	for i, e := range arr {
		if i == 0 || e.(float64) > m {
			m = e.(float64)
		}
	}
	return m
}

func minVal(arr []interface{}) float64 {
	var m float64
	for i, e := range arr {
		if i == 0 || e.(float64) < m {
			m = e.(float64)
		}
	}
	return m
}

func buildFmJson(t map[string]interface{}, ref *reflect.StructField) bson.M {
	cond := bson.M{}
	// fmt.Println("t_len :", len(t))
	for k, v := range t {
		// fmt.Println("KEY:", k)
		fmt.Println("v:", v, ",type:", reflect.TypeOf(v))

		isExist := Ch.Chain(queryRume).FindIndex(func(e string, _ int) bool {
			return k == e
		})

		// fmt.Println("isExist:", isExist)
		if isExist > -1 {
			// fmt.Println("ref.Type:", ref.Type)
			// fmt.Println("ref", ref)

			// fmt.Println("v:", reflect.TypeOf(v).String())

			switch k {
			case "in", "nin":
				{ // in or not-in case
					switch ref.Type.String() {
					
					case "struct", "interface":
						fmt.Println("igorn case")
					case "int", "int32", "int64", "float32", "float64", "uint64", "uint32", "uint",
						"*int", "*int32", "*int64", "*float32", "*float64", "*uint64", "*uint32", "*uint":
						if reflect.TypeOf(v).String() == "[]interface {}" {
							array := []float64{}
							for _, vv := range v.([]interface{}) {
								array = append(array, vv.(float64))
							}
							if cond["$"+k] != nil {
								cond["$"+k] = append(cond["$"+k].([]float64), array...)
							} else {
								cond["$"+k] = append([]float64{}, array...)
							}
						} else {
							if cond["$"+k] != nil {
								cond["$"+k] = append(cond["$"+k].([]float64), v.(float64))
							} else {
								cond["$"+k] = append([]float64{}, v.(float64))
							}
						}
					default:
						if reflect.TypeOf(v).String() == "[]interface {}" {
							array := []string{}
							for _, vv := range v.([]interface{}) {
								array = append(array, vv.(string))
							}
							if cond["$"+k] != nil {
								cond["$"+k] = append(cond["$"+k].([]string), array...)
							} else {
								cond["$"+k] = append([]string{}, array...)
							}
						} else {
							if cond["$"+k] != nil {
								cond["$"+k] = append(cond["$"+k].([]string), v.(string))
							} else {
								cond["$"+k] = append([]string{}, v.(string))
							}
						}
					}
				}
			case "gte", "lte", "lt", "gt":
				{ // number-value
					switch ref.Type.String() {
					case "int", "int32", "int64", "float32", "float64", "uint64", "uint32", "uint",
						"*int", "*int32", "*int64", "*float32", "*float64", "*uint64", "*uint32", "*uint":
						if reflect.TypeOf(v).String() == "[]interface {}" {
							fmt.Println("error, array type cannot use in gt/lt, only get max/min")
							if k == "gte" || k == "gt" {
								cond["$"+k] = maxVal(v.([]interface{}))
							} else {
								cond["$"+k] = minVal(v.([]interface{}))
							}
							fmt.Println("cond[$k]", cond["$"+k])
						} else {
							cond["$"+k] = v.(float64)
						}
					default:
						fmt.Println("refTyper:", ref.Type.String())
						fmt.Println("error, string / interface/ struct type cannot use in gt/lt")
					}
				}
			case "eq", "neq":
				{
					switch ref.Type.String() {
					case "struct", "interface", "*struct", "*interface":
						cond["$"+k] = v
					case "int", "int32", "int64", "float32", "float64", "uint64", "uint32", "uint",
						"*int", "*int32", "*int64", "*float32", "*float64", "*uint64", "*uint32", "*uint":
						if reflect.TypeOf(v).String() == "[]interface {}" {
							fmt.Println("warn, eq/neq is not good for array, use in/nin")
							array := []float64{}
							for _, vv := range v.([]interface{}) {
								array = append(array, vv.(float64))
							}
							if k == "eq" {
								if cond["$in"] != nil {
									cond["$in"] = append(cond["$in"].([]float64), array...)
								} else {
									cond["$in"] = append([]float64{}, array...)
								}
							} else {
								if cond["$nin"] != nil {
									cond["$nin"] = append(cond["$nin"].([]float64), array...)
								} else {
									cond["$nin"] = append([]float64{}, array...)
								}
							}
						} else {
							cond["$"+k] = v.(float64)
						}
					default:
						if reflect.TypeOf(v).String() == "[]interface {}" {
							fmt.Println("warn, eq/neq is not good for array, use in/nin")
							array := []string{}
							for _, vv := range v.([]interface{}) {
								array = append(array, vv.(string))
							}
							if k == "eq" {
								if cond["$in"] != nil {
									cond["$in"] = append(cond["$in"].([]string), array...)
								} else {
									cond["$in"] = append([]string{}, array...)
								}
							} else {
								if cond["$nin"] != nil {
									cond["$nin"] = append(cond["$nin"].([]string), array...)
								} else {
									cond["$nin"] = append([]string{}, array...)
								}
							}
						} else {
							cond["$"+k] = map[string]string{
								"$regex": v.(string),
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("cond:", cond)
	return cond
}

func buildFmBson(q string, ref *reflect.StructField) bson.M {
	cond := bson.M{}
	// fmt.Println("q:", q)
	strcon := strings.Split(q, "_")
	// fmt.Println("strcon:", strcon)
	for i := 0; i < len(strcon); i++ {
		fmt.Println(i, ":strcon:", strcon[i])
		tem := strings.Split(strcon[i], ":")
		arrVal := []string{}
		if strings.Contains(tem[1], "[]") {
			len1 := len([]rune(tem[1]))
			p1 := tem[1]
			arrVal = strings.Split(p1[1:len1-1], ",")
		}
		// fmt.Println("arrVal:", arrVal)
		var val interface{}
		var err error
		switch ref.Type.String() {
		case "int", "int32", "int64", "float32", "float64", "uint64", "uint32", "uint",
			"*int", "*int32", "*int64", "*float32", "*float64", "*uint64", "*uint32", "*uint":
			if len(arrVal) > 0 {
				val = make([]float64, len(arrVal))
				for vk, vv := range arrVal {
					val.([]float64)[vk], err = strconv.ParseFloat(vv, 64)
				}
			} else {
				val, err = strconv.ParseFloat(tem[1], 64)
			}

		case "string", "*string":
			if len(arrVal) > 0 {
				val = arrVal
			} else {
				val = tem[1]
			}

		default:
			val = tem[1]
		}
		if err != nil {
			fmt.Println(err)
		}
		// NOTE: check the tem[0] type
		if cond[tem[0]] == nil {
			vt := reflect.TypeOf(val).String()
			if (tem[0] == "$in" || tem[0] == "$nin") && strings.Contains(vt, "[]") == false {
				cond[tem[0]] = []interface{}{val}
			} else {
				cond[tem[0]] = val
			}
		}

	}
	return cond
}
