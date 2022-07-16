package cmdexec

import (
	"log"
	"bytes"
	"golang.org/x/crypto/ssh"
)


func ExecuteTerminalCmd(sshConn *ssh.Client,command string)(string,error){

	session,err := sshConn.NewSession()
	if err != nil {
		return "Creating New Session Failed",err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b

	err = session.Run(command)

	// if err := session.Run("ls -ltr"); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(b.String())

	if err != nil {
		log.Fatal(err)
	}

	return b.String(),nil
}