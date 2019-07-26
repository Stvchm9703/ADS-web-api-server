package model

import (
	"webserver/common"

	"github.com/go-bongo/bongo"
)

func ConnectDB(temp *common.ConfigTemp) (*bongo.Connection, error) {
	return bongo.Connect(&bongo.Config{
		ConnectionString: temp.Database.Username + ":" + temp.Database.Password + "@" + temp.Database.Host + ":" + temp.Database.Password,
		Database:         temp.Database.Database,
	})
}
