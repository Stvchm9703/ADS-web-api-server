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

var runCMDInput = struct {
	cfPath       string
	mode         string
	checkImpTree bool
	rootPath     string
	skipFol      string
}{}

var runCmd = &cobra.Command{
	Use:   "start",
	Short: "start the server of webserver",
	Long:  `webserver server start run `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webserver Static Site Generator v0.9 -- HEAD")
		if len(args) > 0 {
			fmt.Println(args)
		}

		fmt.Println(runCMDInput.cfPath)

		var configPoint *common.ConfigTemp
		var err error
		if strings.Contains(runCMDInput.cfPath, ".toml") {
			configPoint, err = Cf.OpenToml(runCMDInput.cfPath)
		} else if strings.Contains(runCMDInput.cfPath, ".yaml") {
			configPoint, err = Cf.OpenYaml(runCMDInput.cfPath)
		}
		log.Println(configPoint)
		log.Println(runCMDInput.mode)
		if err == nil {
			Wb.ServerMainProcess(configPoint, runCMDInput.mode)
		} else {
			panic(err)
		}
	},
}

func init() {
	callPath, _ := os.Getwd()
	runCmd.Flags().StringVarP(
		&runCMDInput.cfPath,
		"conf", "c",
		filepath.Join(callPath, "config.toml"),
		"start server with specific config file")

	runCmd.Flags().StringVarP(
		&runCMDInput.mode,
		"mode", "m",
		"prod",
		"server running mode [prod / dev]")

	rootCmd.AddCommand(runCmd)
}
