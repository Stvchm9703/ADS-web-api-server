package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	// "webserver"
	conf "webserver/server/common"
	"webserver/server/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func ServerMainProcess(configPara *conf.ConfigTemp, path string, mode string) {
	// NOTE: add Config reading
	log.Println(configPara)

	conf.ConfigInRun = configPara
	conf.PathInRun = path
	gin.SetMode(gin.ReleaseMode)
	if mode != "test" {
		DB, err := model.ConnectDB(configPara)
		fmt.Println()
		log.Println("Svr main procx")
		fmt.Println(DB)
		fmt.Println()
		if err != nil {
			log.Println(err)
		}
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
		if mode != "test" {
			model.DisconnDB()
		}
	}
}

func ServerInitProc(configPara *conf.ConfigTemp, exportPath *string) {
	// NOTE: add Config reading
	log.Println(configPara)
	DB, err := model.ConnectDB(configPara)
	fmt.Println()
	fmt.Print(DB)
	if err != nil {
		log.Println(err)
	}
	strct, r, e := model.CreateDBTable(configPara)

	fmt.Println("r:", r, "\n\te:", e)
	time.Sleep(750 * time.Millisecond)
	if exportPath != nil {
		var err error
		file, _ := json.MarshalIndent(strct, "", " ")
		if strings.Index(*exportPath, ".json") == len(*exportPath)-5 {
			err = ioutil.WriteFile(*exportPath, file, 0644)
		} else {
			err = ioutil.WriteFile(path.Join(*exportPath, "create_schema.json"), file, 0644)
		}
		if err != nil {
			fmt.Println(err.Error())
		}
	}
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
