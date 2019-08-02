package CFfile

import (
	"fmt"
	"os"
	"path/filepath"
	// "strings"
	"errors"
	"strconv"

	// _ "webserver/cmd"

	"github.com/manifoldco/promptui"
	// "github.com/iancoleman/strcase"
	conf "webserver/server/common"
	// "github.com/spf13/cobra"
)
// SECTION: BaseConfig

var rootFilepath, _ = os.Getwd()

// 
// SECTION 

func ServerForm(initForm *conf.ConfigTemp) bool {
	//  Server
	fmt.Println("= --- [ Develop Server ]")
	fmt.Println("=")
	// -------------------------
	
	prompt2 := promptui.Prompt{
		Label:   "= ----- Server IP (default: 127.0.0.1) : ",
		Default: "127.0.0.1",
	}
	result2, err := prompt2.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.Server.IP = result2
	fmt.Println("=")
	// -------------------------
	prompt3 := promptui.Prompt{
		Label:   "= ----- Server Port (default: 8081) : ",
		Default: "8081",
		Validate: func(input string) error {
			_, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("Invalid number")
			}
			return nil
		},
	}
	result3, err := prompt3.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	portNum, _ := strconv.ParseInt(result3, 10, 32)
	initForm.Server.Port = int(portNum)
	if portNum == 0 {
		initForm.Server.Port = 8081
	}
	fmt.Println("=")
	// -------------------------
	devRootpath := filepath.Join(initForm.Server.RootFilePath, "dev")
	outputPath := initForm.Server.IP + ":" + strconv.Itoa(initForm.Server.Port) 
	//
	fmt.Println("= \t the system root filepath of the website / web-app")
	fmt.Println("= \t (default: {call path}/dev ) : ")

	prompt5 := promptui.Prompt{
		Label:   "= ----- Server root filepath : ",
		Default: devRootpath,
	}
	result5, err := prompt5.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.Server.RootFilePath = result5
	fmt.Println("=")
	// -------------------------

	fmt.Println("= \t the system root output path (url) of the website / web-app ")
	fmt.Println("= \t (default: ./dev => "+ outputPath +"/dev ) : ")
	
	prompt4 := promptui.Prompt{
		Label:   "= ----- Server root output path : ",
		Default: "./dev",
	}
	result4, err := prompt4.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.Server.MainPath = result4
	fmt.Println("=")
	if result4[len(result4)-1:] != "/" {
		outputPath = outputPath + (result4[1:len(result4)])
	} else {
		outputPath = outputPath + (result4[1:len(result4)-1])
	}

	// -------------------------
	fmt.Println("= \t the static source filepath of the website / web-app ")
	fmt.Println("= \t (default: {call path}/static ) : ")

	prompt6 := promptui.Prompt{
		Label:   "= ----- Server static filepath : ",
		Default: filepath.Join(initForm.Server.RootFilePath, "static"),
	}
	result6, err := prompt6.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.Server.StaticFilepath = result6
	fmt.Println("=")
	// -------------------------

	fmt.Println("= \t the static output path (url) the website / web-app ")
	fmt.Println("= \t (default: ./static => " + outputPath+ "/static )")

	prompt7 := promptui.Prompt{
		Label:   "= ----- Server static output path : ",
		Default: "./static",
	}
	result7, err := prompt7.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.Server.StaticOutpath = result7
	fmt.Println("=")
	// -------------------------
	fmt.Println("= \t the template source filepath of the website / web-app")
	fmt.Println("= \t (default: {call path}/template ) : ")

	prompt8 := promptui.Prompt{
		Label:   "= ----- Server template filepath : ",
		Default: filepath.Join(initForm.Server.RootFilePath, "template"),
	}
	result8, err := prompt8.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.Server.TemplateFilepath = result8
	fmt.Println("=")
	// -------------------------
	fmt.Println("= \t the template output path (url) the website / web-app")
	fmt.Println("= \t (default: ./template => " + outputPath+ "/template ) :")
	
	prompt9 := promptui.Prompt{
		Label:   "= ----- Server template output path : ",
		Default: "./template",
	}
	result9, err := prompt9.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.Server.TemplateOutpath = result9
	fmt.Println("=")

	// -------------------------
	fmt.Println("= \t the api table filepath of the website / web-app")
	fmt.Println("= \t (default: {call path}/api.form.yaml ) : ")
	
	prompt10 := promptui.Prompt{
		Label:   "= ----- Server template filepath : ",
		Default: filepath.Join(initForm.Server.RootFilePath, "api.form.yaml"),
	}
	result10, err := prompt10.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.Server.APITablePath = result10
	fmt.Println("=")
	// -------------------------
	fmt.Println("= \t the api output path (url) the website / web-app")
	fmt.Println("= \t (default: ./api => " + outputPath + "/api ) :")

	prompt11 := promptui.Prompt{
		Label:   "= ----- Server template output path : ",
		Default: "./api",
	}
	result11, err := prompt11.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.Server.APIOutpath = result11
	fmt.Println("=")

	fmt.Println("= -------------------------------------- ")
	// ---------------------------------------------------
	// DatabaseForm()
	return true
}

