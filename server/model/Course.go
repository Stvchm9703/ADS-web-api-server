package model

import (
	"fmt"
	"time"
	"webserver/server/common"

	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type Courses struct {
	bongo.DocumentBase `bson:",inline"`
	ID                 string `bson:"_id", json:"_id"`
	CourseID           string `bson:"course_id", json:"course_id"`
	Title              string `bson:"title",json:"title"`
	Level              int    `bson:"level",bson:"level"`
}

func FetchCouse(param *bson.M) (*bongo.ResultSet, error) {
	if DBConn != nil {
		result := DBConn.Collection("Course").Find(param)
		// resultSet := []*Courses
		// for results.Next(person) {
		// 	fmt.Println(person.FirstName)
		// }
		return result, nil
	} else {
		return nil, common.ErrorMessage{
			When: time.Now(),
			What: "MongoDB is not Connect",
		}
	}
}

func GetCouse(param *bson.M) (*Courses, error) {
	if DBConn != nil {
		qwe := &Courses{}
		err := DBConn.Collection("Course").FindOne(param, qwe)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		} else {
			return qwe, nil
		}
	} else {
		return nil, common.ErrorMessage{
			When: time.Now(),
			What: "MongoDB is not Connect",
		}
	}
}

func CreateCourse(cp *Courses) (*Courses, error) {
	if DBConn != nil {
		err := DBConn.Collection("Course").Save(cp)
		if vErr, ok := err.(*bongo.ValidationError); ok {
			fmt.Println("Validation errors are:", vErr.Errors)
			return nil, err
		} else {
			fmt.Println("Got a real error:", err.Error())
			return nil , err
		}
		return cp, nil
	} else {
		return nil, common.ErrorMessage{
			When: time.Now(),
			What: "MongoDB is not Connect",
		}
	}
}

func UpdateCourse(cp *Course) (*Course, error) {
	if DBConn != nil {
		err := DBConn.Collection("Course").Save(cp)
		if vErr, ok := err.(*bongo.ValidationError); ok {
			fmt.Println("Validation errors are:", vErr.Errors)
			return nil, err
		} else {
			fmt.Println("Got a real error:", err.Error())
			return nil , err
		}
		return cp, nil
		
	} else {
		return nil, common.ErrorMessage{
			When: time.Now(),
			What: "MongoDB is not Connect",
		}
	}
}

func DeleteCourse(cpid *string) (bool, error){
	if DBConn != nil {
		err := DBConn.Collection("Course").DeleteOne(&bson.M{"_id":cpid})
		if err != nil{
			fmt.Println("Got a real error:", err.Error())
			return false,err
		} else {
			return true , err
		}
		
	} else {
		return nil, common.ErrorMessage{
			When: time.Now(),
			What: "MongoDB is not Connect",
		}
	} 
} 

func CreateSchemaCourse() (bool,error){
	
}