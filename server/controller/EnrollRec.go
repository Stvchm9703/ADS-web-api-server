package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"webserver/server/common"

	m "webserver/server/model"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// Comtroller
func CreateEnroll(c *gin.Context) {
	id, e := c.Params.Get("stud_id")
	if !e {
		RespondJSONWithError(c, 500, map[string]interface{}{
			"err": "no param of deot_id",
		})
	} else {
		var tem m.EnrollMod
		var test map[string]interface{}
		if c.BindJSON(&test) == nil {
			testB, err := m.TestEnroll(id, test)
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

					if k, err := m.CreateEnroll(id, &tem); err != nil {
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

func GetEnrollList(c *gin.Context) {
	id, e := c.Params.Get("stud_id")
	if !e {
		RespondJSONWithError(c, 500, map[string]interface{}{
			"err": "no param of course_id",
		})
	} else {
		var PS m.PageMeta
		PS.PageLimit, _ = strconv.Atoi(c.Query("pl"))
		PS.PageNum, _ = strconv.Atoi(c.Query("pn"))
		fmt.Println("query map", c.Request.URL.Query())
		fmt.Println("PS", PS)
		// search
		o := BindQuery(c.Request.URL.Query(), m.EnrollMod{})
		if len(id) == 24 {
			k, PS1, err2 := m.FetchStudEnroll(id, o, &PS)
			if err2 != nil {
				RespondJSONWithError(c, 500, err2)
			} else {
				RespondJSON(c, 200, k, PS1)
			}
		} else {
			fmt.Println("FetchAllEnrolled")
			k, PS1, err2 := m.FetchAllEnroll(o, &PS)
			if err2 != nil {
				RespondJSONWithError(c, 500, err2)
			} else {
				RespondJSON(c, 200, k, PS1)
			}
		}
	}

}

func GetEnroll(c *gin.Context) {
	fmt.Println(c.Params)
	id, err := c.Params.Get("stud_id")
	e_id, e3 := c.Params.Get("e_id")
	if !err || !e3 {
		RespondJSONWithError(c, 500, map[string]interface{}{
			"err": "no param of dept_id or id",
		})
	} else {
		k, err2 := m.GetEnroll(id, e_id)
		if (err2) != nil {
			RespondJSONWithError(c, 500, err)
		} else {
			RespondJSON(c, 200, k, nil)
		}
	}
}

func UpdateEnroll(c *gin.Context) {
	id, err3 := c.Params.Get("stud_id")
	e_id, er2 := c.Params.Get("e_id")
	var ftem map[string]interface{}
	var tem m.EnrollMod
	if !err3 || !er2 {
		RespondJSONWithError(c, 500, map[string]interface{}{
			"err": "no param of dept_id or course_id",
		})
	} else {
		if c.BindJSON(&ftem) == nil {
			k1, errr := m.GetEnroll(id, e_id)
			fmt.Println("k1", k1)
			if errr != nil {
				RespondJSONWithError(c, 500, errr)
			} else {
				if k1 != nil {
					newO, err := json.Marshal(ftem)
					fmt.Println("newO", string(newO))
					if err = json.Unmarshal(newO, &tem); err == nil {
						k, err := m.UpdateEnroll(id, k1, &tem)
						if err != nil {
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

func DeleteEnroll(c *gin.Context) {
	id, err3 := c.Params.Get("stud_id")
	cid, er2 := c.Params.Get("e_id")
	var ftem map[string]interface{}
	if !err3 || !er2 {
		RespondJSONWithError(c, 500, map[string]interface{}{
			"err": "no param of dept_id or course_id",
		})
	} else {
		if c.BindJSON(&ftem) == nil {
			notexist, err := m.TestEnroll(id, map[string]interface{}{
				"_id": bson.ObjectIdHex(cid),
			})
			if err != nil {
				RespondJSONWithError(c, 500, err)
			} else {
				if !notexist {
					_, err := m.DeleteEnroll(id, cid)
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
		} else {
			BindingErr(c, &m.EnrollMod{})
		}
	}
}
