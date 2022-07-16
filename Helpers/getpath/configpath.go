package getpath

import (
	"fmt"
	"flag"
)

func GetConfigPath()(string){

	configFilePath := flag.String("path","/home","Config file path should be provided.")
	flag.Parse()
	configPath := fmt.Sprintf("%s",*configFilePath)
	return configPath

}