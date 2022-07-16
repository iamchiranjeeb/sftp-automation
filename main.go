package main

import(
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"com.ssh.connection/Helpers/getpath"
	"com.ssh.connection/Helpers/readfile"
	"com.ssh.connection/downloadupload"
	"com.ssh.connection/connection"
	"com.ssh.connection/cmdexec"
	"com.ssh.connection/Helpers/authType"
	"com.ssh.connection/Helpers/operationType"
)


// go run main.go --path=/home/iamchiranjeeb/Desktop/Demo/go/sshTest/config.json
func main() {

	configFilePath := getpath.GetConfigPath()
	
	if configFilePath == "/home" {
		fmt.Println("Wrong File path")
		os.Exit(1)
	}

	jsonContent := readfile.ReadFile(configFilePath)
	fmt.Println(jsonContent)

	for _,item := range jsonContent {
		
		creds := map[string]string{
			"user":fmt.Sprintf("%v",item["user"]),
			"auth":fmt.Sprintf("%v",item["auth"]),
			"ip":fmt.Sprintf("%v",item["ip"]),
			"port":fmt.Sprintf("%v",item["port"]),
		}

		var fileCreds = map[string]string{
			"local": fmt.Sprintf("%v",item["localPath"]),
			"target":fmt.Sprintf("%v",item["remotePath"]),
		}

		auth := authType.Password
		var opt operationType.PerformType

		reqOpt := fmt.Sprintf("%v",item["operation"])
		optToUpper := strings.ToUpper(reqOpt)

		fmt.Println(optToUpper)


		switch optToUpper {

		case "UPLOAD":
			opt = operationType.Upload
		case "DOWNLOAD":
			opt = operationType.Download
		default:
			log.Fatalf("Operation type %+v could not found.\nChange to upload or download in you yml/json file.",optToUpper)
			os.Exit(1)

		}

		if item["password"] != true {
			auth = authType.Key
		}

		conn,err := connection.GetConnection(auth,creds["user"], creds["auth"],creds["ip"],creds["port"])

		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()

		files := reflect.ValueOf(item["files"])

		fmt.Println(files)

		for i := 0; i < files.Len(); i++ {

			singleFile := files.Index(i).Elem()
			singleFileStr := fmt.Sprintf("%v",singleFile)

			downloadupload.SftpOperations(conn,creds,fileCreds,singleFileStr,opt)
		}

		if item["execCommand"]==true{

			commands := reflect.ValueOf(item["commands"])
			for i := 0; i < commands.Len(); i++{

				singleCmd := commands.Index(i).Elem()
				singleCmdStr := fmt.Sprintf("%v",singleCmd)
				outPut,err := cmdexec.ExecuteTerminalCmd(conn,singleCmdStr)

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(outPut)
			}

		}
	}

}