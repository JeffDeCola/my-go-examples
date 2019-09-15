package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"
)

// Check your error
func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func encrypt(key []byte, plaintext string) string {

	// WILL ONLY ENCRYPT 16 BYTES (128 bits)
	plaintextByte := []byte(plaintext)
	cipherTextByte := make([]byte, 16)

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(key)
	checkErr(err)

	// ENCRYPT DATA - PLACE IN cipherTextByte
	block.Encrypt(cipherTextByte, plaintextByte)

	// RETURN HEX
	cipherText := hex.EncodeToString(cipherTextByte)
	return cipherText
}

func decrypt(key []byte, cipherText string) string {

	// IT IS 16 BYTES (128 bits)
	cipherTextByte, _ := hex.DecodeString(cipherText)
	plainTextByte := make([]byte, 16)

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(key)
	checkErr(err)

	// DECRYPT DATA - PLACE IN plainTextByte
	block.Decrypt(plainTextByte, cipherTextByte)

	// RETRUN STRING
	plainText := string(plainTextByte[:])
	return plainText
}

func main() {

	// KEY
	keyText := "myverystrongpasswordo32bitlength"
	key := []byte(keyText)

	// DATA
	// Must be at least 16 BYTES
	plainText := "Hello Jeff, only 16 Bytes of this will be encrypted"
	fmt.Printf("Original Text:   %s\n", plainText)

	// ENCRYPT
	cipherText := encrypt(key, plainText)
	fmt.Printf("Encrypted Text:  %s\n", cipherText)

	// DECRYPT
	plainText = decrypt(key, cipherText)
	fmt.Printf("Decrypted Text:  %s\n", plainText)

}
