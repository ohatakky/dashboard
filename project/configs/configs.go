package configs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var E *Env

type Env struct {
	Atcoder Atcoder `yaml:"atcoder"`
}

type Atcoder struct {
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
