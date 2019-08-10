package model

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"webserver/server/common"

	oid "github.com/coolbed/mgo-oid"
)

type CourseMod struct {
	ID        *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	CourseID  *string        `bson:"course_id,omitempty" json:"course_id,omitempty"`
	Title     *string        `bson:"title,omitempty" json:"title,omitempty"`
	Level     *int           `bson:"level,omitempty" json:"level,omitempty"`
	CreatedAt *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type CourseMP CourseMod

func (P *CourseMP) SetBSON(raw bson.Raw) error {
	return raw.Unmarshal(P)
}

// func CourseIF interface {
// 	FetchCourse()
// }
var course_mod_name = "Course"

var course_json_schema = bson.M{
	"$jsonSchema": bson.M{
		"bsonType": "object",

		"properties": bson.M{
			"_id": bson.M{
				"bsonType":    "objectId",
				"description": "reference id",
			},
			"course_id": bson.M{
				"bsonType":    "string",
				"description": "course id",
			},
			"title": bson.M{
				"description": "course title",
			},
			"level": bson.M{
				"bsonType":    "int",
				"description": "course level",
			},
			"created_at": bson.M{
				"bsonType":    "date",
				"description": "data created time",
			},
			"updated_at": bson.M{
				"bsonType":    "date",
				"description": "data last updated time",
			},
		},
	},
}

// FetchCourse : GEt the Course list
func FetchCourse(param interface{}, ps *PageMeta) ([]*CourseMod, *PageMeta, error) {
	var record []*CourseMod
	nps := PageMeta{}
	fmt.Println("req. params", param)
	if DBConn != nil {
		count, err := DBConn.C(course_mod_name).Find(&param).Count()
		if err != nil {
			log.Fatalln("error")
			log.Fatalln(err)
			log.Fatalln("param")
			log.Fatalln(param)
			return nil, nil, err
		}
		// fmt.Println("count:", count)
		Q := DBConn.C(course_mod_name).Find(param)
		if ps.PageLimit > 0 {
			Q = Q.Limit(ps.PageLimit)
			nps.PageLimit = ps.PageLimit
		} else {
			Q = Q.Limit(common.QueryDefaultPageLimit)
			nps.PageLimit = common.QueryDefaultPageLimit // default Page Limit
		}
		fmt.Println("req. pageNum", ps.PageNum)
		// defAULT : 1
		if ps.PageNum > 0 {
			Q = Q.Skip((ps.PageNum - 1) * ps.PageLimit)
			nps.PageNum = ps.PageNum
		} else {
			nps.PageNum = 1
		}
		fmt.Println("Q:", Q)
		err1 := Q.All(&record)
		if err1 != nil {
			log.Fatalln("error")
			log.Fatalln(err1)
			log.Fatalln("param")
			log.Fatalln(param)
			return nil, nil, err1
		}
		nps.Count = count
		fmt.Println(record)
		return record, &nps, nil
	}
	_, err := NotConn()
	return nil, nil, err

}

// GetCourse : get one Course object
func GetCourse(id string) (*CourseMod, error) {
	if DBConn != nil {
		var result *CourseMod
		err := DBConn.C(course_mod_name).Find(bson.M{
			"_id": bson.ObjectIdHex(id),
		}).One(&result)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return result, nil
	}
	_, err := NotConn()
	return nil, err
}

// CreateCourse : Create a Course Object
func CreateCourse(cp *CourseMod) (*CourseMod, error) {
	if DBConn != nil {
		tnow := time.Now()
		fmt.Println("hi create")
		if cp.ID == nil {
			temID := bson.ObjectId(string(oid.NewOID().Bytes()))
			fmt.Println("temId:", temID.String())
			cp.ID = &temID
			fmt.Println(cp.ID)
		}
		cp.CreatedAt = &tnow
		cp.UpdatedAt = &tnow
		err := DBConn.C(course_mod_name).Insert(&cp)
		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}
		return cp, nil
	}
	_, err := NotConn()
	return nil, err
}

// UpdateCourse : Update a Course Object
func UpdateCourse(Old *CourseMod, New *CourseMod) (*CourseMod, error) {
	if DBConn != nil {
		tnow := time.Now()
		New.UpdatedAt = &tnow
		if New.CreatedAt != Old.CreatedAt {
			// Note: prevent edited
			New.CreatedAt = Old.CreatedAt
		}

		temp, _ := bson.Marshal(New)
		upNew := bson.M{}
		bson.Unmarshal(temp, upNew)
		Returned := CourseMod{}
		_, err := DBConn.C(course_mod_name).Find(bson.M{"_id": Old.ID}).Apply(
			mgo.Change{
				Update:    bson.M{"$set": upNew},
				ReturnNew: true,
			},
			&Returned,
		)
		// fmt.Println("info", info)
		// fmt.Println("Returned", Returned)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return &Returned, nil
	}
	_, err := NotConn()
	return nil, err
}

// DeleteCourse : Delete a Course
func DeleteCourse(cpid string) (bool, error) {
	if DBConn != nil {
		err := DBConn.C(course_mod_name).Remove(&bson.M{"_id": bson.ObjectIdHex(cpid)})
		if err != nil {
			log.Fatal("Got a real error:", err.Error())
			return false, err
		}
		return true, nil
	}
	_, err := NotConn()
	return false, err

}

// TestCourse : Test Course is not existed
func TestCourse(param map[string]interface{}) (bool, error) {
	if DBConn != nil {
		count, err := DBConn.C(course_mod_name).Find(&param).Count()
		if err != nil {
			return false, err
		}
		return (count == 0), nil
	}
	_, err := NotConn()
	return false, err
}

// // CreateCourseFormat : Adv Create Course Object Format
// func CreateCourseFormat() (bool, error) {
// 	if DBConn != nil {
// 		result := bson.M{}
// 		err := DBConn.Run(bson.M{
// 			"createCollection": bson.M{
// 				"validationAction": "warn",
// 				"validator":        course_json_schema,
// 			},
// 		}, &result)
// 		if err != nil {
// 			return false, err
// 		}
// 		fmt.Println(result)
// 		return true, nil
// 	}
// 	_, err := NotConn()
// 	return false, err
// }
