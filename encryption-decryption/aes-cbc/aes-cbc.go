package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {

	// Must Kept Secret No Hardcoding , This is for Demo purpose.
	key := "myverystrongpasswordo32bitlength"
	plainText := "Hello 8gwifi.org"
	fmt.Printf("Original Text:  %s\n", plainText)
	fmt.Println()
	fmt.Println("====GCM Encryption/ Decryption Without AAD====")

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	iv := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err.Error())
	}
	ciphertext := GCM_encrypt(key, plainText, iv, nil)
	fmt.Printf("GCM Encrypted Text:  %s\n", ciphertext)
	ret := GCM_decrypt(key, ciphertext, iv, nil)
	fmt.Printf("GCM Decrypted Text:  %s\n", ret)
	fmt.Println()

	fmt.Println("====GCM Encryption/ Decryption Using AAD====")

	// Never Use Same IV or Nonce
	iv = make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err.Error())
	}
	additionalData := "Not Secret AAD Value"
	additionalData2 := "Not Shecret AAD Value"
	ciphertext = GCM_encrypt(key, plainText, iv, []byte(additionalData))
	fmt.Printf("GCM Encrypted Text:  %s\n", ciphertext)
	ret = GCM_decrypt(key, ciphertext, iv, []byte(additionalData2))

	fmt.Printf("GCM Decrypted Text:  %s\n", ret)

}

func GCM_encrypt(key string, plaintext string, iv []byte, additionalData []byte) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, iv, []byte(plaintext), additionalData)
	fmt.Printf("%x\n", ciphertext)
	return hex.EncodeToString(ciphertext)
}

func GCM_decrypt(key string, ct string, iv []byte, additionalData []byte) string {
	ciphertext, _ := hex.DecodeString(ct)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, iv, ciphertext, additionalData)
	if err != nil {
		panic(err.Error())
	}
	s := string(plaintext[:])
	return s
}
