package configs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var E *Env

type Env struct {
	Atcoder   Atcoder   `yaml:"atcoder"`
	Bookmater Bookmater `ymal:"bookmater"`
	Note      Note      `yaml:"note"`
	Life      Life      `yaml:"life"`
	Twitter   Twitter   `yaml:"twitter"`
}

type Atcoder struct {
	User string `yaml:"user"`
}

type Bookmater struct {
	User string `yaml:"user"`
}

type Note struct {
	User string `yaml:"user"`
}

type Life struct {
	SpreadsheetID string `yaml:"spreadsheet_id"`
	SheetName     string `yaml:"sheet_name"`
}

type Twitter struct {
	User          string `yaml:"user"`
	SpreadsheetID string `yaml:"spreadsheet_id"`
	SheetName     string `yaml:"sheet_name"`
}

func InitConfigs() {
	buf, err := ioutil.ReadFile("./env.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(buf, &E)
	if err != nil {
		log.Fatal(err)
	}
}
