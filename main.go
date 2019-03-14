package main

import (
	"fmt"
	"log"
	"os"

	remocloud "github.com/NaoyaTabakomori/go-nature-remo/cloud"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Token   string    `yaml:"token"`
	Signals []*signal `yaml:"signals"`
}

type signal struct {
	Name string `yaml:"name"`
	ID   string `yaml:"id"`
}

func main() {

	token := "YOURTOKEN"
	client := remocloud.NewClient(token)

	var signals []*signal
	config := config{Token: token, Signals: signals}
	apps, err := client.GetAppliances()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range apps {
		for _, v2 := range v.Signals {
			sig := signal{}
			sig.Name = v2.Name
			sig.ID = v2.ID
			config.Signals = append(config.Signals, &sig)
		}
	}

	y, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	output := fmt.Sprintf("%s", y)

	file, err := os.OpenFile("config.yaml", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(file, output)
}
