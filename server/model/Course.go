package model

import (
	"fmt"

	// "log"
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
	Offers    []*OfferMod    `bson:"offers,omitempty" json:"offers,omitempty"`
	CreatedAt *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

var course_mod_name = "Course"

// FetchCourse : GEt the Course list
func FetchCourse(param interface{}, ps *PageMeta) ([]*CourseMod, *PageMeta, error) {
	var record []*CourseMod
	nps := PageMeta{}
	fmt.Println("req. params", param)
	if DBConn != nil {
		count, err := DBConn.C(course_mod_name).Find(&param).Count()
		if err != nil {
			return nil, nil, err
		}
		Q := DBConn.C(course_mod_name).Find(param)
		if ps.PageLimit > 0 {
			Q = Q.Limit(ps.PageLimit)
			nps.PageLimit = ps.PageLimit
		} else {
			Q = Q.Limit(common.QueryDefaultPageLimit)
			nps.PageLimit = common.QueryDefaultPageLimit // default Page Limit
		}
		if ps.PageNum > 0 {
			Q = Q.Skip((ps.PageNum - 1) * ps.PageLimit)
			nps.PageNum = ps.PageNum
		} else {
			nps.PageNum = 1
		}
		err1 := Q.All(&record)
		if err1 != nil {
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
			fmt.Println(err)
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
			fmt.Println(err.Error())
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
		if err != nil {
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
			fmt.Println("Got a real error:", err.Error())
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
