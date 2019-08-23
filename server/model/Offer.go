package model

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	oid "github.com/coolbed/mgo-oid"
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
	jsonf, _ := json.Marshal(param)
	var dfg map[string]interface{}
	var t = make(map[string]interface{})
	_ = json.Unmarshal(jsonf, &dfg)
	for k, v := range dfg {
		t["offers."+k] = v
	}
	fmt.Println("t", t)
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
				bson.M{"$match": t},
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
		err := DBConn.C(dept_mod_name).Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"courses._id":        bson.ObjectIdHex(courseId),
				"courses.offers._id": bson.ObjectIdHex(id),
			}},
			bson.M{"$unwind": "$courses"},
			bson.M{"$unwind": "$courses.offers"},
			bson.M{"$replaceRoot": bson.M{"newRoot": "$courses.offers"}},
			bson.M{"$match": bson.M{"_id": bson.ObjectIdHex(id)}},
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

		resultCursor := bson.M{}
		err := DBConn.Run(bson.M{
			"update": dept_mod_name,
			"updates": []bson.M{bson.M{
				"q": bson.M{
					"courses._id": bson.ObjectIdHex(courseId),
				},
				"u": bson.M{"$push": bson.M{
					"courses.$[ele].offers": &cp,
				}},
				"arrayFilters": []bson.M{
					bson.M{"ele._id": bson.M{"$eq": bson.ObjectIdHex(courseId)}},
				},
			}},
		}, &resultCursor)
		// return true, nil
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

	ojson, _ := json.Marshal(Old)
	// njson, _ := json.Marshal(New)
	// g := map[string]string{}
	// _ = json.Unmarshal(ojson, &g)
	fmt.Println("Old", string(ojson))

	if DBConn != nil {
		tnow := time.Now()
		New.UpdatedAt = &tnow
		if New.CreatedAt != Old.CreatedAt {
			New.CreatedAt = Old.CreatedAt
		}
		// did := bson.ObjectIdHex(g["_id"])
		New.ID = Old.ID

		fmt.Println("New", *New)
		cmdList := []bson.M{
			bson.M{
				"q": bson.M{
					"courses._id":        bson.ObjectIdHex(courseId),
					"courses.offers._id": Old.ID,
				},
				"u": bson.M{
					"$set": bson.M{"courses.$[ele].offers.$[elem]": New},
				},
				"arrayFilters": []bson.M{
					bson.M{"ele._id": bson.M{"$eq": bson.ObjectIdHex(courseId)}},
					bson.M{"elem._id": bson.M{"$eq": Old.ID}},
				},
			},
		}
		resultCursor := bson.M{}
		fmt.Println("DBConn", DBConn)
		err := DBConn.Run(bson.D{
			{"update", dept_mod_name},
			{"updates", &cmdList},
			{"ordered", true},
		}, &resultCursor)
		if err != nil {
			fmt.Print(err.Error())
			return nil, err
		}

		var result *OfferMod
		_ = DBConn.C(dept_mod_name).Pipe([]bson.M{
			bson.M{"$match": bson.M{
				"courses._id":        bson.ObjectIdHex(courseId),
				"courses.offers._id": Old.ID,
			}},
			bson.M{"$unwind": "$courses"},
			bson.M{"$unwind": "$courses.offers"},
			bson.M{"$replaceRoot": bson.M{"newRoot": "$courses.offers"}},
			bson.M{"$match": bson.M{"_id": Old.ID}},
		}).One(&result)
		fmt.Println(result)
		return result, nil
	}
	_, err := NotConn()
	return nil, err
}

// DeleteOffer : Delete a Offer
func DeleteOffer(courseId string, cpid string) (bool, error) {
	if DBConn != nil {

		resultCursor := bson.M{}
		cmdd := []bson.M{
			bson.M{
				"q": bson.M{
					"courses._id":        bson.ObjectIdHex(courseId),
					"courses.offers._id": bson.ObjectIdHex(cpid),
				},
				"u": bson.M{"$pull": bson.M{
					"courses.$[ele].offers": bson.M{"_id": bson.ObjectIdHex(cpid)},
				}},
				"arrayFilters": []bson.M{
					bson.M{"ele._id": bson.M{"$eq": bson.ObjectIdHex(courseId)}},
				},
			},
		}

		err := DBConn.Run(bson.D{
			{"update", dept_mod_name},
			{"updates", &cmdd},
		}, &resultCursor)
		fmt.Println(resultCursor)
		if err != nil || resultCursor["ok"] == 0 {
			log.Println("err in DeleteOffer", err, " resultCursor", resultCursor)
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
