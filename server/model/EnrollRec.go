package model

import (
	"encoding/json"
	"fmt"
	"time"

	oid "github.com/coolbed/mgo-oid"
	"gopkg.in/mgo.v2/bson"
)

type EnrollMod struct {
	ID         *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Year       *int           `bson:"year,omitempty" json:"year,omitempty"`
	CourseID   *string        `bson:"course_id,omitempty" json:"course_id,omitempty"`
	EnrollDate *time.Time     `bson:"enroll_date,omitempty" json:"enroll_date,omitempty"`
	CreatedAt  *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type StudentEnrollMod struct {
	ID        *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	StudentID *string        `bson:"student_id,omitempty" json:"student_id,omitempty"`
	StuName   *string        `bson:"student_name,omitempty" json:"student_name,omitempty"`
	DOB       *time.Time     `bson:"dob,omitempty" json:"dob,omitempty"`
	Enrolled  *EnrollMod     `bson:"enrolled,omitempty" json:"enrolled,omitempty"`
	CreatedAt *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// FetchEnroll : Get Enroll list
func FetchStudEnroll(sid string, param interface{}, ps *PageMeta) ([]*EnrollMod, *PageMeta, error) {
	var record *StudentMod
	nps := PageMeta{}
	fmt.Println("req. params", param)
	if DBConn != nil {
		err := DBConn.C(student_mod_name).Find(bson.M{
			"_id": bson.ObjectIdHex(sid),
			"enrolled": bson.M{
				"$elemMatch": param,
			},
		}).One(&record)
		if err != nil {
			return nil, nil, err
		}
		fmt.Println(record)
		return record.Enrolled, &nps, nil
	}
	_, err := NotConn()
	return nil, nil, err
}
func FetchAllEnroll(param interface{}, ps *PageMeta) ([]*StudentEnrollMod, *PageMeta, error) {
	var record []*StudentEnrollMod
	nps := PageMeta{}
	fmt.Println("req. params", param)
	if DBConn != nil {
		qry :=[]bson.M{
			bson.M{"$match": bson.M{
				"enrolled": bson.M{
					"$elemMatch": param,
				},
			}},
			bson.M{"$unwind": "$enrolled"},
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
			qry = append(qry, bson.M{"$skip" : (ps.PageNum-1) *ps.PageLimit})
			nps.PageNum = ps.PageNum
		}
		err1 := DBConn.C(student_mod_name).Pipe(qry).All(&record)
		if err1 != nil {
			return nil, nil, err1
		}
		return record, &nps, nil
	}
	_, err := NotConn()
	return nil, nil, err
}

// GetEnroll : get one Enroll object
func GetEnroll(sid string, id string) (*EnrollMod, error) {
	if DBConn != nil {
		var result *StudentEnrollMod
		// err := DBConn.C(student_mod_name).Find(bson.M{
		// 	"_id":          bson.ObjectIdHex(sid),
		// 	"enrolled._id": bson.ObjectIdHex(id),
		// }).One(&result)
		err := DBConn.C(student_mod_name).Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"_id": bson.ObjectIdHex(sid),
				"enrolled": bson.M{
					"$elemMatch": bson.M{"_id": bson.ObjectIdHex(id)},
				},
			}},
			bson.M{"$unwind": "$enrolled"},
			bson.M{"$match": bson.M{
				"_id":          bson.ObjectIdHex(sid),
				"enrolled._id": bson.ObjectIdHex(id),
			}},
		}).One(&result)
		if err != nil {
			return nil, err
		}
		return result.Enrolled, nil
	}
	_, err := NotConn()
	return nil, err
}

