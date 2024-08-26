package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Values struct {
	Global struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address struct {
			City  string `json:"city"`
			State string `json:"state"`
		} `json:"address"`
	} `json:"global"`
}

func main() {
	var v Values
	jsonData := []byte(`
{
  "global": {
    "name": "Terrance Yeakey",
    "age": 30,
    "address": {
      "city": "Oklahoma City",
      "state": "OK"
    },
    "extra1": 1,
    "extra2": 2,
    "extra3": 3
  }
}`)
	err := json.Unmarshal(jsonData, &v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", v)
}
