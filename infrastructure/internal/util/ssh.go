package util

import (
	"golang.org/x/crypto/ssh"
	"log"
)

func GetSshClient(signer ssh.Signer, user string, ip string) *ssh.Client {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //delete for production
	}

	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	return client
}
