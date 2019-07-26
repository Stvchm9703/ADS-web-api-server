package wildbase

import (
	"log"
	"net/http"
	"strconv"
	conf "webserver/common"
	"webserver/model"

	"github.com/gin-gonic/gin"
	"github.com/go-bongo/bongo"
	"golang.org/x/sync/errgroup"
)

var (
	g  errgroup.Group
	DB *bongo.Connection = nil
)

func ServerMainProcess(configPara *conf.ConfigTemp, mode string) {
	// NOTE: add Config reading

	DB, err := model.ConnectDB(configPara)

	log.Print(DB)
	if err != nil {
		log.Println(err)
	}
	if mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Println(configPara.Server.IP + ":" + strconv.Itoa(configPara.Server.Port))
	router := RouterSetting(configPara)
	server01 := &http.Server{
		Addr:           configPara.Server.IP + ":" + strconv.Itoa(configPara.Server.Port),
		Handler:        router,
		ReadTimeout:    conf.ServerConf.ReadTimeout,
		WriteTimeout:   conf.ServerConf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
		model.DB.Close()
	}
}
