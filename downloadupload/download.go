package downloadupload


import (
	"log"
	"fmt"
	"io"
	"os"
	"github.com/pkg/sftp"
)

func Download(sftpObj *sftp.Client,remoteFile string,localFile string){

	// open source file(Remote File)
	srcFile, err := sftpObj.Open(remoteFile)
	if err != nil {
	   log.Fatal(err)
	}

	// create destination file(Local File)
	dstFile, err := os.Create(localFile)
	if err != nil {
	   log.Fatal(err)
	}
	defer dstFile.Close()

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
	   log.Fatal(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)

	// flush in-memory copy
	err = dstFile.Sync()
	if err != nil {
	   log.Fatal(err)
	}

	// if _, err = srcFile.WriteTo(dstFile); err != nil {
	// 	log.Fatal(err)
	// }
}