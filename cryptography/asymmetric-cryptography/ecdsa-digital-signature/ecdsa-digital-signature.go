// my-go-examples ecdsa-asymmetric-cryptography.go

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
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

func generateECDSAKeys() (ecdsa.PrivateKey, []byte) {

	// GENERATE PUBLIC/PRIVATE KEY PAIR (32 Bytes/256 bits)
	curve := elliptic.P256()
	privateKeyRAW, err := ecdsa.GenerateKey(curve, rand.Reader)
	checkErr(err)

	// GET DIFFERENT FORMATS
	privateKeys509Binary, _ := x509.MarshalECPrivateKey(privateKeyRAW)
	privateKeyPEMByte := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: privateKeys509Binary,
		})
	privateKeyHex := hex.EncodeToString(privateKeyPEMByte)
	privateKeyBase64 := base64.StdEncoding.EncodeToString(privateKeyPEMByte)

	fmt.Printf("BINARY\n%b\n", privateKeys509Binary)
	fmt.Printf("BYTE\n%v\n", privateKeyPEMByte)
	fmt.Printf("HEX\n%v\n", privateKeyHex)
	fmt.Printf("BASE64\n%v\n", privateKeyBase64)

	// GET PUBLIC KEY FROM PRIVATE KEY
	publicKeyByte := append(privateKeyECDSA.PublicKey.X.Bytes(), privateKeyECDSA.PublicKey.Y.Bytes()...)
	//publicKey := &privateKey.PublicKey

	// GET DIFFERENT FORMATS
	/* publicKeyx509Binary, _ := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyPEMByte := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyx509Binary,
		}) */
	publicKeyHex := hex.EncodeToString(publicKeyByte)
	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKeyByte)

	//fmt.Printf("BINARY\n%b\n", publicKeyx509Binary)
	//fmt.Printf("BYTE\n%v\n", publicKeyPEMByte)
	fmt.Printf("HEX\n%v\n", publicKeyHex)
	fmt.Printf("BASE64\n%v\n", publicKeyBase64)

	return *privateKeyRAW, publicKeyByte

}

func calculateSHA256Hash(plainText string) [32]byte {

	plainTextByte := []byte(plainText)

	// HASH
	sha256HashByte := sha256.Sum256(plainTextByte)

	return sha256HashByte
}

func signatureCreator(privateKeyECDSA ecdsa.PrivateKey, plainText string) (*big.Int, *big.Int) {

	plainTextByte := []byte(plainText)

	r, s, err := ecdsa.Sign(rand.Reader, &privateKeyECDSA, plainTextByte)

	signature := append(r.Bytes(), s.Bytes()...)

	return r, s
}

func signatureVerifier(publicKeyByte []byte, plainText string, r big.Int, s big.Int) bool {

	plainTextByte := []byte(plainText)

	result := ecdsa.Verify(&publicKeyByte, plainTextByte, &r, &s)

	if result {
		return true
	} else {
		return false
	}
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

	// READ FILE INTO A STRING
	plainText := readFile(filename)
	fmt.Printf("The file %s contains:\n\n%s\n\n", filename, plainText)

	// GENERATE ECDSA KEYS
	privateKeyECDSA, publicKeyByte := generateECDSAKeys()
	fmt.Printf("The private key:\n\n%v\n\n", privateKeyECDSA)
	fmt.Printf("The public key:\n\n%v\n\n", publicKeyByte)

	// CREATE SIGNATURE
	r, s := signatureCreator(privateKeyECDSA, plainText)

	// SIGNATURE VERIFIER
	result := signatureVerifier(publicKeyByte, plainText, r, s)

	if result {
		fmt.Println("Signature Verified - Everything worked great")
	} else {
		fmt.Println("Signature NOT Verified - We don't know who this is")
	}
}