// CreateEnroll : Create a Enroll Object
func CreateEnroll(sid string, cp *EnrollMod) (*EnrollMod, error) {
	if DBConn != nil {
		tnow := time.Now()
		if cp.ID == nil {
			temID := bson.ObjectId(string(oid.NewOID().Bytes()))
			fmt.Println("temId:", temID.String())
			cp.ID = &temID
			fmt.Println(cp.ID)
		}
		cp.CreatedAt = &tnow
		cp.UpdatedAt = &tnow

		// session

		resultCursor := MgoCursorRes{}
		err := DBConn.Run(bson.M{
			"update": dept_mod_name,
			"updates": []bson.M{bson.M{
				"q": bson.M{
					"courses.course_id":               cp.CourseID,
					"courses.offers.year":             cp.Year,
					"courses.offers.available_places": bson.M{"$gt": 0},
				},
				"u": bson.M{
					"$inc": bson.M{"courses.$[ele].offers.$[elem].available_places": -1},
				},
				"arrayFilters": []bson.M{
					bson.M{"ele.course_id": bson.M{"$eq": cp.CourseID}},
					bson.M{"elem.year": bson.M{"$eq": cp.Year}},
				},
			}},
		}, &resultCursor)
		if err != nil {
			return nil, err
		}

		err = DBConn.C(student_mod_name).Update(bson.M{
			"_id": bson.ObjectIdHex(sid),
		}, bson.M{
			"$push": bson.M{"enrolled": &cp},
		})
		if err != nil {
			return nil, err
		}
		return cp, nil
	}
	_, err := NotConn()
	return nil, err
}

// UpdateEnroll : Update a Enroll Object
func UpdateEnroll(stud_id string, Old *EnrollMod, New *EnrollMod) (*EnrollMod, error) {
	if DBConn != nil {
		tnow := time.Now()
		New.UpdatedAt = &tnow
		if New.CreatedAt != Old.CreatedAt {
			New.CreatedAt = Old.CreatedAt
		}
		tempO, _ := json.Marshal(New)
		fmt.Println(string(tempO))
		temp, _ := bson.Marshal(New)
		upNew := bson.M{}
		bson.Unmarshal(temp, upNew)
		upNew["_id"] = Old.ID
		fmt.Println(upNew)

		resultCursor := bson.M{}
		err := DBConn.Run(bson.M{
			"update": student_mod_name,
			"updates": []bson.M{bson.M{
				"q": bson.M{
					"_id":          bson.ObjectIdHex(stud_id),
					"enrolled._id": Old.ID,
				},
				"u": bson.M{"$set": bson.M{
					"enrolled.$[ele]": upNew,
				}},
				"arrayFilters": []bson.M{
					bson.M{"ele._id": bson.M{"$eq": Old.ID}},
				},
			}},
		}, &resultCursor)
		// fmt.Println()
		// fmt.Println("resultCursor:", resultCursor)
		// fmt.Println()

		if err != nil {
			return nil, err
		}
		Returned := StudentEnrollMod{}
		err = DBConn.C(student_mod_name).Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"_id": bson.ObjectIdHex(stud_id),
				"enrolled": bson.M{
					"$elemMatch": bson.M{"_id": Old.ID},
				},
			}},
			bson.M{"$unwind": "$enrolled"},
			bson.M{"$match": bson.M{
				"_id":          bson.ObjectIdHex(stud_id),
				"enrolled._id": Old.ID,
			}},
		}).One(&Returned)

		if err != nil {
			return nil, err
		}
		return Returned.Enrolled, nil
	}
	_, err := NotConn()
	return nil, err
}

// DeleteEnroll : Delete a Enroll
func DeleteEnroll(sid string, cpid string) (bool, error) {
	if DBConn != nil {
		err := DBConn.C(student_mod_name).Update(bson.M{
			"_id": bson.ObjectIdHex(sid),
		}, bson.M{
			"$pull": bson.M{
				"enrolled": bson.M{"_id": bson.ObjectIdHex(cpid)},
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

// TestEnroll : Test Enroll is not existed
func TestEnroll(sid string, param map[string]interface{}) (bool, error) {
	if DBConn != nil {
		count, err := DBConn.C(student_mod_name).Find(bson.M{
			"_id": bson.ObjectIdHex(sid),
			"enrolled": bson.M{
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
