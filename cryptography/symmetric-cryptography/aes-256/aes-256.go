// my-go-examples aes.go

package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func encrypt(keyByte []byte, plaintext string) string {

	// WILL ONLY ENCRYPT 16 BYTES (128 bits)
	plaintextByte := []byte(plaintext)
	cipherTextByte := make([]byte, 16)

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	checkErr(err)

	// ENCRYPT DATA
	block.Encrypt(cipherTextByte, plaintextByte)

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

func main() {

	// DATA
	// Must be at least 16 bytes
	plainText := "Hello Jeff, only 16 Bytes of this will be encrypted."
	fmt.Printf("\nOriginal Text:           %s\n\n", plainText)

	// KEY
	keyText := "myverystrongpasswordo32bitlength"
	keyByte := []byte(keyText)
	fmt.Printf("The 32-byte Key:         %s\n\n", keyText)

	// ENCRYPT
	cipherText := encrypt(keyByte, plainText)
	fmt.Printf("Encrypted Text:          %s\n", cipherText)

	// DECRYPT
	plainText = decrypt(keyByte, cipherText)
	fmt.Printf("Decrypted Text:          %s\n\n", plainText)

}
