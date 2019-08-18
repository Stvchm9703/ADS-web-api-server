package model

import (
	"fmt"
	"time"

	oid "github.com/coolbed/mgo-oid"
	"gopkg.in/mgo.v2"
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

var student_mod_name = "Enroll"

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
		err1 := DBConn.C(student_mod_name).Pipe(
			[]bson.M{
				bson.M{"$match": bson.M{
					"enrolled": bson.M{
						"$elemMatch": param,
					},
				}},
				bson.M{"$unwind": "$enrolled"},
			},
		).All(&record)
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
		err := DBConn.C(student_mod_name).Find(bson.M{
			"_id":          bson.ObjectIdHex(sid),
			"enrolled._id": bson.ObjectIdHex(id),
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
		fmt.Println("hi create")
		if cp.ID == nil {
			temID := bson.ObjectId(string(oid.NewOID().Bytes()))
			fmt.Println("temId:", temID.String())
			cp.ID = &temID
			fmt.Println(cp.ID)
		}
		cp.CreatedAt = &tnow
		cp.UpdatedAt = &tnow
		err := DBConn.C(student_mod_name).Update(bson.M{
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

		temp, _ := bson.Marshal(New)
		upNew := bson.M{}
		bson.Unmarshal(temp, upNew)
		Returned := EnrollMod{}
		_, err := DBConn.C(student_mod_name).Find(
			bson.M{
				"_id": bson.ObjectIdHex(stud_id),
			},
		).Apply(
			mgo.Change{
				Update: bson.M{
					"$set": bson.M{"$$": upNew}},
				ReturnNew: true,
			},
			&Returned,
		)
		// fmt.Println("info", info)
		// fmt.Println("Returned", Returned)
		if err != nil {
			return nil, err
		}
		return &Returned, nil
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
				"enrolled": bson.M{
					"_id": bson.ObjectIdHex(cpid),
				},
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
