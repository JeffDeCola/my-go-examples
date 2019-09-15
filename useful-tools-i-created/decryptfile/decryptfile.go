package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// HASH THE PARAPHRASE TO GET 32 BYTE KEY
func createKey(paraphrase string) (string, error) {
	log.Trace("hashing the paraphrase")
	hasher := md5.New()
	hasher.Write([]byte(paraphrase))
	hash := hex.EncodeToString(hasher.Sum(nil))
	log.Trace("hash is ", hash)
	return hash, nil
}

// DECRYPT DATA WITH 32 BYTE KEY AND RETURN PLAINTEXT
func decrypt(data []byte, hashKey string) []byte {

	key := []byte(hashKey)
	block, err := aes.NewCipher(key)
	checkErr(err)

	gcm, err := cipher.NewGCM(block)
	checkErr(err)

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	checkErr(err)

	return plaintext
}

// Check your error
func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func init() {

	// SET LOG LEVEL
	log.SetLevel(log.InfoLevel)
	// log.SetLevel(log.TraceLevel)

	// SET FORMAT
	log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	log.SetOutput(os.Stdout)

}

func main() {

	// GET FILE NAME
	flag.Parse()
	filenameSlice := flag.Args()
	if len(filenameSlice) != 2 {
		err := errors.New("only two files allowed")
		checkErr(err)
	}
	filename := filenameSlice[0]    // Make it a string
	filenameout := filenameSlice[1] // Make it a string

	// READ THE FILE- Will be a slice of bytes
	log.Trace("Read the file to decrypt")
	fileDataToDecrypt, err := ioutil.ReadFile(filename)
	checkErr(err)
	// fmt.Printf("Data/File to decrypt\n--------------------\n%x\n--------------------\n", fileDataToDecrypt)

	// GET THE PARAPHRASE
	log.Trace("Get the paraphrase")
	paraphrase := ""
	fmt.Printf("What is your secret paraphrase? ")
	_, err = fmt.Scan(&paraphrase)
	checkErr(err)

	// HASH THE PARAPHRASE TO GET 32 BYTE KEY STRING
	log.Trace("hash the paraphrase to get 32 byte key")
	hashKey, err := createKey(paraphrase)
	checkErr(err)

	// DECRYPT
	log.Trace("Decrypt fileDataToDecrypt with key")
	fmt.Println("Decrypting input file")
	plainText := decrypt(fileDataToDecrypt, hashKey)
	// fmt.Printf("Decrypted Data\n--------------------\n%s\n--------------------\n", plainText)

	// WRITE TO A FILE (ADD DATE)
	log.Trace("Write plainText to a file")
	f, err := os.Create(filenameout)
	checkErr(err)
	defer f.Close()
	f.Write(plainText)
	fmt.Println("Wrote output file")

}
