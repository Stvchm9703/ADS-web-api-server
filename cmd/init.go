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

var initCMDInput = struct {
	cfPath       string
	mode         string
	checkImpTree bool
	rootPath     string
	skipFol      string
	refDataPath  string
}{}

var initCMD = &cobra.Command{
	Use:   "init",
	Short: "init the mongodb server of webserver",
	Long:  `webserver server start run `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webserver Static Site Generator v0.9 -- HEAD")
		if len(args) > 0 {
			fmt.Println(args)
		}
		// fmt.Println(initCMDInput.cfPath)
		var configPoint *common.ConfigTemp
		var err error
		if strings.Contains(initCMDInput.cfPath, ".toml") {
			configPoint, err = Cf.OpenToml(initCMDInput.cfPath)
		} else if strings.Contains(initCMDInput.cfPath, ".yaml") {
			configPoint, err = Cf.OpenYaml(initCMDInput.cfPath)
		}
		log.Println(configPoint)
		log.Println(initCMDInput.mode)
		if err == nil {
			Wb.ServerInitProc(configPoint)
		} else {
			panic(err)
		}
	},
}

func init() {
	callPath, _ := os.Getwd()
	initCMD.Flags().StringVarP(
		&initCMDInput.cfPath,
		"conf", "c",
		filepath.Join(callPath, "config.toml"),
		"start server with specific config file")
	rootCmd.AddCommand(initCMD)
}
