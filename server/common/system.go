package common

import (
	"fmt"
	"time"
)

const (
	QueryDefaultPageLimit = 500
	DBDefaultConn         = "root:@tcp(127.0.0.1:3306)/wildbase"
)
var (
	ConfigInRun *ConfigTemp
	PathInRun   string
)
var ServerConf = struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DefaultMode  string
}{
	ReadTimeout:  10 * time.Second,
	WriteTimeout: 10 * time.Second,
	DefaultMode:  "develop",
}

type ErrorMessage struct {
	When     time.Time
	What     string
	ErrorObj interface{}
}

func (e ErrorMessage) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}
