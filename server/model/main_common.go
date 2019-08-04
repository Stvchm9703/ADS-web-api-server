package model

import (
	"log"
	"strconv"
	"webserver/server/common"

	"gopkg.in/mgo.v2"
)

var DBSess *Session := nil
var DBConn *Database := nil
func ConnectDB(temp *common.ConfigTemp) (*Session, error) {
	Host := []string{
		temp.Database.Host + ":" + strconv.Itoa(temp.Database.Port),
	}
	DBSess, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		Username: temp.Database.Username,
		Password: temp.Database.Password,
		// Database: temp.Database.Database,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer DBSess.Close()
	DBConn = DBSess.DB(temp.Database.Database)
	return DBSess, nil
}

func DisconnDB()(bool,err){
	return DBSess.Close()
}

func CreateDBTable(config *common.ConfigTemp) {

}

func NotConn() ( *struct, error){
	return nil , &common.ErrorMessage{
		When: time.Now(),
		What: "MongoDB is not Connect",
	}
}

type PageMeta struct {
	PageLimit int `json:"page_limit,omitempty"`,
	PageNum int `json:"page_num,omitempty"`,
	Count int `json:"count,omitempty"`,
}



