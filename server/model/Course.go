package model

import (
	"fmt"

	// "log"
	"time"

	"gopkg.in/mgo.v2/bson"

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
type CourseListM struct {
	ID       *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	DeptID   *string        `bson:"dept_id,omitempty" json:"dept_id,omitempty"`
	DeptName *string        `bson:"dept_name,omitempty" json:"dept_name,omitempty"`
	Location *string        `bson:"location,omitempty" json:"location,omitempty"`
	Courses  *CourseMod     `bson:"courses,omitempty" json:"courses,omitempty"`
}

// var course_mod_name = "Course"

// FetchCourse : GEt the Course list
func FetchDeptCourse(dept_id string, param interface{}, ps *PageMeta) ([]*CourseMod, *PageMeta, error) {
	var record *DepartmentMod
	nps := PageMeta{}
	fmt.Println("req. params", param)
	if DBConn != nil {

		err1 := DBConn.C(dept_mod_name).Find(
			bson.M{
				"_id": bson.ObjectIdHex(dept_id),
				"courses": bson.M{
					"$elemMatch": param,
				},
			},
		).One(&record)
		if err1 != nil {
			fmt.Println(err1)
			return nil, nil, err1
		}
		// nps.Count = count
		fmt.Println("record", record)
		return record.Courses, &nps, nil
	}
	_, err := NotConn()
	return nil, nil, err

}

// FetchCourse : GEt the Course list
func FetchAllCourse(param interface{}, ps *PageMeta) ([]CourseListM, *PageMeta, error) {
	var record []CourseListM
	nps := PageMeta{}
	fmt.Println("req. params", param)
	if DBConn != nil {
		qry := []bson.M{
			bson.M{"$match": bson.M{
				"courses": bson.M{
					"$elemMatch": param,
				},
			}},
			bson.M{"$unwind": "$courses"},
		}
		if len(ps.Sort) > 0 {
			qry = append(qry, bson.M{"$sort": ps.Sort})
			nps.Sort = ps.Sort
		}
		if ps.PageLimit > 0 {
			qry = append(qry, bson.M{"$sort": ps.PageLimit})
			nps.PageLimit = ps.PageLimit
		}
		if ps.PageNum > 0 {
			qry = append(qry, bson.M{"$skip": (ps.PageNum - 1) * ps.PageLimit})
			nps.PageNum = ps.PageNum
		}
		err1 := DBConn.C(dept_mod_name).Pipe(qry).All(&record)
		if err1 != nil {
			fmt.Println(err1)
			return nil, nil, err1
		}

		return record, &nps, nil
	}
	_, err := NotConn()
	return nil, nil, err

}

// GetCourse : get one Course object
func GetDeptCourse(dept_id string, id string) (*CourseMod, error) {
	if DBConn != nil {
		var result *CourseListM

		qry := []bson.M{
			bson.M{"$match": bson.M{
				"_id":         bson.ObjectIdHex(dept_id),
				"courses._id": bson.ObjectIdHex(id),
			}},
			bson.M{"$unwind": "$courses"},
			bson.M{"$match": bson.M{
				"_id":         bson.ObjectIdHex(dept_id),
				"courses._id": bson.ObjectIdHex(id),
			}},
		}
		err1 := DBConn.C(dept_mod_name).Pipe(qry).One(&result)
		if err1 != nil {
			fmt.Println(err1)
			return nil, err1
		}
		return result.Courses, nil
	}
	_, err := NotConn()
	return nil, err
}

// CreateCourse : Create a Course Object
func CreateCourse(deptId string, cp *CourseMod) (*CourseMod, error) {
	if DBConn != nil {
		tnow := time.Now()
		if cp.ID == nil {
			temID := bson.ObjectId(string(oid.NewOID().Bytes()))
			cp.ID = &temID
			fmt.Println(cp.ID)
		}
		cp.CreatedAt = &tnow
		cp.UpdatedAt = &tnow
		cp.Offers = []*OfferMod{}
		err := DBConn.C(dept_mod_name).Update(bson.M{
			"_id": bson.ObjectIdHex(deptId),
		}, bson.M{
			"$push": bson.M{
				"courses": &cp,
			},
		})
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
func UpdateCourse(dept_id string, Old *CourseMod, New *CourseMod) (*CourseMod, error) {
	if DBConn != nil {
		tnow := time.Now()
		New.UpdatedAt = &tnow
		if New.CreatedAt != Old.CreatedAt {
			New.CreatedAt = Old.CreatedAt
		}
		temp, _ := bson.Marshal(New)
		upNew := bson.M{}
		bson.Unmarshal(temp, upNew)
		delete(upNew, "Offers")
		upNew["_id"] = Old.ID
		fmt.Println("update upNew", upNew)

		resultCursor := bson.M{}
		err := DBConn.Run(bson.M{
			"update": dept_mod_name,
			"updates": []bson.M{bson.M{
				"q": bson.M{
					"_id":         bson.ObjectIdHex(dept_id),
					"courses._id": Old.ID,
				},
				"u": bson.M{"$set": bson.M{
					"courses.$[ele]": upNew,
				}},
				"arrayFilters": []bson.M{
					bson.M{"ele._id": bson.M{"$eq": Old.ID}},
				},
			}},
		}, &resultCursor)
		if err != nil {
			return nil, err
		}

		var result *CourseListM

		qry := []bson.M{
			bson.M{"$match": bson.M{
				"_id":         bson.ObjectIdHex(dept_id),
				"courses._id": Old.ID,
			}},
			bson.M{"$unwind": "$courses"},
			bson.M{"$match": bson.M{
				"_id":         bson.ObjectIdHex(dept_id),
				"courses._id": Old.ID,
			}},
		}
		err = DBConn.C(dept_mod_name).Pipe(qry).One(&result)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return result.Courses, nil
		// return GetDeptCourse(dept_id, Old.ID.Hex())
	}
	_, err := NotConn()
	return nil, err
}

// DeleteCourse : Delete a Course
func DeleteCourse(dept_id string, cpid string) (bool, error) {
	if DBConn != nil {
		// err := DBConn.C(course_mod_name).Remove(&bson.M{"_id": bson.ObjectIdHex(cpid)})
		err := DBConn.C(dept_mod_name).Update(bson.M{
			"_id": bson.ObjectIdHex(dept_id),
		}, bson.M{
			"$pull": bson.M{
				"courses": bson.M{"_id": bson.ObjectIdHex(cpid)},
			},
		})

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
func TestCourse(dept_id string, param map[string]interface{}) (bool, error) {
	if DBConn != nil {
		count, err := DBConn.C(dept_mod_name).Find(
			bson.M{
				"_id": bson.ObjectIdHex(dept_id),
				"courses": bson.M{
					"$elemMatch": param,
				},
			},
		).Count()
		if err != nil {
			return false, err
		}
		return (count == 0), nil
	}
	_, err := NotConn()
	return false, err
}
