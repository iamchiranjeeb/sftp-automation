package connection

import (
	"fmt"
	"log"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)


func GetPrivateKeyConnection(user string,keyPath string,ip string,port string)(*ssh.Client,error){

	fmt.Println("Looseeeeeerr..............")

	ipPort := ip + ":" + port

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			PublicKey(keyPath),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn,err := ssh.Dial("tcp",ipPort,config)

	if err != nil {
		return nil,err
	}

	return conn,nil
}

func PublicKey(keyPath string)(ssh.AuthMethod){

	key,err := ioutil.ReadFile(keyPath)

	if err != nil {
		log.Fatal(err)
	}

	signer,err := ssh.ParsePrivateKey(key)

	if err != nil {
		log.Fatal(err)
	}

	return ssh.PublicKeys(signer)
}