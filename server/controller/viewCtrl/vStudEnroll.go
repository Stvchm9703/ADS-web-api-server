package viewCtrl

import (
	"fmt"
	"strconv"

	ctrl "webserver/server/controller"
	m "webserver/server/model"
	v "webserver/server/model/viewMod"

	"github.com/gin-gonic/gin"
)

func GetEnrollList(c *gin.Context) {
	// k, err := IF.Fetch("")
	var PS m.PageMeta
	PS.PageLimit, _ = strconv.Atoi(c.Query("pl"))
	PS.PageNum, _ = strconv.Atoi(c.Query("pn"))
	PS.Sort, PS.SortAr = ctrl.BindSort( c.Query("sort"),v.VStudentEnrolledMod{})
	fmt.Println("query map", c.Request.URL.Query())
	fmt.Println("PS", PS)
	// search
	o := ctrl.BindQuery(c.Request.URL.Query(), v.VStudentEnrolledMod{})
	fmt.Println("o", o)

	k, PS1, err2 := v.FetchVStudentEnrolled(o, &PS)
	fmt.Println(k, PS1, err2)
	if err2 != nil {
		ctrl.RespondJSONWithError(c, 500, err2)
	} else {
		ctrl.RespondJSON(c, 200, k, PS1)
	}

}

func GetEnroll(c *gin.Context) {
	fmt.Println(c.Params)
	id, err := c.Params.Get("id")
	if err == false {
		ctrl.RespondJSONWithError(c, 500, err)
	} else {
		k, err2 := v.GetVStudentEnrolled(id)
		if (err2) != nil {
			ctrl.RespondJSONWithError(c, 500, err)
		} else {
			ctrl.RespondJSON(c, 200, k, nil)
		}
	}
}
