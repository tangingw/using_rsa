package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"

	myRSA "github.com/tangingw/rsacrypto"
)

func readEncrypredFile(filename string) []byte {

	content, err := ioutil.ReadFile(filename)

	if err != nil {

		fmt.Printf("file: %s not found!\n", filename)
		os.Exit(1)
	}

	return content
}

func main() {

	if _, err := os.Stat("private.pem"); os.IsNotExist(err) {

		bitLength := 4096
		keyPair := myRSA.GenerateRSAKey(bitLength)

		myRSA.SavePublicPEMKey("public.pem", keyPair.PublicKey)
		myRSA.SavePEMKey("private.pem", keyPair.PrivateKey)
	}

	key := myRSA.RetrievePEMKey("private.pem")

	pubKey := myRSA.RetrievePEMPubKey("public.pem")
	keyPair := &myRSA.RSAKeyPair{pubKey, key}

	if _, err := os.Stat("encrypted_data.txt"); err == nil {

		decrypted := readEncrypredFile("encrypted_data.txt")
		decryptMessage := keyPair.DecryptOAEP(sha256.New(), decrypted, nil)
		fmt.Println(string(decryptMessage))

		serverMsg := []byte("I got your message!")
		encrypted := keyPair.EncryptOAEP(sha256.New(), serverMsg, nil)

		err := ioutil.WriteFile("server_message.txt", encrypted, 0644)

		if err != nil {
			panic(err)
		}

	} else if os.IsNotExist(err) {

		fmt.Println("We don't receive your message")

	} else if os.IsExist(err) {

		fmt.Println("Something must be wrong!")
	}

}
