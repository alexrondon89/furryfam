package util

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

func GetSigner(address string) ssh.Signer {
	err := os.Chmod(address, 0400)
	if err != nil {
		log.Fatalf("Failed to change file permissions: %v", err)
	}

	key, err := os.ReadFile(address)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	return signer
}
