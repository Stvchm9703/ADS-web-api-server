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

/**
url : "/c/dept/:dept_id/course"
methods : "POST"
request-body : {
	"course_id" : string , optional,
	"title" : string, optional,
	"level" : int, oprional,
	"offers" : []offer_obj , not-necessary,
}
*/
func CreateCourse(c *gin.Context) {
	id, e := c.Params.Get("dept_id")
	if !e {
		RespondJSONWithError(c, 500, map[string]interface{}{
			"err": "no param of deot_id",
		})
	} else {
		var tem m.CourseMod
		var test map[string]interface{}
		if c.BindJSON(&test) == nil {
			testB, err := m.TestCourse(id, test)
			if err != nil {
				RespondJSONWithError(c, 500, err)
			} else if !testB {
				RespondJSONWithError(c, 500, common.ErrorMessage{
					When: time.Now(),
					What: "create object error, existed oject",
				})
			} else {
				newO, err := json.Marshal(test)
				if err = json.Unmarshal(newO, &tem); err == nil {
					if k, err := m.CreateCourse(id, &tem); err != nil {
						RespondJSONWithError(c, 500, err)
					} else {
						RespondJSON(c, 200, k, nil)
					}
				} else {
					BindingErr(c, tem)
				}
			}
		} else {
			BindingErr(c, test)
		}
	}
}

func GetDeptCourseList(c *gin.Context) {
	id, e := c.Params.Get("dept_id")
	if !e {
		RespondJSONWithError(c, 500, map[string]interface{}{
			"err": "no param of deot_id",
		})
	} else {
		var PS m.PageMeta
		PS.PageLimit, _ = strconv.Atoi(c.Query("pl"))
		PS.PageNum, _ = strconv.Atoi(c.Query("pn"))
		// search
		o := BindQuery(c.Request.URL.Query(), m.CourseMod{})
		var k []*m.CourseMod
		var PS1 *m.PageMeta
		var err2 error
		if len(id) == 24 {
			k, PS1, err2 = m.FetchDeptCourse(id, o, &PS)
		} else {
			k, PS1, err2 = m.FetchAllCourse(o, &PS)
		}
		if err2 != nil {
			RespondJSONWithError(c, 500, err2)
		} else {
			RespondJSON(c, 200, k, PS1)
		}
	}
}

func GetDeptCourse(c *gin.Context) {
	fmt.Println(c.Params)
	id, err := c.Params.Get("course_id")
	dept_id, e3 := c.Params.Get("dept_id")
	if !err || !e3 {
		RespondJSONWithError(c, 500, map[string]interface{}{
			"err": "no param of dept_id or id",
		})
	} else {
		k, err2 := m.GetDeptCourse(dept_id, id)
		if (err2) != nil {
			RespondJSONWithError(c, 500, err)
		} else {
			RespondJSON(c, 200, k, nil)
		}
	}
}

func UpdateCourse(c *gin.Context) {
	id, err3 := c.Params.Get("dept_id")
	course_id, er2 := c.Params.Get("course_id")
	var ftem map[string]interface{}
	var tem m.CourseMod
	if !err3 || !er2 {
		RespondJSONWithError(c, 500, map[string]interface{}{
			"err": "no param of dept_id or course_id",
		})
	} else {
		if c.BindJSON(&ftem) == nil {
			k1, errr := m.GetCourse(id, course_id)
			if errr != nil {
				RespondJSONWithError(c, 500, errr)
			} else {
				if k1 != nil {
					newO, err := json.Marshal(ftem)
					if err = json.Unmarshal(newO, &tem); err == nil {
						k, err := m.UpdateCourse(id, k1, &tem)
						if err != nil {
							log.Println(err)
							RespondJSONWithError(c, 500, err)
						} else {
							RespondJSON(c, 200, k, nil)
						}
					} else {
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
}

func DeleteCourse(c *gin.Context) {
	// var ftem map[string]interface{}
	id, err3 := c.Params.Get("dept_id")
	cid, er2 := c.Params.Get("course_id")
	if !err3 || !er2 {
		RespondJSONWithError(c, 500, map[string]interface{}{
			"err": "no param of dept_id or course_id",
		})
	} else {

		notexist, err := m.TestCourse(id, map[string]interface{}{
			"_id": bson.ObjectIdHex(cid),
		})
		if err != nil {
			RespondJSONWithError(c, 500, err)
		} else {
			if !notexist {
				_, err := m.DeleteCourse(id, cid)
				if err != nil {
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

	}
}
