package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	// // "log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
	"webserver/server/common"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PageMeta struct {
	PageLimit int `json:"page_limit,omitempty"`
	PageNum   int `json:"page_num,omitempty"`
	Count     int `json:"count,omitempty"`
}

type MgoCursorRes struct {
	OK       int  `bson:"ok" json:"ok"`
	WaitedMS *int `bson:"waitedMS,omitempty" json:"waitedMS,omitempty"`
	Cursor   struct {
		ID         interface{}              `bson:"id" json:"id"`
		NS         string                   `bson:"ns" json:"ns"`
		FirstBatch []map[string]interface{} `bson:"firstBatch" json:"firstBatch"`
	} `bson:"cursor" json:"cursor"`
}

var DBSess *mgo.Session
var DBConn *mgo.Database
var refstructlist = []interface{}{
	CourseMod{},
	DepartmentMod{},
	EnrollMod{},
	OfferMod{},
	StudentMod{},
}

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
		fmt.Println(err)
		return nil, err
	}
	DBConn = DBSess.DB(temp.Database.Database)
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

func CreateDBTable(config *common.ConfigTemp) (bool, error) {
	if DBConn != nil {
		var structlist = []interface{}{
			DepartmentMod{},
			StudentMod{},
		}
		tmpBD := []bson.M{}
		for k, v := range structlist {
			fmt.Println("k", k)
			r := reflect.TypeOf(v)
			tmpProp := createBMap(v)

			// fmt.Println(tmpProp)
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
		fmt.Println()
		nameJ, _ := json.Marshal(tmpBD)
		fmt.Println("runcmd : ")
		fmt.Println(string(nameJ))
		nameList := MgoCursorRes{}
		DBConn.Run(bson.M{
			"listCollections": 1.0,
			"nameOnly":        true,
		}, &nameList)
		fmt.Println()
		nameJ, _ = json.Marshal(nameList)
		fmt.Println("resultCursor: ")
		fmt.Println(string(nameJ))

		// Note: Cursor Result
		fb := nameList.Cursor.FirstBatch
		finPush := []bson.M{}
	I:
		for ja := 0; ja < len(tmpBD); ja++ {
			fmt.Println("check:", tmpBD[ja]["create"])
			for i := 0; i < len(fb); i++ {
				// fmt.Println("loop check:", fb[i]["name"])
				if tmpBD[ja]["create"] == fb[i]["name"].(string) {
					// fmt.Println("-->break:")
					continue I
				}
			}
			fmt.Println("= not exist ")
			finPush = append(finPush, tmpBD[ja])
		}

		// resultCursor := mgoCursorRes{}
		// DBConn.Run(tmpBD, &resultCursor)

		nameJ, _ = json.Marshal(finPush)
		fmt.Println("result: ")
		fmt.Println(string(nameJ))

		if len(finPush) > 0 {
			failCount := 0
			for ks := 0; ks < len(finPush); {
				time.Sleep(750 * time.Millisecond)
				fmt.Println("check:", finPush[ks]["create"])
				resultCursor := MgoCursorRes{}
				DBConn.Run(finPush[ks], &resultCursor)
				nameJ, _ := json.Marshal(resultCursor)
				fmt.Println("resultCursor: ")
				fmt.Println(string(nameJ))
				if resultCursor.OK == 1 {
					ks++
					failCount = 0
				} else {
					failCount++
					time.Sleep(750 * time.Millisecond)
					if failCount > 10 {
						return false, common.ErrorMessage{
							When: time.Now(),
							What: "Failure : " + string(nameJ),
						}
						break
					}
				}
			}

		} else {
			return false, common.ErrorMessage{
				When: time.Now(),
				What: "the DB is Already Declare",
			}
		}

	}
	_, err := NotConn()
	return false, err
}

func createBMap(v interface{}) bson.M {
	r := reflect.TypeOf(v)
	tmpProp := bson.M{}
	for ri := 0; ri < r.NumField(); ri++ {
		smp := strings.Split(r.Field(ri).Tag.Get("bson"), ",")[0]
		ftmp := typeConv(r.Field(ri))
		// SECTION
		if ftmp == "object" {
			xx := reflect.New(r.Field(ri).Type).Elem().Interface()
			ccc := reflect.Type(r.Field(ri).Type).Kind()

			if ccc == reflect.Ptr {
				xx = reflect.New(r.Field(ri).Type.Elem()).Elem().Interface()
			}
			tmpProp[smp] = bson.M{
				"bsonType":    ftmp,
				"description": "Go System AutoGen : " + r.Field(ri).Name,
				"properties":  createBMap(xx),
			}
		} else if ftmp == "array" {

			xx := reflect.New(r.Field(ri).Type.Elem()).Elem().Interface()
			ccc := reflect.Type(r.Field(ri).Type.Elem()).Kind()

			if ccc == reflect.Ptr {
				xx = reflect.New(r.Field(ri).Type.Elem().Elem()).Elem().Interface()
			}
			tmpProp[smp] = bson.M{
				"bsonType":    ftmp,
				"description": "Go System AutoGen : " + r.Field(ri).Name,
				"properties":  createBMap(xx),
			}
			// SECTION
		} else {
			tmpProp[smp] = bson.M{
				"bsonType":    ftmp,
				"description": "Go System AutoGen : " + r.Field(ri).Name,
			}
		}
		fmt.Println("tmpProp[smp]:")
		fmt.Println(tmpProp[smp])
	}
	eee, _ := json.Marshal(tmpProp)
	fmt.Println(eee, string(eee))
	return tmpProp
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

func remove(s []bson.M, i int) []bson.M {
	return append(s[:i], s[i+1:]...)
}

// for futune
func ImportData(folderPath *string) (bool, error) {
	tmpBD := []string{
		"Course",
		"Department",
		"Enroll",
		"Offer",
		"Student",
	}
	if folderPath == nil {
		return false, common.ErrorMessage{
			When: time.Now(),
			What: "folderPath is nil",
		}
	}
	if DBConn != nil {
		failList := []map[string]interface{}{}
		for k := 0; k < len(tmpBD); {
			jsonFile, err := os.Open(filepath.Join(*folderPath, tmpBD[k]+".json"))
			if err != nil {
				fmt.Println(err)
				return false, err
			}
			fmt.Printf("Successfully Opened %v.json", tmpBD[k])
			byteValue, _ := ioutil.ReadAll(jsonFile)
			fil := []map[string]interface{}{}
			json.Unmarshal(byteValue, &fil)
			fmt.Println("fil:\n\t", fil)
			jsonFile.Close()
			for index := 0; index < len(fil); index++ {
				err := DBConn.C(tmpBD[k]).Insert(fil[index])
				fmt.Println("mod", tmpBD[k], "\nbatch:", index, "\n\terr:", err)
				fmt.Println("refDoc: \t", fil[index])
				if err != nil {
					errmsg := map[string]interface{}{
						"modname":    tmpBD[k],
						"batchindex": index,
						"refDoc":     fil[index],
						"err":        err,
					}
					failList = append(failList, errmsg)
				}
			}
			k++
		}
		if len(failList) > 0 {
			return false, common.ErrorMessage{
				When:     time.Now(),
				What:     "run import batch err:",
				ErrorObj: failList,
			}
		}
		return true, nil
	}
	_, err := NotConn()
	return false, err
}

func ExportData(targetPath *string) (bool, error) {
	tmpBD := []string{
		"Course",
		"Department",
		"Enroll",
		"Offer",
		"Student",
	}

	if targetPath == nil {
		return false, common.ErrorMessage{
			When: time.Now(),
			What: "targetPath is nil",
		}
	}
	if DBConn != nil {

		failList := []map[string]interface{}{}
		for k := 0; k < len(tmpBD); k++ {
			raw := []bson.M{}
			err := DBConn.C(tmpBD[k]).Find(nil).All(&raw)
			if err != nil {
				errmsg := map[string]interface{}{
					"modname": tmpBD[k],
					"refDoc":  raw,
					"step":    "get from db",
					"err":     err,
				}
				failList = append(failList, errmsg)
			}
			file, _ := json.MarshalIndent(raw, "", " ")
			err = ioutil.WriteFile(filepath.Join(*targetPath, tmpBD[k]+".json"), file, 0644)
			if err != nil {
				errmsg := map[string]interface{}{
					"modname": tmpBD[k],
					"refDoc":  raw,
					"step":    "write json file",
					"err":     err,
				}
				failList = append(failList, errmsg)
			}

		}
		if len(failList) > 0 {
			return false, common.ErrorMessage{
				When:     time.Now(),
				What:     "run import batch err:",
				ErrorObj: failList,
			}
		}
		return true, nil
	}
	_, err := NotConn()
	return false, err
}
