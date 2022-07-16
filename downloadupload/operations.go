package downloadupload

import (
	"fmt"
	"log"
	"path"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"com.ssh.connection/Helpers/operationType"
)


func SftpOperations(sshConn *ssh.Client,creds map[string]string,locCred map[string]string,fileName string,opt operationType.PerformType){

	client, err := sftp.NewClient(sshConn)
	if err != nil {
		log.Fatal(err)
	}

	switch opt {
	case 0:
		fmt.Println("Uploading file to " + creds["ip"])
		var localFile string = locCred["local"] + "/" + fileName
		//create destination file
		//copying file same name as local
		var remoteFileName = path.Base(localFile)
		var fullRemoteFileName = path.Join(path.Join(locCred["target"], remoteFileName))
		//to copy in a different name we can concatenate a file name with the target path
		Upload(client,fullRemoteFileName,localFile)
	case 1:
		fmt.Println("Downloading from " + creds["ip"])
		var localFile string = locCred["target"] + "/" + fileName
		//copying file same name as remote
		var localFileName string = path.Base(localFile)
		var fullLocalFileName string = path.Join(path.Join(locCred["local"], localFileName))
		Download(client,localFile,fullLocalFileName)
	default:
		fmt.Println("Option not found.")
	}
}