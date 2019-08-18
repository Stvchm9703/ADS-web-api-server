package model

import (
	"fmt"
	"time"

	oid "github.com/coolbed/mgo-oid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type OfferMod struct {
	ID              *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Year            *int           `bson:"year,omitempty" json:"year,omitempty"`
	ClassSize       *int           `bson:"class_size,omitempty" json:"class_size,omitempty"`
	AvailablePlaces *int           `bson:"available_places,omitempty" json:"available_places,omitempty"`
	CreatedAt       *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt       *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
type CourseOfferMod struct {
	ID       *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	CourseID *string        `bson:"course_id,omitempty" json:"course_id,omitempty"`
	Title    *string        `bson:"title,omitempty" json:"title,omitempty"`
	Level    *int           `bson:"level,omitempty" json:"level,omitempty"`
	Offers   *struct {
		ID              *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
		Year            *int           `bson:"year,omitempty" json:"year,omitempty"`
		ClassSize       *int           `bson:"class_size,omitempty" json:"class_size,omitempty"`
		AvailablePlaces *int           `bson:"available_places,omitempty" json:"available_places,omitempty"`
		CreatedAt       *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
		UpdatedAt       *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	} `bson:"offers,omitempty" json:"offers,omitempty"`
	CreatedAt *time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

var offer_mod_name = "Offer"

// FetchCourse : GEt the Course list by Course id
func FetchCourseOffer(courseId string, param interface{}, ps *PageMeta) ([]*CourseMod, *PageMeta, error) {
	var record []*CourseMod
	nps := PageMeta{}
	fmt.Println("req. params", param)
	if DBConn != nil {
		err1 := DBConn.C(dept_mod_name).Pipe(
			[]bson.M{
				bson.M{"$match": bson.M{
					"courses._id": bson.ObjectIdHex(courseId),
					"courses.offers": bson.M{
						"$elemMatch": param,
					},
				}},
				bson.M{"$unwind": "$courses"},
				// bson.M{"$unwind": "$courses.offers"},
				bson.M{"$replaceRoot": bson.M{"newRoot": "$courses"}},
			},
		).All(&record)
		if err1 != nil {
			fmt.Println(err1)
			return nil, nil, err1
		}
		// nps.Count = count
		fmt.Println("record", record)
		return record, &nps, nil
	}
	_, err := NotConn()
	return nil, nil, err

}

// FetchAllOffer : Globle search in all Offer
func FetchAllOffer(param interface{}, ps *PageMeta) ([]*CourseOfferMod, *PageMeta, error) {
	var record []*CourseOfferMod
	nps := PageMeta{}
	fmt.Println("req. params", param)
	if DBConn != nil {
		err1 := DBConn.C(dept_mod_name).Pipe(
			[]bson.M{
				bson.M{"$match": bson.M{
					"courses.offers": bson.M{
						"$elemMatch": param,
					},
				}},
				bson.M{"$unwind": "$courses"},
				bson.M{"$unwind": "$courses.offers"},
				bson.M{"$replaceRoot": bson.M{"newRoot": "$courses"}},
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

// GetOffer : get one Offer object
func GetOffer(courseId string, id string) (*OfferMod, error) {
	if DBConn != nil {
		var result *OfferMod
		err := DBConn.C(dept_mod_name).Find(bson.M{
			"courses._id":        bson.ObjectIdHex(courseId),
			"courses.offers._id": bson.ObjectIdHex(id),
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

// CreateOffer : Create a Offer Object
func CreateOffer(courseId string, cp *OfferMod) (*OfferMod, error) {
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
		err := DBConn.C(dept_mod_name).Update(bson.M{
			"courses._id": bson.ObjectIdHex(courseId),
		}, bson.M{
			"$push": bson.M{
				"courses.$.offers": &cp,
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

// UpdateOffer : Update a Offer Object
func UpdateOffer(courseId string, Old *OfferMod, New *OfferMod) (*OfferMod, error) {
	if DBConn != nil {
		tnow := time.Now()
		New.UpdatedAt = &tnow
		if New.CreatedAt != Old.CreatedAt {
			New.CreatedAt = Old.CreatedAt
		}

		temp, _ := bson.Marshal(New)
		upNew := bson.M{}
		bson.Unmarshal(temp, upNew)
		Returned := OfferMod{}
		_, err := DBConn.C(offer_mod_name).Find(bson.M{
			"courses._id":        bson.ObjectIdHex(courseId),
			"courses.offers._id": Old.ID,
		}).Apply(
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

// DeleteOffer : Delete a Offer
func DeleteOffer(courseId string, cpid string) (bool, error) {
	if DBConn != nil {
		err := DBConn.C(dept_mod_name).Update(&bson.M{
			"courses._id": bson.ObjectIdHex(courseId),
		}, bson.M{
			"$pull": bson.M{
				"courses.$.offers._id": bson.ObjectIdHex(cpid),
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

// TestOffer : Test Offer is not existed
func TestOffer(id string, param map[string]interface{}) (bool, error) {
	if DBConn != nil {
		count, err := DBConn.C(dept_mod_name).Find(
			bson.M{"courses": bson.M{
				"$elemMatch": bson.M{
					"_id":    bson.ObjectIdHex(id),
					"offers": bson.M{"$elemMatch": param},
				},
			}},
		).Count()
		if err != nil {
			return false, err
		}
		return (count == 0), nil
	}
	_, err := NotConn()
	return false, err
}
