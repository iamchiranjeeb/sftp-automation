package downloadupload


import (
	"log"
	"fmt"
	"io"
	"os"
	"github.com/pkg/sftp"
)


func Upload(sftpObj *sftp.Client,remoteFile string,localFile string){

	//create destination file
	dstFile,err := sftpObj.Create(remoteFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	//create source file
	srcFile, err := os.Open(localFile)
	if err != nil {
		log.Fatal(err)
	}

	//copy source to destination
	bytes,err := io.Copy(dstFile,srcFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)

	// buf := make([]byte, 1024)

	// for {
	// 	n, _ := srcFile.Read(buf)
	// 	if n == 0 {
	// 		break
	// 	}
	// 	dstFile.Write(buf)
	// }


}