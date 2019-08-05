package model

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"webserver/server/common"
)

type CourseMod struct {
	ID       string `bson:"_id", json:"_id"`
	CourseID string `bson:"course_id", json:"course_id"`
	Title    string `bson:"title",json:"title"`
	Level    int    `bson:"level",bson:"level"`
}

// FetchCourse : GEt the Course list
func FetchCourse(param interface{}, ps *PageMeta) ([]*CourseMod, *PageMeta, error) {
	var record []*CourseMod
	var nps *PageMeta

	if DBSess != nil && DBConn != nil {
		count, err := DBConn.C("Course").Find(param).Count()
		Q := DBConn.C("Course").Find(param)
		if ps.PageLimit > 0 {
			Q = Q.Limit(ps.PageLimit)
			nps.PageLimit = ps.PageLimit
		} else {
			Q = Q.Limit(common.QueryDefaultPageLimit)
			nps.PageLimit = common.QueryDefaultPageLimit // default Page Limit
		}
		if ps.PageNum > 0 {
			Q = Q.Skip((ps.PageNum - 1) * ps.PageLimit)
		}
		err1 := Q.All(&record)
		// if ps != nil {
		// 	err1 = DBConn.C("Course").Find(param).Limit(nps.PageLimit).Skip((nps.PageNum - 1) * nps.PageLimit).All(&record)
		// } else {
		// 	err1 = DBConn.C("Course").Find(param).Limit(nps.PageLimit).Skip((nps.PageNum - 1) * nps.PageLimit).All(&record)
		// }

		if err != nil || err1 != nil {
			log.Fatalln("error")
			log.Fatalln(err)
			log.Fatalln(err1)
			log.Fatalln("param")
			log.Fatalln(param)
			if err != nil {
				return nil, nil, err
			} else {
				return nil, nil, err1
			}
		}
		nps.Count = count
		return record, nps, nil
	} else {
		_, err := NotConn()
		return nil, nil, err
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
		_, err := NotConn()
		return nil, err
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
		_, err := NotConn()
		return nil, err
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
		_, err := NotConn()
		return nil, err
	}
}

func DeleteCourse(cpid *string) (bool, error) {
	if DBSess != nil && DBConn != nil {
		err := DBConn.C("Course").Remove(&bson.M{"_id": cpid})
		if err != nil {
			log.Fatal("Got a real error:", err.Error())
			return false, err
		}
		return true, nil
	} else {
		_, err := NotConn()
		return false, err
	}
}

func TestCourse(param *CourseMod) (bool, error) {
	if DBSess != nil && DBConn != nil {
		count, err := DBConn.C("Course").Find(param).Count()
		if err != nil {
			return false, err
		}
		return (count == 0), nil

	} else {
		_, err := NotConn()
		return false, err
	}
}
