package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	Cf "webserver/cmd/CFfile"
	Wb "webserver/server"
	"webserver/server/common"

	"github.com/spf13/cobra"
)

var importCMDInput = struct {
	cfPath       string
	mode         string
	checkImpTree bool
	rootPath     string
	skipFol      string
	refDataPath  string
}{}

var importCMD = &cobra.Command{
	Use:   "import",
	Short: "import the mongodb server of webserver",
	Long:  `webserver server start run `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webserver Static Site Generator v0.9 -- HEAD")
		if len(args) > 0 {
			fmt.Println(args)
		}
		// fmt.Println(importCMDInput.cfPath)
		var configPoint *common.ConfigTemp
		var err error
		if strings.Contains(importCMDInput.cfPath, ".toml") {
			configPoint, err = Cf.OpenToml(importCMDInput.cfPath)
		} else if strings.Contains(importCMDInput.cfPath, ".yaml") {
			configPoint, err = Cf.OpenYaml(importCMDInput.cfPath)
		}
		log.Println(configPoint)
		log.Println(importCMDInput.mode)
		if err == nil {
			Wb.ServerImportProc(configPoint, &importCMDInput.refDataPath)
		} else {
			panic(err)
		}
	},
}

func init() {
	callPath, _ := os.Getwd()
	importCMD.Flags().StringVarP(
		&importCMDInput.cfPath,
		"conf", "c",
		filepath.Join(callPath, "config.toml"),
		"start db server with specific config file")
	importCMD.Flags().StringVarP(
		&importCMDInput.refDataPath,
		"import", "i",
		filepath.Join(callPath, "/data"),
		"import json file for mongodb")
	rootCmd.AddCommand(importCMD)
}
