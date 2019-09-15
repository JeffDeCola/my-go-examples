package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
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

// ENCRYPT DATA WITH 32 BYTE KEY AND RETURN CIPHERTEXT
func encrypt(data []byte, hashKey string) []byte {

	// generate a new aes cipher using our 32 byte long key
	block, err := aes.NewCipher([]byte(hashKey))
	checkErr(err)
	gcm, err := cipher.NewGCM(block)
	checkErr(err)

	// Creates a new byte array the size of the nonce
	nonce := make([]byte, gcm.NonceSize())

	// Populates our nonce with a cryptographically secure random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// Encrypt our text using the Seal function
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
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
		err := errors.New("only two file allowed")
		checkErr(err)
	}
	filename := filenameSlice[0]    // Make it a string
	filenameout := filenameSlice[1] // Make it a string

	// READ THE FILE- Will be a slice of bytes
	log.Trace("Read the file to encrypt")
	fileDataToEncrypt, err := ioutil.ReadFile(filename)
	checkErr(err)
	// fmt.Printf("Data/File to encrypt\n--------------------\n%s\n--------------------\n", fileDataToEncrypt)

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

	// ENCRYPT DATA WITH KEY
	log.Trace("Encrypt file with key")
	fmt.Println("Encrypting input file")
	cipherText := encrypt(fileDataToEncrypt, hashKey)
	s := string(cipherText[:])
	fmt.Printf("Encrypted Data\n--------------------\n%s\n--------------------\n", s)
	//You will get back a slice of bytes
	fmt.Printf("Encrypted Data\n--------------------\n%x\n--------------------\n", cipherText)
	fmt.Printf("Encrypted Data\n--------------------\n%v\n--------------------\n", cipherText)
	fmt.Printf("Encrypted Data\n--------------------\n%s\n--------------------\n", cipherText)

	// WRITE TO A FILE
	log.Trace("Write cipherText to a file")
	f, err := os.Create(filenameout)
	checkErr(err)
	defer f.Close()
	f.Write(cipherText)
	fmt.Println("Wrote output file")

	ioutil.WriteFile("myfile.data", cipherText, 0777)
}
