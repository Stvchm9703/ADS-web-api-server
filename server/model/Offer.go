package model

import (
	"fmt"
	"time"
	"webserver/server/common"

	oid "github.com/coolbed/mgo-oid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type OfferMod struct {
	ID              *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	DepartmentID    *string        `bson:"dept_id,omitempty" json:"dept_id,omitempty"`
	CourseID        *string        `bson:"course_id,omitempty" json:"course_id,omitempty"`
	Year            *int           `bson:"year,omitempty" json:"year,omitempty"`
	ClassSize       *int           `bson:"class_size,omitempty" json:"class_size,omitempty"`
	AvailablePlaces *int           `bson:"available_places,omitempty" json:"available_places,omitempty"`
	CreatedAt       *time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt       *time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

var offer_mod_name = "Offer"

// FetchOffer : Get Offer list
func FetchOffer(param interface{}, ps *PageMeta) ([]*OfferMod, *PageMeta, error) {
	var record []*OfferMod
	nps := PageMeta{}
	fmt.Println("req. params", param)
	if DBConn != nil {
		count, err := DBConn.C(offer_mod_name).Find(&param).Count()
		if err != nil {
			fmt.Println("error")
			fmt.Println(err)
			fmt.Println("param")
			fmt.Println(param)
			return nil, nil, err
		}
		// fmt.Println("count:", count)
		Q := DBConn.C(offer_mod_name).Find(param)
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

// GetOffer : get one Offer object
func GetOffer(id string) (*OfferMod, error) {
	if DBConn != nil {
		var result *OfferMod
		err := DBConn.C(offer_mod_name).Find(bson.M{
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

// CreateOffer : Create a Offer Object
func CreateOffer(cp *OfferMod) (*OfferMod, error) {
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
		err := DBConn.C(offer_mod_name).Insert(&cp)
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
func UpdateOffer(Old *OfferMod, New *OfferMod) (*OfferMod, error) {
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
		Returned := OfferMod{}
		_, err := DBConn.C(offer_mod_name).Find(bson.M{"_id": Old.ID}).Apply(
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

// DeleteOffer : Delete a Offer
func DeleteOffer(cpid string) (bool, error) {
	if DBConn != nil {
		err := DBConn.C(offer_mod_name).Remove(&bson.M{"_id": bson.ObjectIdHex(cpid)})
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
func TestOffer(param map[string]interface{}) (bool, error) {
	if DBConn != nil {
		count, err := DBConn.C(offer_mod_name).Find(&param).Count()
		if err != nil {
			return false, err
		}
		return (count == 0), nil
	}
	_, err := NotConn()
	return false, err
}
