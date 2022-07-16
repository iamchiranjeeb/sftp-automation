package readfile


import (
	"gopkg.in/yaml.v2"
	"log"
	"io/ioutil"
	// "fmt"
)


func ReadYml(ymlPath string)([]map[string]interface{}){
	log.Printf("Reading : %v", ymlPath)

	data, err := ioutil.ReadFile(ymlPath)

	if err != nil {
		log.Fatalf("Failed to open Yml file: %v", err)
	}

	dataBytes := []byte(data)

	var ymlData []map[string]interface{}
	
	yaml.Unmarshal(dataBytes, &ymlData)
	// fmt.Printf("Yml File Data : %+v",ymlData)
	// fmt.Printf("%T",ymlData)

	return ymlData
}