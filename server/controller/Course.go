package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
	"webserver/server/common"

	m "webserver/server/model"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// Comtroller
func CreateCourse(c *gin.Context) {
	var tem m.CourseMod
	var test map[string]interface{}
	if c.BindJSON(&test) == nil {
		fmt.Println(test)
		testB, err := m.TestCourse(test)
		if err != nil {
			log.Fatalln(err)
			RespondJSONWithError(c, 500, err)
		} else if !testB {
			log.Print("exist")
			RespondJSONWithError(c, 500, common.ErrorMessage{
				When: time.Now(),
				What: "create object error, existed oject",
			})
		} else {
			newO, err := json.Marshal(test)
			if err = json.Unmarshal(newO, &tem); err == nil {
				fmt.Println(tem)
				if k, err := m.CreateCourse(&tem); err != nil {
					log.Println(err)
					RespondJSONWithError(c, 500, err)
				} else {
					RespondJSON(c, 200, k, nil)
				}
			} else {
				fmt.Println("err fall:", tem)
				BindingErr(c, tem)
			}
		}
	} else {
		BindingErr(c, test)
	}
}

func GetCourseList(c *gin.Context) {
	// k, err := IF.Fetch("")
	var PS m.PageMeta
	PS.PageLimit, _ = strconv.Atoi(c.Query("pl"))
	PS.PageNum, _ = strconv.Atoi(c.Query("pn"))
	fmt.Println("query map", c.Request.URL.Query())
	fmt.Println("PS", PS)
	// search
	o := BindQuery(c.Request.URL.Query(), m.CourseMod{})
	fmt.Println("o", o)
	// NOTE: test case query
	// o = &bson.M{
	// 	"level": bson.M{
	// 		"$in": []float64{1.0, 1, 0, 2},
	// 	},
	// }
	k, PS1, err2 := m.FetchCourse(o, &PS)
	fmt.Println(k, PS1, err2)
	if err2 != nil {
		RespondJSONWithError(c, 500, err2)
	} else {
		RespondJSON(c, 200, k, PS1)
	}

}

func GetCourse(c *gin.Context) {
	fmt.Println(c.Params)
	id, err := c.Params.Get("id")
	if err == false {
		RespondJSONWithError(c, 500, err)
	} else {
		k, err2 := m.GetCourse(id)
		if (err2) != nil {
			RespondJSONWithError(c, 500, err)
		} else {
			RespondJSON(c, 200, k, nil)
		}
	}
}

func UpdateCourse(c *gin.Context) {
	var ftem map[string]interface{}
	var tem m.CourseMod
	if c.BindJSON(&ftem) == nil {
		k1, errr := m.GetCourse(ftem["_id"].(string))
		if errr != nil {
			log.Fatalln(errr)
			RespondJSONWithError(c, 500, errr)
		} else {
			if k1 != nil {
				newO, err := json.Marshal(ftem)
				// fmt.Println("newO:", string(newO))
				if err = json.Unmarshal(newO, &tem); err == nil {
					// fmt.Println(tem)
					k, err := m.UpdateCourse(k1, &tem)
					if err != nil {
						log.Println(err)
						RespondJSONWithError(c, 500, err)
					} else {
						RespondJSON(c, 200, k, nil)
					}
				} else {
					// fmt.Println("newO", string(newO))
					// fmt.Println("err fall:", tem)
					BindingErr(c, tem)
				}

			} else {
				RespondJSONWithError(c, 500, common.ErrorMessage{
					When: time.Now(),
					What: "no Object is find for update",
				})
			}
		}
	} else {
		BindingErr(c, &tem)
	}
}

func DeleteCourse(c *gin.Context) {
	var ftem map[string]interface{}
	if c.BindJSON(&ftem) == nil {
		fmt.Println(ftem)
		notexist, err := m.TestCourse(map[string]interface{}{
			"_id": bson.ObjectIdHex(ftem["_id"].(string)),
		})
		if err != nil {
			RespondJSONWithError(c, 500, err)
		} else {
			if !notexist {
				_, err := m.DeleteCourse(ftem["_id"].(string))
				if err != nil {
					log.Println(err)
					RespondJSONWithError(c, 500, err)
				} else {
					RespondJSON(c, 200, true, nil)
				}
			} else {
				RespondJSONWithError(c, 500, common.ErrorMessage{
					When: time.Now(),
					What: "no exist Object is find for delete",
				})
			}
		}
	} else {
		fmt.Println(ftem)
		BindingErr(c, &m.CourseMod{})
	}
}
