package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func encrypt(keyByte []byte, plaintext string, additionalData string) string {

	plaintextByte := []byte(plaintext)
	additionalDataByte := []byte(additionalData)

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	checkErr(err)

	// GET GCM INSTANCE THAT USES THE AES CIPHER
	gcm, err := cipher.NewGCM(block)
	checkErr(err)

	// CREATE A NONCE
	nonce := make([]byte, gcm.NonceSize())
	// Populates the nonce with a cryptographically secure random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	fmt.Printf("The Nonce:               %x\n\n", nonce)

	// ENCRYPT DATA
	// Note how we put the Nonce in the beginging,
	// So we can rip it out when we decrypt
	cipherTextByte := gcm.Seal(nonce, nonce, plaintextByte, additionalDataByte)

	// RETURN HEX
	cipherText := hex.EncodeToString(cipherTextByte)
	return cipherText
}

func decrypt(keyByte []byte, cipherText string, additionalData string) string {

	cipherTextByte, _ := hex.DecodeString(cipherText)
	additionalDataByte := []byte(additionalData)

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	checkErr(err)

	// GET GCM BLOCK
	gcm, err := cipher.NewGCM(block)
	checkErr(err)

	// EXTRACT NONCE FROM cipherTextByte
	// Because I put it there
	nonceSize := gcm.NonceSize()
	nonce, cipherTextByte := cipherTextByte[:nonceSize], cipherTextByte[nonceSize:]

	// DECRYPT DATA
	plainTextByte, err := gcm.Open(nil, nonce, cipherTextByte, additionalDataByte)
	checkErr(err)

	// RETURN STRING
	plainText := string(plainTextByte[:])
	return plainText
}

func main() {

	// DATA
	plainText := "This is an example of AES-256 GCM."
	fmt.Printf("\nOriginal Text:           %s\n\n", plainText)

	// KEY
	keyText := "myverystrongpasswordo32bitlength"
	keyByte := []byte(keyText)
	fmt.Printf("The 32-byte Key:         %s\n", keyText)

	// ADDITIONAL DATA (NON ENCRYPTED AUTHORIZATION)
	additionalData := "Jeff's additional data for authorization"
	fmt.Printf("Additional Data:         %s\n", additionalData)

	// ENCRYPT
	cipherText := encrypt(keyByte, plainText, additionalData)
	fmt.Printf("Encrypted Text:          %s\n", cipherText)

	// DECRYPT
	plainText = decrypt(keyByte, cipherText, additionalData)
	fmt.Printf("Decrypted Text:          %s\n\n", plainText)

}
