package connection

import(
	"com.ssh.connection/Helpers/authType"
	"golang.org/x/crypto/ssh"
)

func GetConnection(auth authType.Auth,user string,keyPath string,ip string,port string)(*ssh.Client,error){

	var conn *ssh.Client
	var err error

	switch auth {
		case 0:
			conn,err = GetPrivateKeyConnection(user, keyPath, ip, port)
		case 1:
			conn,err = GetPasswordConnection(user, keyPath, ip, port)
	}

	return conn, err
}