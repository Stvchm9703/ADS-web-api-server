package viewMod

import (
	"encoding/json"
	"fmt"
	"time"
	"webserver/server/common"
	m "webserver/server/model"

	"gopkg.in/mgo.v2/bson"
)

type VCourseOfferMod struct {
	ID              *bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty" `
	// Department
	DeptObjID  *bson.ObjectId `bson:"dept_obj_id,omitempty" json:"dept_obj_id,omitempty"`
	DeptID          *string        `bson:"dept_id,omitempty" json:"dept_id,omitempty"`
	DeptName        *string        `bson:"dept_name,omitempty" json:"dept_name,omitempty"`
	Location          *string        `bson:"location,omitempty" json:"location,omitempty"`
	// Course
	CourseObjID 	*bson.ObjectId `bson:"course_obj_id,omitempty" json:"course_obj_id,omitempty"`
	CourseID        *string        `bson:"course_id,omitempty" json:"course_id,omitempty"`
	Title           *string        `bson:"title,omitempty" json:"title,omitempty"`
	Level           *int           `bson:"level,omitempty" json:"level,omitempty"`
	// Offer i
	CourseOfferId  *string    `bson:"course_offer_id,omitempty" json:"course_offer_id,omitempty"`
	Year            *int           `bson:"year,omitempty" json:"year,omitempty"`
	ClassSize       *int           `bson:"class_size,omitempty" json:"class_size,omitempty"`
	AvailablePlaces *int           `bson:"available_places,omitempty" json:"available_places,omitempty"`
	NumOfStud       *int           `bson:"num_of_stud,omitempty" json:"num_of_stud,omitempty"`
}

var v_course_offer_name = "VCourseOffer"
var v_course_offer_basemod = "department"
var v_course_offer_schema = []bson.M{
	bson.M{ "$unwind": "$courses" },
	bson.M{ "$unwind": "$courses.offers" },
	bson.M{
		"$project": bson.M{
            "_id": "$courses.offers._id",
            "dept_obj_id": "$_id",
            "dept_id": "$dept_id",
            "dept_name": "$dept_name",
            "location": "$location",
            "course_obj_id":"$courses._id",
            "course_id" : "$courses.course_id",
            "title": "$courses.title",
			"level": "$courses.level",
			"year" : "$courses.offers.year",
            "class_size": "$courses.offers.class_size",
			"available_place": "$courses.offers.available_place",
			"course_offer_id": bson.M{
				"$concat": []interface{}{ "$courses.course_id", "_",  bson.M{"$toStirng" : "$courses.offers.year"} },
			},
            "num_of_stud": bson.M{
                "$subtract": []string{ "$courses.offers.class_size","$courses.offers.available_places" },
			},	
        },
	},
}

// FetchVCourseOffer : Get VCourseOffer Object list
func FetchVCourseOffer(param interface{}, ps *m.PageMeta) ([]*VCourseOfferMod, *m.PageMeta, error) {
	var record []*VCourseOfferMod
	nps := m.PageMeta{}
	// fmt.Println("req. params", param)
	if m.DBConn != nil {
		count, err := m.DBConn.C(v_course_offer_name).Find(&param).Count()
		if err != nil {
			return nil, nil, err
		}
		// fmt.Println("count:", count)
		Q := m.DBConn.C(v_course_offer_name).Find(param)
		if len( ps.SortAr ) > 0 {
			Q = Q.Sort(ps.SortAr...)
			nps.Sort = ps.Sort
			nps.SortAr = ps.SortAr
		} 
		if ps.PageLimit > 0 {
			Q = Q.Limit(ps.PageLimit)
			nps.PageLimit = ps.PageLimit
		} else {
			Q = Q.Limit(common.QueryDefaultPageLimit)
			nps.PageLimit = common.QueryDefaultPageLimit // default Page Limit
		}
		// fmt.Println("req. pageNum", ps.PageNum)
		// defAULT : 1
		if ps.PageNum > 0 {
			Q = Q.Skip((ps.PageNum - 1) * ps.PageLimit)
			nps.PageNum = ps.PageNum
		} else {
			nps.PageNum = 1
		}
		// fmt.Println("Q:", Q)
		err1 := Q.All(&record)
		if err1 != nil {
			return nil, nil, err1
		}
		nps.Count = count
		// fmt.Println(record)
		return record, &nps, nil
	}
	_, err := m.NotConn()
	return nil, nil, err

}

// GetVCourseOffer : get one VCourseOffer object
func GetVCourseOffer(id string) (*VCourseOfferMod, error) {
	if m.DBConn != nil {
		var result *VCourseOfferMod
		err := m.DBConn.C(v_course_offer_name).Find(bson.M{
			"_id": bson.ObjectIdHex(id),
		}).One(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	_, err := m.NotConn()
	return nil, err
}

// TestVCourseOffer : Test VCourseOffer is not existed
func TestVCourseOffer(param map[string]interface{}) (bool, error) {
	if m.DBConn != nil {
		count, err := m.DBConn.C(v_course_offer_name).Find(&param).Count()
		if err != nil {
			return false, err
		}
		return (count == 0), nil
	}
	_, err := m.NotConn()
	return false, err
}

func CreateSchemaVCourseOffer() (bool, error) {
	bsonCheckName := []string{}
	// for ks := 0; ks < len(v_course_offer_schema); ks++ {
	// 	tmp := v_course_offer_schema[ks]
	// 	if tmp["$lookup"] != nil {
	// 		tmp1 := tmp["$lookup"].(bson.M)
	// 		bsonCheckName = append(bsonCheckName, tmp1["from"].(string))
	// 	}
	// }
	// bsonCheckName = append(bsonCheckName, v_course_offer_basemod)
	// fmt.Println("bsonCheckName: ", bsonCheckName)
	nameList := m.MgoCursorRes{}
	m.DBConn.Run(bson.M{
		"listCollections": 1.0,
		"nameOnly":        true,
	}, &nameList)
	fmt.Println("nameList", nameList.Cursor.FirstBatch)

	list := nameList.Cursor.FirstBatch
I:
	for i := 0; i < len(list); i++ {
		tmp := list[i]["name"].(string)
		if tmp == v_course_offer_name {
			return false, nil
			break I
		// } else {
		// 	for j := 0; j < len(bsonCheckName); j++ {
		// 		if bsonCheckName[j] == tmp {
		// 			bsonCheckName = append(bsonCheckName[:j], bsonCheckName[j+1:]...)
		// 			continue I
		// 		}
		// 	}
		}
	}
	fmt.Println("bsonCheckName:", bsonCheckName, ", :num:", len(bsonCheckName))
	if len(bsonCheckName) > 0 {
		// fmt.Println("failed , some of Collection is not declared")
		nameJ, _ := json.Marshal(bsonCheckName)
		return false, common.ErrorMessage{
			When: time.Now(),
			What: "failed , some of Collection is not declared : " + string(nameJ),
		}
	}
	resultCursor := m.MgoCursorRes{}
	m.DBConn.Run(bson.M{
		"create":   v_course_offer_name,
		"viewOn":   v_course_offer_basemod,
		"pipeline": v_course_offer_schema,
	}, resultCursor)
	if resultCursor.OK == 1 {
		return true, nil
	}
	nameJ, _ := json.Marshal(resultCursor)
	return false, common.ErrorMessage{
		When: time.Now(),
		What: string(nameJ),
	}
	return false, nil
}
