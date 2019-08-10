package model

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
	"webserver/server/common"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PageMeta struct {
	PageLimit int `json:"page_limit,omitempty"`
	PageNum   int `json:"page_num,omitempty"`
	Count     int `json:"count,omitempty"`
}

var DBSess *mgo.Session
var DBConn *mgo.Database

func ConnectDB(temp *common.ConfigTemp) (*mgo.Session, error) {
	// CreateDBTable(temp)
	Host := []string{
		temp.Database.Host + ":" + strconv.Itoa(temp.Database.Port),
	}
	DBSess, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		Username: temp.Database.Username,
		Password: temp.Database.Password,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	DBConn = DBSess.DB(temp.Database.Database)

	CreateDBTable(temp)
	return DBSess, nil
}

func DisconnDB() (bool, error) {
	DBSess.Close()
	return true, nil
}

func NotConn() (*struct{}, error) {
	fmt.Println(DBSess)
	fmt.Println(DBConn)
	return nil, &common.ErrorMessage{
		When: time.Now(),
		What: "MongoDB is not Connect",
	}
}

func CreateDBTable(config *common.ConfigTemp) {
	var structlist = []interface{}{
		CourseMod{},
		DepartmentMod{},
		EnrollMod{},
		LocationMod{},
		OfferMod{},
		StudentMod{},
	}
	tmpBD := []bson.M{}
	for k, v := range structlist {
		fmt.Println("k", k)
		tmpProp := bson.M{}
		// temp := map[string]interface{}{}
		r := reflect.TypeOf(v)
		for ri := 0; ri < r.NumField(); ri++ {
			smp := strings.Split(r.Field(ri).Tag.Get("bson"), ",")[0]
			tmpProp[smp] = bson.M{
				"bsonType":    typeConv(r.Field(ri)),
				"description": "Go System AutoGen : " + r.Field(ri).Name,
			}
			fmt.Println("tmpProp[smp]:")
			fmt.Println(tmpProp[smp])
		}
		fmt.Println("tmpProp:")
		fmt.Println(tmpProp)
		mod_name := strings.Replace(r.Name(), "Mod", "", -1)
		tmp_schema := bson.M{
			"create":           mod_name,
			"validationAction": "warn",
			"validator": bson.M{
				"$jsonSchema": bson.M{
					"bsonType":   "object",
					"properties": tmpProp,
				},
			},
		}
		tmpBD = append(tmpBD, tmp_schema)
	}
	fmt.Println("structlist", structlist)
	spew.Printf("runcmd : %v\n", tmpBD)
	nameList := bson.M{}
	DBConn.Run(bson.M{
		"listCollections":       1.0,
		"authorizedCollections": true,
		"nameOnly":              true,
	}, &nameList)
	spew.Printf("nameList: %v\n", nameList["cursor"])
}

func typeConv(t reflect.StructField) string {
	targ := strings.Replace(t.Type.String(), "*", "", -1)

	if strings.Contains(targ, "[]") && !strings.Contains(targ, "byte") {
		return "array"
	}
	switch targ {
	case "int", "int32", "int64", "uint32", "uint64", "uint":
		return "int"
	case "time.Time":
		return "date"
	case "float32", "float64":
		return "double"
	case "interface{}", "struct":
		return "object"
	case "[]byte", "byte":
		return "binData"
	case "bool":
		return "bool"
	case "bson.ObjectId":
		return "objectId"
	default:
		return "string"
	}
}
