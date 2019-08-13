package viewMod

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"webserver/server/common"
	m "webserver/server/model"

	"gopkg.in/mgo.v2/bson"
)

type VStudentEnrolledMod struct {
	// EnrollMod
	ID         *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty" `
	StudentID  *string        `bson:"student_id,omitempty" json:"student_id,omitempty"`
	Year       *int           `bson:"year,omitempty" json:"year,omitempty"`
	EnrollDate *time.Time     `bson:"enroll_date,omitempty" json:"enroll_date,omitempty"`
	CourseID   *string        `bson:"course_id,omitempty" json:"course_id,omitempty"`
	CreatedAt  *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	// StudentMod
	StuName *string `bson:"student_name,omitempty" json:"student_name,omitempty"`
	// VCourseOfferMod
	Title           *string `bson:"title,omitempty" json:"title,omitempty"`
	Level           *int    `bson:"level,omitempty" json:"level,omitempty"`
	ClassSize       *int    `bson:"class_size,omitempty" json:"class_size,omitempty"`
	AvailablePlaces *int    `bson:"available_places,omitempty" json:"available_places,omitempty"`
	DeptID          *string `bson:"dept_id,omitempty" json:"dept_id,omitempty"`
	DeptName        *string `bson:"dept_name,omitempty" json:"dept_name,omitempty"`
	Address         *string `bson:"address,omitempty" json:"address,omitempty"`
}

var v_stud_enroll_name = "VStudentEnrolled"
var v_stud_enroll_basemod = "Enroll"
var v_stud_enroll_schema = []bson.M{
	bson.M{
		"$lookup": bson.M{
			"localField":   "student_id",
			"from":         "Student",
			"foreignField": "student_id",
			"as":           "std",
		},
	},
	bson.M{
		"$lookup": bson.M{
			"localField":   "course_id",
			"from":         "VCourseOffer",
			"foreignField": "course_id",
			"as":           "vcs",
		},
	},
	bson.M{
		"$project": bson.M{
			"_id":              1,
			"year":             1,
			"student_id":       1,
			"enroll_date":      1,
			"course_id":        1,
			"created_at":       1,
			"updated_at":       1,
			"student_name":     "std.student_name",
			"title":            "vcs.title",
			"level":            "vcs.level",
			"class_size":       "vcs.class_size",
			"available_places": "vcs.available_places",
			"dept_id":          "vcs.dept_id",
			"dept_name":        "vcs.dept_name",
			"address":          "vcs.address",
		},
	},
}

// FetchVStudentEnrolled : Get VStudentEnrolled Object list
func FetchVStudentEnrolled(param interface{}, ps *m.PageMeta) ([]*VStudentEnrolledMod, *m.PageMeta, error) {
	var record []*VStudentEnrolledMod
	nps := m.PageMeta{}
	// fmt.Println("req. params", param)
	if m.DBConn != nil {
		count, err := m.DBConn.C(v_stud_enroll_name).Find(&param).Count()
		if err != nil {
			return nil, nil, err
		}
		// // fmt.Println("count:", count)
		Q := m.DBConn.C(v_stud_enroll_name).Find(param)
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
			log.Fatal(err)
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
	for ks := 0; ks < len(v_stud_enroll_schema); ks++ {
		tmp := v_stud_enroll_schema[ks]
		if tmp["$lookup"] != nil {
			tmp1 := tmp["$lookup"].(bson.M)
			bsonCheckName = append(bsonCheckName, tmp1["from"].(string))
		}
	}
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
		} else {
			for j := 0; j < len(bsonCheckName); j++ {
				if bsonCheckName[j] == tmp {
					bsonCheckName = append(bsonCheckName[:j], bsonCheckName[j+1:]...)

					continue I
				}
			}
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