func DatabaseForm( initForm *conf.ConfigTemp) bool{
	//  Database
	dbDefaultCF := []struct {
		DisplayName string 
		Connector 	string 
		Host        string 
		Port        int
	} {
		{DisplayName: "MySql", Connector: "mysql", Host: "127.0.0.1", Port: 3306 },
		{DisplayName: "PostgreSQL", Connector: "postgres", Host: "127.0.0.1", Port: 5432 },
		{DisplayName: "SQLite", Connector: "sqlite3", Host: "127.0.0.1", Port: 0 },
		{DisplayName: "MS SQL", Connector: "mssql", Host: "127.0.0.0", Port: 3306 },
	}

	formTemp := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .DisplayName | cyan }} ",
		Inactive: "  {{ .DisplayName | cyan }}",
		Selected: " -> {{ .DisplayName | red | cyan }}",
		Details: `
--------- Engine ----------
{{ "DB Engine:" | faint }}	{{ .DisplayName }}
{{ "Host:" | faint }}	{{ .Host }}
{{ "Port:" | faint }}	{{ .Port }}`,
	}
	
	fmt.Println("= --- [ Database Server Connector ]")
	fmt.Println("=")
	// -------------------------
	prompt := promptui.Select{
		Label:     "= Select Database Engine (default: mysql) : ",
		Items:     dbDefaultCF,
		Templates: formTemp,
		IsVimMode: false,
		Size:      4,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("= Prompt failed %v = \n", err)
		return false
	}

	fmt.Println("= --- [ Custom Conneter ]")

	fmt.Println("= \t Host : " + dbDefaultCF[i].Host)
	
	fmt.Println("= \t Port : " + strconv.Itoa(dbDefaultCF[i].Port))

	prmpt := promptui.Select{
		Label: "= change the Connector Default Setting?",
		Items: []string{"Yes", "No"},
	}

	isCustom, _, err := prmpt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}


	initForm.Database.Connector = dbDefaultCF[i].Connector

	if isCustom == 1  {
		initForm.Database.Host = dbDefaultCF[i].Host
		initForm.Database.Port = dbDefaultCF[i].Port
	} else {
		prompt2 := promptui.Prompt{
			Label:   "= ----- Database Server IP/Host (default: " + dbDefaultCF[i].Host +") : ",
			Default: dbDefaultCF[i].Host,
		}
		result2, err := prompt2.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return false
		}
		initForm.Database.Host = result2
		fmt.Println("=")
		// -------------------------
		prompt3 := promptui.Prompt{
			Label:   "= ----- Database Server Port (default: "+ strconv.Itoa(dbDefaultCF[i].Port)+") : ",
			Default: strconv.Itoa(dbDefaultCF[i].Port),
			Validate: func(input string) error {
				_, err := strconv.ParseInt(input, 10, 32)
				if err != nil {
					return errors.New("Invalid number")
				}
				return nil
			},
		}
		result3, err := prompt3.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return false
		}
		portNum, _ := strconv.Atoi(result3)
		initForm.Database.Port = portNum
		fmt.Println("=")
	}
	
	// -------------------------
	// devRootpath := filepath.Join(initForm.Server.RootFilePath, "dev")
	// outputPath := initForm.Server.IP + ":" + strconv.Itoa(initForm.Server.Port) 
	//
	fmt.Println("= \t the database connection account username ")
	fmt.Println("= \t (default: root ) : ")

	prompt5 := promptui.Prompt{
		Label:   "= ----- Username : ",
		Default: "root",
	}
	result5, err := prompt5.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.Database.Username = result5
	fmt.Println("=")
	// -------------------------

	fmt.Println("= \t the database connection account Password")
	fmt.Println("= \t (default:  ) : ")

	prompt4 := promptui.Prompt{
		Label:   "= ----- Password : ",
		Default: "",
	}
	result4, err := prompt4.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.Database.Password = result4
	fmt.Println("=")
	

	
	return true
}


func DatabaseFallbackForm(initForm *conf.ConfigTemp) bool{
	fmt.Println("= --- [ Database Server Fallback ]")
	fmt.Println("=")
	
	fmt.Println("= \t the storage filepath of database fallback file ")
	fmt.Println("= \t (default: {call path}/db/{System Name}.fall.db => " + filepath.Join (rootFilepath,"db", "system" +".fall.db" ) + " )")

	prompt7 := promptui.Prompt{
		Label:   "= ----- filepath : ",
		Default: filepath.Join (rootFilepath,"db", "system" +".fall.db" ) ,
	}
	
	result7, err := prompt7.Run()
	
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	initForm.DatabaseFallback.Schema = "db"
	initForm.DatabaseFallback.Filepath = result7
	fmt.Println("=")

	return true
}
