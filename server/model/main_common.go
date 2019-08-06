package model

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"webserver/server/common"

	"gopkg.in/mgo.v2"
)

var DBSess *mgo.Session
var DBConn *mgo.Database

func ConnectDB(temp *common.ConfigTemp) (*mgo.Session, error) {
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
	return DBSess, nil
}

func DisconnDB() (bool, error) {
	DBSess.Close()
	return true, nil
}

func CreateDBTable(config *common.ConfigTemp) {

}

func NotConn() (*struct{}, error) {
	fmt.Println(DBSess)
	fmt.Println(DBConn)
	return nil, &common.ErrorMessage{
		When: time.Now(),
		What: "MongoDB is not Connect",
	}
}

type PageMeta struct {
	PageLimit int `json:"page_limit,omitempty"`
	PageNum   int `json:"page_num,omitempty"`
	Count     int `json:"count,omitempty"`
}
