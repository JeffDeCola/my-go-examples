// my-go-examples ecdsa-asymmetric-cryptography.go

package main

import (
	"crypto/aes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
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

func generateECDSAKeys() ([]byte, []byte) {

	// GENERATE PUBLIC/PRIVATE KEY PAIR (32 Bytes/256 bits)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	checkErr(err)

	// GET DIFFERENT FORMATS
	privateKeys509Binary, _ := x509.MarshalECPrivateKey(privateKey)
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
	publicKey := &privateKey.PublicKey

	// GET DIFFERENT FORMATS
	publicKeyx509Binary, _ := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyPEMByte := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyx509Binary,
		})
	publicKeyHex := hex.EncodeToString(publicKeyPEMByte)
	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKeyPEMByte)

	fmt.Printf("BINARY\n%b\n", publicKeyx509Binary)
	fmt.Printf("BYTE\n%v\n", publicKeyPEMByte)
	fmt.Printf("HEX\n%v\n", publicKeyHex)
	fmt.Printf("BASE64\n%v\n", publicKeyBase64)

	return privateKeyPEMByte, publicKeyPEMByte

}

func encrypt(keyByte []byte, plainText string) string {

	// WILL ONLY ENCRYPT 16 BYTES (128 bits)
	plainTextByte := []byte(plainText)
	cipherTextByte := make([]byte, 16)

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	checkErr(err)

	// ENCRYPT DATA
	block.Encrypt(cipherTextByte, plainTextByte)

	// RETURN HEX
	cipherText := hex.EncodeToString(cipherTextByte)
	return cipherText
}

func decrypt(keyByte []byte, cipherText string) string {

	// IT IS 16 BYTES (128 bits)
	cipherTextByte, _ := hex.DecodeString(cipherText)
	plainTextByte := make([]byte, 16)

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	checkErr(err)

	// DECRYPT DATA
	block.Decrypt(plainTextByte, cipherTextByte)

	// RETURN STRING
	plainText := string(plainTextByte[:])
	return plainText
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
	fmt.Printf("The file %s contains:\n\n%s\n\n", filename, plainText)

	// GENERATE ECDSA KEYS
	privateKeyPEMByte, publicKeyPEMByte := generateECDSAKeys()

	// ENCRYPT
	cipherText := encrypt(privateKeyPEMByte, plainText)

	// DECRYPT
	plainText = decrypt(keyByte, cipherText)

	// VERIFY

}
