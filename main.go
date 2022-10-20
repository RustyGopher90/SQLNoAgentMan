package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct  {
	UnmarshallMe string `json:"unmarshallMe"`

}

func main() {
	config := new(Config)
	configData, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal([]byte(configData), &config)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(config.UnmarshallMe)
}