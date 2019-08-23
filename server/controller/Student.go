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
func CreateStudent(c *gin.Context) {
	var tem m.StudentMod
	var test map[string]interface{}
	if c.BindJSON(&test) == nil {
		fmt.Println(test)
		testB, err := m.TestStudent(test)
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
				fmt.Println(tem)
				if k, err := m.CreateStudent(&tem); err != nil {
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

func GetStudentList(c *gin.Context) {
	var PS m.PageMeta
	PS.PageLimit, _ = strconv.Atoi(c.Query("pl"))
	PS.PageNum, _ = strconv.Atoi(c.Query("pn"))
	fmt.Println("query map", c.Request.URL.Query())
	fmt.Println("PS", PS)
	o := BindQuery(c.Request.URL.Query(), m.StudentMod{})
	fmt.Println("o", o)
	k, PS1, err2 := m.FetchStudent(o, &PS)
	fmt.Println(k, PS1, err2)
	if err2 != nil {
		RespondJSONWithError(c, 500, err2)
	} else {
		RespondJSON(c, 200, k, PS1)
	}

}

func GetStudent(c *gin.Context) {
	fmt.Println(c.Params)
	id, err := c.Params.Get("stud_id")
	if err == false {
		RespondJSONWithError(c, 500, err)
	} else {
		k, err2 := m.GetStudent(id)
		if (err2) != nil {
			RespondJSONWithError(c, 500, err)
		} else {
			RespondJSON(c, 200, k, nil)
		}
	}
}

func UpdateStudent(c *gin.Context) {
	var ftem map[string]interface{}
	var tem m.StudentMod
	if c.BindJSON(&ftem) == nil {
		k1, errr := m.GetStudent(ftem["_id"].(string))
		if errr != nil {
			log.Fatalln(errr)
			RespondJSONWithError(c, 500, errr)
		} else {
			if k1 != nil {
				newO, err := json.Marshal(ftem)
				if err = json.Unmarshal(newO, &tem); err == nil {
					k, err := m.UpdateStudent(k1, &tem)
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

func DeleteStudent(c *gin.Context) {
	var ftem map[string]interface{}
	if c.BindJSON(&ftem) == nil {
		fmt.Println(ftem)
		notexist, err := m.TestStudent(map[string]interface{}{
			"_id": bson.ObjectIdHex(ftem["_id"].(string)),
		})
		if err != nil {
			RespondJSONWithError(c, 500, err)
		} else {
			if !notexist {
				_, err := m.DeleteStudent(ftem["_id"].(string))
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
		BindingErr(c, &m.StudentMod{})
	}
}
