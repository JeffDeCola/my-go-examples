// my-go-examples ecdsa-asymmetric-cryptography.go

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"

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

func generateECDSAKeys() (string, string) {

	// GENERATE PRIVATE & PUBLIC KEY PAIR
	curve := elliptic.P256()
	privateKeyRaw, err := ecdsa.GenerateKey(curve, rand.Reader)
	checkErr(err)

	// EXTRACT PUBLIC KEY
	publicKeyRaw := &privateKeyRaw.PublicKey

	// ENCODE
	privateKeyHex, publicKeyHex := encodeKeys(privateKeyRaw, publicKeyRaw)

	return privateKeyHex, publicKeyHex

}

// encodeKeys - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex
func encodeKeys(privateKeyRaw *ecdsa.PrivateKey, publicKeyRaw *ecdsa.PublicKey) (string, string) {

	privateKeyx509Encoded, _ := x509.MarshalECPrivateKey(privateKeyRaw)
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: privateKeyx509Encoded,
		})
	privateKeyHex := hex.EncodeToString(privateKeyPEM)

	publicKeyx509Encoded, _ := x509.MarshalPKIXPublicKey(publicKeyRaw)
	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyx509Encoded,
		})
	publicKeyHex := hex.EncodeToString(publicKeyPEM)

	return privateKeyHex, publicKeyHex

}

func createSignature(privateKeyHex string, plainText string) string {

	// DECODE PRIVATE KEY
	privateKeyPEM, _ := hex.DecodeString(privateKeyHex)
	block, _ := pem.Decode([]byte(privateKeyPEM))
	privateKeyx509Encoded := block.Bytes
	privateKeyRaw, _ := x509.ParseECPrivateKey(privateKeyx509Encoded)

	// HASH plainText
	hashedPlainText := sha256.Sum256([]byte(plainText))
	hashedPlainTextByte := hashedPlainText[:]

	r := big.NewInt(0)
	s := big.NewInt(0)

	// CREATE SIGNATURE
	r, s, err := ecdsa.Sign(
		rand.Reader,
		privateKeyRaw,
		hashedPlainTextByte,
	)
	checkErr(err)

	signatureByte := r.Bytes()
	signatureByte = append(signatureByte, s.Bytes()...)

	// ENCODE - RETURN HEX
	signature := hex.EncodeToString(signatureByte)

	return signature

}

func verifySignature(publicKeyHex string, signature string, plainText string) bool {

	// DECODE PUBLIC KEY
	publicKeyPEM, _ := hex.DecodeString(publicKeyHex)
	blockPub, _ := pem.Decode([]byte(publicKeyPEM))
	publicKeyx509Encoded := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(publicKeyx509Encoded)
	publicKeyRaw := genericPublicKey.(*ecdsa.PublicKey)

	// HASH plainText
	hashedPlainText := sha256.Sum256([]byte(plainText))
	hashedPlainTextByte := hashedPlainText[:]

	// DECODE signature
	signatureByte, _ := hex.DecodeString(signature)

	// EXTRACT R & S
	r := big.NewInt(0)
	s := big.NewInt(0)
	sigLen := len(signatureByte)
	r.SetBytes(signatureByte[:(sigLen / 2)])
	s.SetBytes(signatureByte[(sigLen / 2):])

	// VERIFY SIGNATURE
	verifyStatus := ecdsa.Verify(
		publicKeyRaw,
		hashedPlainTextByte,
		r,
		s,
	)

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
	senderPrivateKeyHex, senderPublicKeyHex := generateECDSAKeys()

	// CREATE SIGNATURE
	signature := createSignature(senderPrivateKeyHex, plainText)
	fmt.Printf("The senders signature:\n\n%s\n\n", signature)

	// VERIFY SIGNATURE
	verifyStatus := verifySignature(senderPublicKeyHex, signature, plainText)
	fmt.Printf("The senders signature is: %v\n\n", verifyStatus)

}
