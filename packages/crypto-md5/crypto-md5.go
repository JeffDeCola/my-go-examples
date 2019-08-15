package main

import (
	"crypto/md5"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println(" ")

	// Flags - Will default to $HOME/.ssh/id_rsa.pub if no input giving
	inputFilenamePathPtr := flag.String("i", os.Getenv("HOME")+"/.ssh/id_rsa.pub", "input file")
	// Parse the flags
	flag.Parse()

	// Read key file
	fmt.Printf("Reading the keyfile %s\n\n", *inputFilenamePathPtr)
	keyfile, err := ioutil.ReadFile(*inputFilenamePathPtr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The keyfile contains: \n%s \n", keyfile)

	// Parse the file because the file looks like `ssh-rsa AAA...ABC comments`
	// Hence parts[1] is the key
	fmt.Printf("Parsing the keyfile...")
	parts := strings.Fields(string(keyfile))
	if len(parts) < 2 {
		log.Fatal("bad parse")
	}
	fmt.Printf("The key is: \n%s \n\n", parts[1])

	// Turn the string into bytes
	fmt.Printf("Turn key into bytes...")
	keyBytes, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The key in bytes is now: \n%x \n\n", keyBytes)

	// Get the hash in md5 bytes
	fmt.Printf("Hash the key (MD5 fingerprint)...")
	md5HashInBytes := md5.Sum([]byte(keyBytes))
	fmt.Printf("The key hashed is now: \n%x \n\n", md5HashInBytes)

	// Print out the md5 fingerprint
	fmt.Println("The md5 fingerprint in a more readable form:")
	for i, b := range md5HashInBytes {
		fmt.Printf("%02x", b)
		if i < len(md5HashInBytes)-1 {
			fmt.Print(":")
		}
	}

	fmt.Printf("\n\n")

}
