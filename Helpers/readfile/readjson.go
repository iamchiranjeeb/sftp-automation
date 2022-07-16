package readfile

import(
	"encoding/json"
	"log"
	"io/ioutil"
	// "fmt"
)


func ReadJson(jsonPath string)([]map[string]interface{}){
	log.Printf("Reading : %v", jsonPath)

	data, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		log.Fatalf("Failed to open json file: %v", err)
	}

	dataBytes := []byte(data)

	var jsonData []map[string]interface{}
	json.Unmarshal(dataBytes, &jsonData)
	// fmt.Printf("Json File Data : %+v",jsonData)

	return jsonData

}