package connection

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

func GetPasswordConnection(user string,pass string,ip string,port string)(*ssh.Client,error){
	fmt.Println("Looseeeeeerr..............")

	ipPort := ip + ":" + port

	config := &ssh.ClientConfig{

		User:user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn,err := ssh.Dial("tcp",ipPort,config)

	if err != nil {
		return nil,err
	}

	return conn,nil
}