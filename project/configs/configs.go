package configs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var E *Env

type Env struct {
	Atcoder Atcoder `yaml:"atcoder"`
	Note    Note    `yaml:"note"`
}

type Atcoder struct {
	User string `yaml:"user"`
}

type Note struct {
	User string `yaml:"user"`
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
