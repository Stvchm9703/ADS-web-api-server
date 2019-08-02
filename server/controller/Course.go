package controller

// import (
// 	_ "webserver/server"
// 	"webserver/server/model"

// 	"github.com/gin-gonic/gin"
// )

// // func InsertCourse(w *gin.Context) {
// // 	_, err := DB.create(model.CoursesMod)
// // }


import (
	// m "wildbase/pkg/wildbase/model"
	"fmt"
	"log"
	"strconv"
	
	"webserver/server/controller/util"
	m "webserver/server/model"

	"github.com/gin-gonic/gin"
)

// Comtroller
func CreateGroup(c *gin.Context) {
	var tem m.UserGroupMod
	if c.BindJSON(&tem) == nil {
		k, err := IF.Create(&tem)
		if len(err) > 0 {
			log.Println(err)
			util.RespondJSONWithError(c, 500, err)
		} else {
			util.RespondJSON(c, 200, k)
		}
	} else {
		util.RespondJSONWithError(c, 500, nil)
	}
}

func GetGroupList(c *gin.Context) {
	// k, err := IF.Fetch("")
	if len(err) > 0 {
		log.Println(err)
		util.RespondJSONWithError(c, 500, err)
	} else {
		util.RespondJSON(c, 200, k)
	}
}

func GetGroup(c *gin.Context) {
	fmt.Println(c.Params)
	id, err := c.Params.Get("id")
	if err == false {
		util.RespondJSONWithError(c, 500, err)
	} else {
		// k, err2 := IF.Get(id)

		if len(err2) > 0 {
			util.RespondJSONWithError(c, 500, err)
		} else {
			util.RespondJSON(c, 200, k)
		}
	}
}

func UpdateGroup(c *gin.Context) {
	var tem m.UserGroupMod
	if c.BindJSON(&tem) == nil {
		k, err := IF.Update(&tem)
		if len(err) > 0 {
			log.Println(err)
			util.RespondJSONWithError(c, 500, err)
		} else {
			util.RespondJSON(c, 200, k)
		}
	} else {
		util.RespondJSONWithError(c, 500, nil)
	}
}

func DeleteGroup(c *gin.Context) {
	var tem m.UserGroupMod
	if c.BindJSON(&tem) == nil {
		val := strconv.FormatUint(uint64(tem.ID), 10)
		k, err := IF.Delete(&val)
		if len(err) > 0 {
			log.Println(err)
			util.RespondJSONWithError(c, 500, err)
		} else {
			util.RespondJSON(c, 200, k)
		}
	} else {
		util.RespondJSONWithError(c, 500, nil)
	}
}
