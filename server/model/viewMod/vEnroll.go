package viewMod

import (
	"encoding/json"
	"fmt"

	// // "log"
	"time"
	"webserver/server/common"
	m "webserver/server/model"

	"gopkg.in/mgo.v2/bson"
)

type VStudentEnrolledMod struct {
	// EnrollMod
	ID *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty" `
	// StudentMod
	StudentID  *string    `bson:"student_id,omitempty" json:"student_id,omitempty"`
	StuName    *string    `bson:"student_name,omitempty" json:"student_name,omitempty"`
	EnrollDate *time.Time `bson:"enroll_date,omitempty" json:"enroll_date,omitempty"`
	// VCourseOfferMod
	CourseOffer *VCourseOfferMod `bson:"course_offer,omitempty" json:"course_offer,omit"`
}

var v_stud_enroll_name = "VStudentEnrolled"
var v_stud_enroll_basemod = "student"
var v_stud_enroll_schema = []bson.M{
	bson.M{"$unwind": "$enrolled"},
	bson.M{
		"$project": bson.M{
			"_id":          "$enrolled._id",
			"student_id":   1,
			"student_name": 1,
			"enroll_date":  "$enrolled.enroll_date",
			"course_offer_id": bson.M{
				"$concat": []interface{}{"$enrolled.course_id", "_", bson.M{"$toString": "$enrolled.year"}},
			},
		},
	},
	bson.M{
		"$lookup": bson.M{
			"from":         "VCourseOffer",
			"localField":   "course_offer_id",
			"foreignField": "course_offer_id",
			"as":           "course_offer",
		},
	},
	bson.M{"$unwind": "$course_offer"},
}

// FetchVStudentEnrolled : Get VStudentEnrolled Object list
func FetchVStudentEnrolled(param interface{}, ps *m.PageMeta) ([]*VStudentEnrolledMod, *m.PageMeta, error) {
	var record []*VStudentEnrolledMod
	var nps m.PageMeta
	// fmt.Println("req. params", param)
	if m.DBConn != nil {
		count, err := m.DBConn.C(v_stud_enroll_name).Find(&param).Count()
		if err != nil {
			return nil, nil, err
		}
		// // fmt.Println("count:", count)
		Q := m.DBConn.C(v_stud_enroll_name).Find(param)
		if len(ps.Sort) > 0 {
			Q = Q.Sort(ps.SortAr...)
			nps.Sort, nps.SortAr = ps.Sort, ps.SortAr
		}
		if ps.PageLimit > 0 {
			Q = Q.Limit(ps.PageLimit)
			nps.PageLimit = ps.PageLimit
		} else {
			Q = Q.Limit(common.QueryDefaultPageLimit)
			nps.PageLimit = common.QueryDefaultPageLimit // default Page Limit
		}
		// fmt.Println("req. pageNum", ps.PageNum)
		// defAULT : 1
		if ps.PageNum > 0 {
			Q = Q.Skip((ps.PageNum - 1) * ps.PageLimit)
			nps.PageNum = ps.PageNum
		} else {
			nps.PageNum = 1
		}
		// fmt.Println("Q:", Q)
		err1 := Q.All(&record)
		if err1 != nil {
			return nil, nil, err1
		}
		nps.Count = count
		// fmt.Println(record)
		return record, &nps, nil
	}
	_, err := m.NotConn()
	return nil, nil, err
}

// GetVStudentEnrolled : get one VStudentEnrolled object
func GetVStudentEnrolled(id string) (*VStudentEnrolledMod, error) {
	if m.DBConn != nil {
		var result *VStudentEnrolledMod
		err := m.DBConn.C(v_stud_enroll_name).Find(bson.M{
			"_id": bson.ObjectIdHex(id),
		}).One(&result)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return result, nil
	}
	_, err := m.NotConn()
	return nil, err
}

// TestVStudentEnrolled : Test VStudentEnrolled is not existed
func TestVStudentEnrolled(param map[string]interface{}) (bool, error) {
	if m.DBConn != nil {
		count, err := m.DBConn.C(v_stud_enroll_name).Find(&param).Count()
		if err != nil {
			return false, err
		}
		return (count == 0), nil
	}
	_, err := m.NotConn()
	return false, err
}

func CreateSchemaVStudentEnroll() (bool, error) {
	bsonCheckName := []string{}
	// for ks := 0; ks < len(v_stud_enroll_schema); ks++ {
	// 	tmp := v_stud_enroll_schema[ks]
	// 	if tmp["$lookup"] != nil {
	// 		tmp1 := tmp["$lookup"].(bson.M)
	// 		bsonCheckName = append(bsonCheckName, tmp1["from"].(string))
	// 	}
	// }
	bsonCheckName = append(bsonCheckName, v_stud_enroll_basemod)
	fmt.Println("bsonCheckName: ", bsonCheckName)
	nameList := m.MgoCursorRes{}
	m.DBConn.Run(bson.M{
		"listCollections": 1.0,
		"nameOnly":        true,
	}, &nameList)
	fmt.Println("nameList", nameList.Cursor.FirstBatch)

	list := nameList.Cursor.FirstBatch
I:
	for i := 0; i < len(list); i++ {
		tmp := list[i]["name"].(string)
		if tmp == v_stud_enroll_name {
			return false, nil
			break I
			// } else {
			// 	for j := 0; j < len(bsonCheckName); j++ {
			// 		if bsonCheckName[j] == tmp {
			// 			bsonCheckName = append(bsonCheckName[:j], bsonCheckName[j+1:]...)

			// 			continue I
			// 		}
			// 	}
		}
	}
	fmt.Println("bsonCheckName:", bsonCheckName, ", :num:", len(bsonCheckName))
	if len(bsonCheckName) > 0 {
		// fmt.Println("failed , some of Collection is not declared")
		nameJ, _ := json.Marshal(bsonCheckName)
		return false, common.ErrorMessage{
			When: time.Now(),
			What: "failed , some of Collection is not declared : " + string(nameJ),
		}
	}
	resultCursor := m.MgoCursorRes{}
	m.DBConn.Run(bson.M{
		"create":   v_stud_enroll_name,
		"viewOn":   v_stud_enroll_basemod,
		"pipeline": v_stud_enroll_schema,
	}, resultCursor)
	if resultCursor.OK == 1 {
		return true, nil
	}
	nameJ, _ := json.Marshal(resultCursor)
	return false, common.ErrorMessage{
		When: time.Now(),
		What: string(nameJ),
	}
	return false, nil
}
