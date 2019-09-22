// my-go-examples md5-hash-from-file.go

package main

import (
	"crypto/md5"
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

func calculateMD5Hash(plainText string) string {

	plainTextBytes := []byte(plainText)

	// HASH
	md5HashByte := md5.Sum(plainTextBytes)

	// CONVERT TO STRING
	md5Hash := hex.EncodeToString(md5HashByte[:])

	return md5Hash
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

	// CALCULATE MD5 HASH FROM STRING
	md5Hash := calculateMD5Hash(plainText)
	fmt.Printf("\nThe md5 hash is: \n%s \n\n", md5Hash)

}
