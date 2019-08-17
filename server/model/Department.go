package model

import (
	// "fmt"

	"fmt"
	"time"
	"webserver/server/common"

	oid "github.com/coolbed/mgo-oid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DepartmentMod struct {
	ID        *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	DeptID    *string        `bson:"dept_id,omitempty" json:"dept_id,omitempty"`
	DeptName  *string        `bson:"dept_name,omitempty" json:"dept_name,omitempty"`
	Location  *string        `bson:"location,omitempty" json:"location,omitempty"`
	Courses   []*CourseMod   `bson:"courses,omitempty" json:"courses,omitempty"`
	CreatedAt *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type DepartmentListM struct {
	ID       *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	DeptID   *string        `bson:"dept_id,omitempty" json:"dept_id,omitempty"`
	DeptName *string        `bson:"dept_name,omitempty" json:"dept_name,omitempty"`
	Location *string        `bson:"location,omitempty" json:"location,omitempty"`
	Courses  []struct {
		ID        *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
		CourseID  *string        `bson:"course_id,omitempty" json:"course_id,omitempty"`
		Title     *string        `bson:"title,omitempty" json:"title,omitempty"`
		Level     *int           `bson:"level,omitempty" json:"level,omitempty"`
		CreatedAt *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
		UpdatedAt *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	} `bson:"courses,omitempty" json:"courses,omitempty"`
	CreatedAt *time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// dept_mod_name : model name
var dept_mod_name = "Department"

// FetchDepartment : Get Department Object list
func FetchDepartment(param interface{}, ps *PageMeta) ([]*DepartmentListM, *PageMeta, error) {
	var record []*DepartmentListM
	nps := PageMeta{}
	if DBConn != nil {
		count, err := DBConn.C(dept_mod_name).Find(&param).Count()

		if err != nil {
			return nil, nil, err
		}
		Q := DBConn.C(dept_mod_name).Find(param)
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
		return record, &nps, nil
	}
	_, err := NotConn()
	return nil, nil, err

}

// GetDepartment : get one Department object
func GetDepartment(id string) (*DepartmentMod, error) {
	if DBConn != nil {
		var result *DepartmentMod
		err := DBConn.C(dept_mod_name).Find(bson.M{
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

// CreateDepartment : Create a Department Object
func CreateDepartment(cp *DepartmentMod) (*DepartmentMod, error) {
	if DBConn != nil {
		tnow := time.Now()
		if cp.ID == nil {
			temID := bson.ObjectId(string(oid.NewOID().Bytes()))
			cp.ID = &temID
		}
		cp.CreatedAt = &tnow
		cp.UpdatedAt = &tnow
		err := DBConn.C(dept_mod_name).Insert(&cp)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		return cp, nil
	}
	_, err := NotConn()
	return nil, err
}

// UpdateDepartment : Update a Department Object
func UpdateDepartment(Old *DepartmentMod, New *DepartmentMod) (*DepartmentMod, error) {
	if DBConn != nil {
		tnow := time.Now()
		New.UpdatedAt = &tnow
		if New.CreatedAt != Old.CreatedAt {
			New.CreatedAt = Old.CreatedAt
		}
		if New.Courses == nil || len(New.Courses) != len(Old.Courses) {
			New.Courses = Old.Courses
		}
		temp, _ := bson.Marshal(New)
		upNew := bson.M{}
		bson.Unmarshal(temp, upNew)
		Returned := DepartmentMod{}
		_, err := DBConn.C(dept_mod_name).Find(bson.M{"_id": Old.ID}).Apply(
			mgo.Change{
				Update:    bson.M{"$set": upNew},
				ReturnNew: true,
			},
			&Returned,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return &Returned, nil
	}
	_, err := NotConn()
	return nil, err
}

// DeleteDepartment : Delete a Department
func DeleteDepartment(cpid string) (bool, error) {
	if DBConn != nil {
		err := DBConn.C(dept_mod_name).Remove(&bson.M{"_id": bson.ObjectIdHex(cpid)})
		if err != nil {
			fmt.Println("Got a real error:", err.Error())
			return false, err
		}
		return true, nil
	}
	_, err := NotConn()
	return false, err

}

// TestDepartment : Test Department is not existed
func TestDepartment(param map[string]interface{}) (bool, error) {
	if DBConn != nil {
		count, err := DBConn.C(dept_mod_name).Find(&param).Count()
		if err != nil {
			return false, err
		}
		return (count == 0), nil
	}
	_, err := NotConn()
	return false, err
}
