// my-go-examples sha256-hash-from-file.go

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func calculateSHA256Hash(plainText string) string {

	plainTextBytes := []byte(plainText)

	// HASH
	sha256HashByte := sha256.Sum256(plainTextBytes)

	// CONVERT TO STRING
	sha256Hash := hex.EncodeToString(sha256HashByte[:])

	return sha256Hash
}

func main() {

	fmt.Println(" ")

	// GET FILE NAME FROM ARGS
	flag.Parse()
	filenameSlice := flag.Args()
	if len(filenameSlice) != 1 {
		err := errors.New("Only one filename allowed")
		checkErr(err)
	}
	filename := filenameSlice[0] // Make it a string

	// READ FILE INTO BYTES
	plainTextByte, err := ioutil.ReadFile(filename)
	checkErr(err)
	// Convert to string
	plainText := string(plainTextByte)
	fmt.Printf("The file %s contains: \n%s\n", filename, plainText)

	// CALCULATE SHA256 HASH FROM STRING
	sha256Hash := calculateSHA256Hash(plainText)
	fmt.Printf("\nThe sha256 hash is: \n%s \n\n", sha256Hash)

}
