package model

import (
	"fmt"
	"log"
	"time"
	"webserver/server/common"

	oid "github.com/coolbed/mgo-oid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EnrollMod struct {
	ID         *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	StudentID  *string        `bson:"student_id,omitempty" json:"student_id,omitempty"`
	Year       *int           `bson:"year,omitempty" json:"year,omitempty"`
	CourseID   *string        `bson:"course_id,omitempty" json:"course_id,omitempty"`
	EnrollDate *time.Time     `bson:"enroll_date,omitempty" json:"enroll_date,omitempty"`
	CreatedAt  *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

var enroll_mod_name = "Enroll"

// FetchEnroll : Get Enroll list
func FetchEnroll(param interface{}, ps *PageMeta) ([]*EnrollMod, *PageMeta, error) {
	var record []*EnrollMod
	nps := PageMeta{}
	fmt.Println("req. params", param)
	if DBConn != nil {
		count, err := DBConn.C(enroll_mod_name).Find(&param).Count()
		if err != nil {
			log.Fatalln("error")
			log.Fatalln(err)
			log.Fatalln("param")
			log.Fatalln(param)
			return nil, nil, err
		}
		// fmt.Println("count:", count)
		Q := DBConn.C(enroll_mod_name).Find(param)
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

// GetEnroll : get one Enroll object
func GetEnroll(id string) (*EnrollMod, error) {
	if DBConn != nil {
		var result *EnrollMod
		err := DBConn.C(enroll_mod_name).Find(bson.M{
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

// CreateEnroll : Create a Enroll Object
func CreateEnroll(cp *EnrollMod) (*EnrollMod, error) {
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
		err := DBConn.C(enroll_mod_name).Insert(&cp)
		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}
		return cp, nil
	}
	_, err := NotConn()
	return nil, err
}

// UpdateEnroll : Update a Enroll Object
func UpdateEnroll(Old *EnrollMod, New *EnrollMod) (*EnrollMod, error) {
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
		Returned := EnrollMod{}
		_, err := DBConn.C(enroll_mod_name).Find(bson.M{"_id": Old.ID}).Apply(
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

// DeleteEnroll : Delete a Enroll
func DeleteEnroll(cpid string) (bool, error) {
	if DBConn != nil {
		err := DBConn.C(enroll_mod_name).Remove(&bson.M{"_id": bson.ObjectIdHex(cpid)})
		if err != nil {
			log.Fatal("Got a real error:", err.Error())
			return false, err
		}
		return true, nil
	}
	_, err := NotConn()
	return false, err

}

// TestEnroll : Test Enroll is not existed
func TestEnroll(param map[string]interface{}) (bool, error) {
	if DBConn != nil {
		count, err := DBConn.C(enroll_mod_name).Find(&param).Count()
		if err != nil {
			return false, err
		}
		return (count == 0), nil
	}
	_, err := NotConn()
	return false, err
}
