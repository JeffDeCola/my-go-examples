package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
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

	// Generate a new aes cipher using our 32 byte long key
	key := []byte(hashKey)
	block, err := aes.NewCipher(key)
	checkErr(err)

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	gcm, err := cipher.NewGCM(block)
	checkErr(err)

	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	checkErr(err)

	return plaintext
}

func getCipherText(inputFile *os.File) []byte {

	cipherTextHex := ""

	// Start scanning the input file line by line
	scanner := bufio.NewScanner(inputFile) // Increment the token

	for scanner.Scan() {

		// Read a line in file
		line := scanner.Text()

		// If you find a delimiter, get all the lines in between and place in a table.
		if line == "--------------------------------------------------------------------------------" {

			// Stay in here until you find another delimiter
			for scanner.Scan() {

				// Read next line nad Start Building the long long string
				line = scanner.Text()

				// Exit and build table when you find another delimiter
				if line == "--------------------------------------------------------------------------------" {
					break
				}
				cipherTextHex = cipherTextHex + line
			}
		}
	}

	cipherText, err := hex.DecodeString(cipherTextHex)
	checkErr(err)
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
		err := errors.New("only two files allowed")
		checkErr(err)
	}
	filename := filenameSlice[0]    // Make it a string
	filenameout := filenameSlice[1] // Make it a string

	// READ THE FILE- Will be a slice of bytes
	log.Trace("Read the file to decrypt")
	// Open input file
	inputFile, err := os.Open(filename)
	checkErr(err)
	defer inputFile.Close()

	// GET CIPHERTEXT (in bytes) FROM INPUT FILE
	cipherText := getCipherText(inputFile)
	// fmt.Printf("Data/File to decrypt\n--------------------\n%x\n--------------------\n", cipherText)

	// GET THE PARAPHRASE - ASK USER
	log.Trace("Get the paraphrase")
	paraphrase := ""
	fmt.Printf("\nWhat is your secret paraphrase? ")
	_, err = fmt.Scan(&paraphrase)
	checkErr(err)

	// HASH THE PARAPHRASE TO GET 32 BYTE KEY STRING
	log.Trace("hash the paraphrase to get 32 byte key")
	hashKey, err := createKey(paraphrase)
	checkErr(err)

	// DECRYPT
	log.Trace("Decrypt cipherText with key")
	fmt.Println("Decrypting input file")
	plainText := decrypt(cipherText, hashKey)
	// fmt.Printf("Decrypted Data\n--------------------\n%s\n--------------------\n", plainText)

	// WRITE TO A FILE (ADD DATE)
	log.Trace("Write plainText to a file")
	f, err := os.Create(filenameout)
	checkErr(err)
	defer f.Close()
	f.Write(plainText)
	fmt.Printf("Wrote output file\n\n")

}
