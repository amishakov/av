package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Token string
}

func (c *config) NewConfig(path string) {
	conf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Can't open conf file %v ", err)
	}
	err = yaml.Unmarshal(conf, c)
	if err != nil {
		log.Fatalf("Can't parse conf file: %v", err)
	}
	return c
}
