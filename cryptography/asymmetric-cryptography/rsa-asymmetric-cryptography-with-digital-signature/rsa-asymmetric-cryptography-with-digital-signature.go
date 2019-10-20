// my-go-examples rsa-asymmetric-cryptography-wth-digital-signature.go

package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

var filename string

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func readFile(filename string) string {

	plainTextByte, err := ioutil.ReadFile(filename)
	checkErr(err)
	// Convert to string
	plainText := string(plainTextByte)
	return plainText

}

func generateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey) {

	// GENERATE PRIVATE & PUBLIC KEY PAIR
	privateKeyRaw, err := rsa.GenerateKey(rand.Reader, 2048)
	checkErr(err)

	// EXTRACT PUBLIC KEY
	publicKeyRaw := &privateKeyRaw.PublicKey

	return privateKeyRaw, publicKeyRaw

}

func encryptMessage(publicKeyRaw *rsa.PublicKey, plainText string) string {

	plainTextByte := []byte(plainText)
	label := []byte("")
	hash := sha256.New()

	// ENCRYPT
	cipherTextByte, err := rsa.EncryptOAEP(
		hash,
		rand.Reader,
		publicKeyRaw,
		plainTextByte,
		label,
	)
	checkErr(err)

	// ENCODE - RETURN HEX
	cipherText := hex.EncodeToString(cipherTextByte)
	return cipherText

}

func decryptMessage(privateKeyRaw *rsa.PrivateKey, cipherText string) string {

	// DECODE cipherText
	cipherTextByte, _ := hex.DecodeString(cipherText)

	label := []byte("")
	hash := sha256.New()

	// DECRYPT DATA
	plainTextByte, err := rsa.DecryptOAEP(
		hash,
		rand.Reader,
		privateKeyRaw,
		cipherTextByte,
		label,
	)
	checkErr(err)

	// RETURN STRING
	plainText := string(plainTextByte[:])
	return plainText

}

func createSignature(senderPrivateKeyRaw *rsa.PrivateKey, plainText string) string {

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	PSSmessage := []byte(plainText)
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	// CREATE SIGNATURE
	signatureByte, err := rsa.SignPSS(
		rand.Reader,
		senderPrivateKeyRaw,
		newhash,
		hashed,
		&opts,
	)
	checkErr(err)

	// ENCODE - RETURN HEX
	signature := hex.EncodeToString(signatureByte)

	return signature

}

func verifySignature(senderPublicKeyRaw *rsa.PublicKey, signature string, plainText string) bool {

	var result bool

	// DECODE signature
	signatureByte, _ := hex.DecodeString(signature)

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	PSSmessage := []byte(plainText)
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	// VERIFY SIGNATURE
	err := rsa.VerifyPSS(
		senderPublicKeyRaw,
		newhash,
		hashed,
		signatureByte,
		&opts,
	)
	if err != nil {
		verifyStatus = false
	} else {
		verifyStatus = true
	}

	return verifyStatus
}

func init() {

	// GET FILE NAME FROM ARGS
	flag.Parse()
	filenameSlice := flag.Args()
	if len(filenameSlice) != 1 {
		err := errors.New("Only one filename allowed")
		checkErr(err)
	}

	filename = filenameSlice[0] // Make it a string

}

func main() {

	fmt.Println(" ")

	// READ FILE INTO STRING
	plainText := readFile(filename)
	fmt.Printf("The original message contains:\n\n%s\n\n", plainText)

	// SENDER GENERATE RSA KEYS
	senderPrivateKeyRaw, senderPublicKeyRaw := generateRSAKeys()

	// RECEIVER GENERATE RSA KEYS
	receiverPrivateKeyRaw, receiverPublicKeyRaw := generateRSAKeys()

	// ENCRYPT MESSAGE USING PRIVATE KEY
	cipherText := encryptMessage(receiverPublicKeyRaw, plainText)
	fmt.Printf("The encrypted message contains:\n\n%s\n\n", cipherText)

	// CREATE SIGNATURE
	signature := createSignature(senderPrivateKeyRaw, plainText)
	fmt.Printf("The senders signature:\n\n%s\n\n", signature)

	// DECRYPT MESSAGE USING PUBLIC KEY
	plainText = decryptMessage(receiverPrivateKeyRaw, cipherText)
	fmt.Printf("The received message contains:\n\n%s\n\n", plainText)

	// VERIFY SIGNATURE
	verifyStatus := verifySignature(senderPublicKeyRaw, signature, plainText)
	fmt.Printf("The senders signature is: %v\n\n", verifyStatus)

}
