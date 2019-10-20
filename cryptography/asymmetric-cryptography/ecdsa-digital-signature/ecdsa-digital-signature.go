// my-go-examples ecdsa-asymmetric-cryptography.go

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"hash"
	"io"
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

func generateECDSAKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {

	// GENERATE PRIVATE & PUBLIC KEY PAIR
	curve := elliptic.P256()
	privateKeyRaw, err := ecdsa.GenerateKey(curve, rand.Reader)
	checkErr(err)

	// EXTRACT PUBLIC KEY
	publicKeyRaw := &privateKeyRaw.PublicKey

	return privateKeyRaw, publicKeyRaw

}

func createSignature(senderPrivateKeyRaw *ecdsa.PrivateKey, plainText string) string {

	// HASH plainText
	var h hash.Hash
	h = sha256.New()
	io.WriteString(h, plainText)
	signHash := h.Sum(nil)

	// CREATE SIGNATURE
	r := big.NewInt(0)
	s := big.NewInt(0)

	r, s, err := ecdsa.Sign(rand.Reader, senderPrivateKeyRaw, signHash)
	checkErr(err)

	signatureByte := r.Bytes()
	signatureByte = append(signatureByte, s.Bytes()...)

	// ENCODE - RETURN HEX
	signature := hex.EncodeToString(signatureByte)

	return signature

}

func verifySignature(senderPublicKeyRaw *ecdsa.PublicKey, signature string, plainText string) bool {

	// HASH plainText
	var h hash.Hash
	h = sha256.New()
	io.WriteString(h, plainText)
	signHash := h.Sum(nil)

	// DECODE signature
	signatureByte, _ := hex.DecodeString(signature)

	// VERIFY SIGNATURE
	r := big.NewInt(0)
	s := big.NewInt(0)
	sigLen := len(signatureByte)
	r.SetBytes(signatureByte[:(sigLen / 2)])
	s.SetBytes(signatureByte[(sigLen / 2):])

	verifyStatus := ecdsa.Verify(senderPublicKeyRaw, signHash, r, s)

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
	senderPrivateKeyRaw, senderPublicKeyRaw := generateECDSAKeys()

	// CREATE SIGNATURE
	signature := createSignature(senderPrivateKeyRaw, plainText)
	fmt.Printf("The senders signature:\n\n%s\n\n", signature)

	// VERIFY SIGNATURE
	verifyStatus := verifySignature(senderPublicKeyRaw, signature, plainText)
	fmt.Printf("The senders signature is: %v\n\n", verifyStatus)

}
