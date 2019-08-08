package model

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"webserver/server/common"
)

type CourseMod struct {
	ID        string    `bson:"_id" json:"_id,omitempty"`
	CourseID  string    `bson:"course_id" json:"course_id,omitempty"`
	Title     string    `bson:"title" json:"title,omitempty"`
	Level     int       `bson:"level" json:"level,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

var modName = "Course"

var jsonSchema = bson.M{
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
		count, err := DBConn.C(modName).Find(&param).Count()
		if err != nil {
			log.Fatalln("error")
			log.Fatalln(err)
			log.Fatalln("param")
			log.Fatalln(param)
			return nil, nil, err
		}
		// fmt.Println("count:", count)
		Q := DBConn.C(modName).Find(param)
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

func GetCourse(id string) (*CourseMod, error) {
	if DBConn != nil {
		var result *CourseMod
		err := DBConn.C(modName).Find(bson.M{
			"_id": id,
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

func CreateCourse(cp *CourseMod) (*CourseMod, error) {
	if DBConn != nil {
		err := DBConn.C(modName).Insert(cp)
		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}
		return cp, nil
	}
	_, err := NotConn()
	return nil, err

}

func UpdateCourse(Old *CourseMod, New *CourseMod) (*CourseMod, error) {
	if DBConn != nil {
		err := DBConn.C(modName).Update(Old, New)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return New, nil
	}
	_, err := NotConn()
	return nil, err

}

func DeleteCourse(cpid *string) (bool, error) {
	if DBConn != nil {
		err := DBConn.C(modName).Remove(&bson.M{"_id": cpid})
		if err != nil {
			log.Fatal("Got a real error:", err.Error())
			return false, err
		}
		return true, nil
	}
	_, err := NotConn()
	return false, err

}

func TestCourse(param *CourseMod) (bool, error) {
	if DBConn != nil {
		count, err := DBConn.C(modName).Find(param).Count()
		if err != nil {
			return false, err
		}
		return (count == 0), nil

	}
	_, err := NotConn()
	return false, err

}

func CreateCourseFormat() (bool, error) {
	if DBConn != nil {
		result := bson.M{}
		err := DBConn.Run(bson.M{
			"createCollection": bson.M{
				"validationAction": "warn",
				"validator":        jsonSchema,
			},
		}, &result)
		if err != nil {
			return false, err
		}
		fmt.Println(result)
		return true, nil
	}
	_, err := NotConn()
	return false, err
}
