package controller

import (
	"fmt"
	"log"
	"strconv"

	m "webserver/server/model"

	"github.com/gin-gonic/gin"
)

// Comtroller
func CreateCourse(c *gin.Context) {
	var tem m.CourseMod
	if c.BindJSON(&tem) == nil {
		fmt.Println(tem)
		test, err := m.TestCourse(&tem)

		if !test {
			log.Fatal(err)
			RespondJSONWithError(c, 500, err)
		} else {
			k, err := m.CreateCourse(&tem)
			if err != nil {
				log.Println(err)
				RespondJSONWithError(c, 500, err)
			} else {
				RespondJSON(c, 200, k, nil)
			}
		}
	} else {
		BindingErr(c, tem)
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
	fmt.Println(o)
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
	var tem m.CourseMod
	if c.BindJSON(&tem) == nil {
		k1, errr := m.GetCourse(tem.ID)
		if errr != nil {
			log.Fatalln(errr)
		}
		if k1 != nil {
			k, err := m.UpdateCourse(k1, &tem)
			if err != nil {
				log.Println(err)
				RespondJSONWithError(c, 500, err)
			} else {
				RespondJSON(c, 200, k, nil)
			}
		}
	} else {
		RespondJSONWithError(c, 500, nil)
	}
}

func DeleteCourse(c *gin.Context) {
	var tem m.CourseMod
	if c.BindJSON(&tem) == nil {
		_, err := m.DeleteCourse(&tem.ID)
		if err != nil {
			log.Println(err)
			RespondJSONWithError(c, 500, err)
		} else {
			RespondJSON(c, 200, true, nil)
		}
	} else {
		BindingErr(c, tem)
	}
}
