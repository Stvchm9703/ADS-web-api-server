package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	// "webserver"
	conf "webserver/server/common"
	"webserver/server/model"

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
	log.Println(configPara)
	DB, err := model.ConnectDB(configPara)
	fmt.Println()
	log.Println("Svr main procx")
	fmt.Print(DB)
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
		model.DisconnDB()
	}
}

func ServerInitProc(configPara *conf.ConfigTemp) {
	// NOTE: add Config reading
	log.Println(configPara)
	DB, err := model.ConnectDB(configPara)
	fmt.Println()
	fmt.Print(DB)
	if err != nil {
		log.Println(err)
	}
	r, e := model.CreateDBTable(configPara)
	fmt.Println("r:", r, "\n\te:", e)
	time.Sleep(750 * time.Millisecond)

	// fmt.Println("create view:")
	// r, e = viewMod.CreateSchemaVStudentEnroll()
	// fmt.Println("r:", r, "\n\te:", e)
	// time.Sleep(750 * time.Millisecond)

	// r, e = viewMod.CreateSchemaVCourseOffer()
	// fmt.Println("r:", r, "\n\te:", e)
	// time.Sleep(750 * time.Millisecond)

	time.Sleep(750 * time.Millisecond)
	DB.Close()
}

func ServerImportProc(configPara *conf.ConfigTemp, importPath *string) {
	// NOTE: add Config reading
	log.Println(configPara)
	DB, err := model.ConnectDB(configPara)
	fmt.Println()
	fmt.Print(DB)
	if err != nil {
		log.Println(err)
	}
	r, e := model.ImportData(importPath)
	fmt.Println("r:", r, "\n\te:", e)
	time.Sleep(750 * time.Millisecond)
	DB.Close()
}

func ServerExportProc(configPara *conf.ConfigTemp, exportPath *string) {
	// NOTE: add Config reading
	log.Println(configPara)
	DB, err := model.ConnectDB(configPara)
	fmt.Println()
	fmt.Print(DB)
	if err != nil {
		log.Println(err)
	}
	r, e := model.ExportData(exportPath)
	fmt.Println("r:", r, "\n\te:", e)
	time.Sleep(750 * time.Millisecond)
	DB.Close()
}
