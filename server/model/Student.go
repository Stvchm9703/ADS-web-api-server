package model

import (
	"fmt"
	"time"
	"webserver/server/common"

	oid "github.com/coolbed/mgo-oid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type StudentMod struct {
	ID        *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	StudentID *string        `bson:"student_id,omitempty" json:"student_id,omitempty"`
	StuName   *string        `bson:"student_name,omitempty" json:"student_name,omitempty"`
	DOB       *time.Time     `bson:"dob,omitempty" json:"dob,omitempty"`
	Enrolled  []*EnrollMod   `bson:"enrolled,omitempty" json:"enrolled,omitempty"`
	CreatedAt *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

var student_mod_name = "Student"

// FetchStudent : Get Student list
func FetchStudent(param interface{}, ps *PageMeta) ([]*StudentMod, *PageMeta, error) {
	var record []*StudentMod
	nps := PageMeta{}
	fmt.Println("req. params", param)
	if DBConn != nil {
		count, err := DBConn.C(student_mod_name).Find(&param).Count()
		if err != nil {

			return nil, nil, err
		}
		// fmt.Println("count:", count)
		Q := DBConn.C(student_mod_name).Find(param)
		if len( ps.SortAr ) > 0 {
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
			fmt.Println("error")
			fmt.Println(err1)
			fmt.Println("param")
			fmt.Println(param)
			return nil, nil, err1
		}
		nps.Count = count
		fmt.Println(record)
		return record, &nps, nil
	}
	_, err := NotConn()
	return nil, nil, err

}

// GetStudent : get one Student object
func GetStudent(id string) (*StudentMod, error) {
	if DBConn != nil {
		var result *StudentMod
		err := DBConn.C(student_mod_name).Find(bson.M{
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

// CreateStudent : Create a Student Object
func CreateStudent(cp *StudentMod) (*StudentMod, error) {
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
		err := DBConn.C(student_mod_name).Insert(&cp)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		return cp, nil
	}
	_, err := NotConn()
	return nil, err
}

// UpdateStudent : Update a Student Object
func UpdateStudent(Old *StudentMod, New *StudentMod) (*StudentMod, error) {
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
		Returned := StudentMod{}
		_, err := DBConn.C(student_mod_name).Find(bson.M{"_id": Old.ID}).Apply(
			mgo.Change{
				Update:    bson.M{"$set": upNew},
				ReturnNew: true,
			},
			&Returned,
		)
		// fmt.Println("info", info)
		// fmt.Println("Returned", Returned)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return &Returned, nil
	}
	_, err := NotConn()
	return nil, err
}

// DeleteStudent : Delete a Student
func DeleteStudent(cpid string) (bool, error) {
	if DBConn != nil {
		err := DBConn.C(student_mod_name).Remove(&bson.M{"_id": bson.ObjectIdHex(cpid)})
		if err != nil {
			fmt.Println("Got a real error:", err.Error())
			return false, err
		}
		return true, nil
	}
	_, err := NotConn()
	return false, err

}

// TestStudent : Test Student is not existed
func TestStudent(param map[string]interface{}) (bool, error) {
	if DBConn != nil {
		count, err := DBConn.C(student_mod_name).Find(&param).Count()
		if err != nil {
			return false, err
		}
		return (count == 0), nil
	}
	_, err := NotConn()
	return false, err
}
