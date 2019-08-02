package CFfile

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	conf "webserver/server/common"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

func CreateConfigToml(path string, initForm *conf.ConfigTemp) {
	fmt.Println("= ---- creating config.toml -----")
	fileLocate, err := os.Create(filepath.Join(path, "config.toml"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	Writer := bufio.NewWriter(fileLocate)
	EncoderA := toml.NewEncoder(Writer)
	EncoderA.Encode(initForm)
	fmt.Println("= ")
}

func CreateConfigYaml(path string, initForm *conf.ConfigTemp) {
	fmt.Println("= ---- creating config.yaml -----")
	fileLocate, err := os.Create(filepath.Join(path, "config.yaml"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	Writer := bufio.NewWriter(fileLocate)
	d, err := yaml.Marshal(initForm)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
	fmt.Fprint(Writer, string(d))
	Writer.Flush()
}

func OpenToml(path string) *conf.ConfigTemp {
	fmt.Println("= ---- open config.toml -----")
	fmt.Println(path)

	var config conf.ConfigTemp
	if _, err := toml.DecodeFile(path, &config); err != nil {
		fmt.Println(err)
		// panic(err)
	}

	return &config
}
