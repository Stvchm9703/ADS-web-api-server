package model

import (
	"context"
	"strconv"
	"time"
	"webserver/server/common"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBConn *mongo.Client = nil

func ConnectDB(temp *common.ConfigTemp) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	DBConn, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+temp.Database.Username+":"+temp.Database.Password+"@"+temp.Database.Host+":"+strconv.Itoa(temp.Database.Port)))
	if err != nil {
		return DBConn, err
	}
	return DBConn, err
}

func CreateDBTable(config *common.ConfigTemp) {

}
