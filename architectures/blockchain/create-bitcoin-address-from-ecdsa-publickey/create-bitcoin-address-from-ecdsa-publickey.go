// my-go-examples create-bitcoin-address-from-ecdsa-publickey.go

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"log"
	"reflect"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func generateECDSASKeys() (string, string) {

	// GET PRIVATE & PUBLIC KEY PAIR
	curve := elliptic.P256()
	privateKeyRaw, err := ecdsa.GenerateKey(curve, rand.Reader)
	checkErr(err)

	// EXTRACT PUBLIC KEY
	publicKeyRaw := &privateKeyRaw.PublicKey

	// ENCODE
	privateKeyHex, publicKeyHex := encodeKeys(privateKeyRaw, publicKeyRaw)

	// DECODE TO CHECK
	privateKeyRawCheck, publicKeyRawCheck := decodeKeys(privateKeyHex, publicKeyHex)

	// CHECK THEY ARE THE SAME
	if !reflect.DeepEqual(privateKeyRaw, privateKeyRawCheck) {
		fmt.Println("ERROR: Private keys do not match.")
	}
	if !reflect.DeepEqual(publicKeyRaw, publicKeyRawCheck) {
		fmt.Println("ERROR: Public keys do not match.")
	}

	return privateKeyHex, publicKeyHex

}

func encodeKeys(privateKeyRaw *ecdsa.PrivateKey, publicKeyRaw *ecdsa.PublicKey) (string, string) {

	privateKeyx509Encoded, _ := x509.MarshalECPrivateKey(privateKeyRaw)
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: privateKeyx509Encoded,
		})
	fmt.Printf("\nPRIVATE ECDSA KEY (privateKeyPEM):\n%v", string(privateKeyPEM))
	privateKeyHex := hex.EncodeToString(privateKeyPEM)

	publicKeyx509Encoded, _ := x509.MarshalPKIXPublicKey(publicKeyRaw)
	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyx509Encoded,
		})
	fmt.Printf("\nPUBLIC ECDSA KEY (publicKeyPEM):\n%v", string(publicKeyPEM))
	publicKeyHex := hex.EncodeToString(publicKeyPEM)

	return privateKeyHex, publicKeyHex

}

func decodeKeys(privateKeyHex string, publicKeyHex string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {

	privateKeyPEM, _ := hex.DecodeString(privateKeyHex)
	block, _ := pem.Decode([]byte(privateKeyPEM))
	privateKeyx509Encoded := block.Bytes
	privateKeyRaw, _ := x509.ParseECPrivateKey(privateKeyx509Encoded)

	publicKeyPEM, _ := hex.DecodeString(publicKeyHex)
	blockPub, _ := pem.Decode([]byte(publicKeyPEM))
	publicKeyx509Encoded := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(publicKeyx509Encoded)
	publicKeyRaw := genericPublicKey.(*ecdsa.PublicKey)

	return privateKeyRaw, publicKeyRaw
}

func generateBitcoinAddress(publicKeyHex string) string {

	verPublicKeyHash := hashPublicKey(publicKeyHex)
	fmt.Printf("VERSIONED PUBLIC KEY HASH:    %s\n", hex.EncodeToString(verPublicKeyHash))

	checkSum := checksumKeyHash(verPublicKeyHash)
	fmt.Printf("CHECKSUM:                     %s\n", hex.EncodeToString(checkSum))

	bitcoinAddressHex := encodeKeyHash(verPublicKeyHash, checkSum)
	//fmt.Printf("\nBITCOIN ADDRESS:              %s\n\n", bitcoinAddressHex)

	return bitcoinAddressHex

}

func hashPublicKey(publicKeyHex string) []byte {

	// 1 - SHA-256 HASH
	publicKeyByte, _ := hex.DecodeString(publicKeyHex) // Don't use []byte(publicKeyHex)
	publicSHA256 := sha256.Sum256(publicKeyByte)
	fmt.Printf("    1 - SHA-256 HASH          %s\n", hex.EncodeToString(publicSHA256[:]))

	// 2 - RIPEMD-160 HASH
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	checkErr(err)
	publicKeyHash := RIPEMD160Hasher.Sum(nil)
	fmt.Printf("    2 - RIPEMD-160 HASH       %s\n", hex.EncodeToString(publicKeyHash))
	fmt.Printf("PUBLIC KEY HASH:              %s\n", hex.EncodeToString(publicKeyHash))

	// VERSION
	version := make([]byte, 1)
	version[0] = 0x00
	fmt.Printf("VERSION:                      %s\n", "00")

	// 3 - CONCAT
	verPublicKeyHash := append(version, publicKeyHash...)
	fmt.Printf("    3 - CONCAT                %s\n", hex.EncodeToString(verPublicKeyHash))

	return verPublicKeyHash
}

func checksumKeyHash(verPublicKeyHash []byte) []byte {

	// 4 - SHA-256 HASH
	firstpublicSHA256 := sha256.Sum256(verPublicKeyHash)
	fmt.Printf("    4 - SHA-256 HASH          %s\n", hex.EncodeToString(firstpublicSHA256[:]))

	// 5 - SHA-256 HASH
	secondPublicSHA256 := sha256.Sum256(firstpublicSHA256[:])
	fmt.Printf("    5 - SHA-256 HASH          %s\n", hex.EncodeToString(secondPublicSHA256[:]))

	// 6 - FIRST FOUR BYTE
	checkSum := secondPublicSHA256[:4]
	fmt.Printf("    6 - FIRST FOUR BYTE       %s\n", hex.EncodeToString(checkSum))

	return checkSum
}

func encodeKeyHash(verPublicKeyHash []byte, checkSum []byte) string {

	// 7 - CONCAT
	addressHex := append(verPublicKeyHash, checkSum...)
	fmt.Printf("    7 - CONCAT                %s\n", hex.EncodeToString(addressHex))
	fmt.Printf("BITCOIN ADDRESS HEX:          %s\n", hex.EncodeToString(addressHex))

	// 8 - BASE58 ENCODING
	bitcoinAddressHex := base58.Encode(addressHex)
	fmt.Printf("    8 - BASE58 ENCODING       %s\n", bitcoinAddressHex)

	return bitcoinAddressHex

}

func main() {

	// GENERATE ECDSA KEYS
	privateKeyHex, publicKeyHex := generateECDSASKeys()
	fmt.Printf("\nPRIVATE ECDSA KEY (privateKeyHex):\n%v\n\n", privateKeyHex)
	fmt.Printf("PUBLIC ECDSA KEY (publicKeyHex):\n%v\n\n", publicKeyHex)

	// GET BITCOIN ADDRESS
	bitcoinAddressHex := generateBitcoinAddress(publicKeyHex)
	fmt.Printf("\nBITCOIN ADDRESS:              %s\n\n", bitcoinAddressHex)

}
