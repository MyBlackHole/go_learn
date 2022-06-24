package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	RabbitMQ struct {
		Host     string `yaml:"host"`
		Port     int64  `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

func readConf(filename string) (*Conf, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var conf Conf
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	return &conf, nil
}

func main() {
	conf, err := readConf("conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", conf)
}
