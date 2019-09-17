// my-go-examples aes-cbc.go

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

func encrypt(keyByte []byte, nonce []byte, plaintext string) string {

	plaintextByte := []byte(plaintext)
	cipherTextByte := make([]byte, len(plaintext))

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	checkErr(err)

	// GET CBC ENCRYPTER
	cbc := cipher.NewCBCEncrypter(block, nonce)

	// ENCRYPT DATA
	cbc.CryptBlocks(cipherTextByte, plaintextByte)

	// RETURN HEX
	cipherText := hex.EncodeToString(cipherTextByte)
	return cipherText
}

func decrypt(keyByte []byte, nonce []byte, cipherText string) string {

	cipherTextByte, _ := hex.DecodeString(cipherText)
	plainTextByte := make([]byte, len(cipherTextByte))

	// CHECK cipherTextByte
	// CBC mode always works in whole blocks.
	if len(cipherTextByte)%aes.BlockSize != 0 {
		panic("cipherTextByte is not a multiple of the block size")
	}

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	checkErr(err)

	// GET CBC DECRYPTER
	cbc := cipher.NewCBCDecrypter(block, nonce)

	// DECRYPT DATA
	cbc.CryptBlocks(plainTextByte, cipherTextByte)

	// RETURN STRING
	plainText := string(plainTextByte[:])
	return plainText
}

func main() {

	// DATA
	// IN CBC Must be Block Size of AES (Multiple of 16)
	plainText := "This is AES-256 CBC (32 Bytes)!!"
	if len(plainText)%aes.BlockSize != 0 {
		panic("Plaintext is not a multiple of the block size")
	}
	fmt.Printf("\nOriginal Text:           %s\n\n", plainText)

	// KEY
	keyText := "myverystrongpasswordo32bitlength"
	keyByte := []byte(keyText)
	fmt.Printf("The 32-byte Key:         %s\n", keyText)

	// CREATE A NONCE
	// For this example I'm not including in the cipherText
	nonce := make([]byte, aes.BlockSize)
	// Populates the nonce with a cryptographically secure random sequence
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	fmt.Printf("The Nonce:               %x\n\n", nonce)

	// ENCRYPT
	cipherText := encrypt(keyByte, nonce, plainText)
	fmt.Printf("Encrypted Text:          %s\n", cipherText)

	// DECRYPT
	plainText = decrypt(keyByte, nonce, cipherText)
	fmt.Printf("Decrypted Text:          %s\n\n", plainText)

}
