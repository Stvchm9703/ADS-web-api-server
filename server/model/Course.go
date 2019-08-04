package model

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

type CourseMod struct {
	ID       string `bson:"_id", json:"_id"`
	CourseID string `bson:"course_id", json:"course_id"`
	Title    string `bson:"title",json:"title"`
	Level    int    `bson:"level",bson:"level"`
}

func FetchCourse(param interface{}, ps *PageMeta) (record *[]CourseMod, nps *PageMeta, err error) {
	if DBSess != nil && DBConn != nil {
		count, err = DBConn.C("Course").Find(param).All(&record).Count()
		if err != nil {
			log.Fatal(err)
			record = nil
			nps = nil
			return
		}
		err = nil
		nps.Count = count
		return
	} else {
		record, err = NotConn()
		return
	}
}

func GetCourse(id string) (*CourseMod, error) {
	if DBSess != nil && DBConn != nil {
		var result *CourseMod
		err := DBConn.C("Course").Find(bson.M{
			"_id": id,
		}).One(&result)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return result, nil
	} else {
ß		return NotConn()
	}
}

func CreateCourse(cp *CourseMod) (*CourseMod, error) {
	if DBSess != nil && DBConn != nil {
		err := DBConn.C("Course").Insert(cp)
		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}
		return cp, nil
	} else {
		return NotConn()
	}
}

func UpdateCourse(Old *CourseMod, New *CourseMod) (*CourseMod, error) {
	if DBSess != nil && DBConn != nil {
		err := DBConn.C("Course").Update(Old, New)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return New, nil
	} else {
		return NotConn()
	}
}

func DeleteCourse(cpid *string) (bool, error) {
	if DBSess != nil && DBConn != nil {
		err := DBConn.C("Course").Remove(&bson.M{"_id": cpid})
		if err != nil {
			log.Fatal("Got a real error:", err.Error())
			return false, err
		} else {
			return true, nil
		}

	} else {
		return NotConn()
	}
}

func TestCourse(param *CourseMod) (bool, error) {
	if DBSess != nil && DBConn != nil {
		count, err := DBConn.C("Course").Find(param).Count()
		if err != nil {
			return false, err
		} else {
			return (count == 0), nil
		}
	} else {
		return NotConn()
	}
}
