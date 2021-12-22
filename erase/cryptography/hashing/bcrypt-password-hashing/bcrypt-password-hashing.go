// my-go-examples cryptography bcrypt-password-hashing.go

package main

import (
	"fmt"

	bcrypt "golang.org/x/crypto/bcrypt"
)

// hashPassword using bcrypt
func hashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

// checkPasswordHash using bcrypt
func checkPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

func main() {

	// SET PASSWORD -----------------------------------------------------------------------------
	passwordSet := "myPassword"

	// CREATE PASSWORD HASH USING BCRYPT
	passwordHash, _ := hashPassword(passwordSet)

	fmt.Println("Password: ", passwordSet)
	fmt.Println("Hash:     ", passwordHash)

	// USE PASSWORD -----------------------------------------------------------------------------
	passwordEntered := "myPassword"

	// CHECK PASSWORD HASH AGAINST PASSWORD USING BCRYPT
	status := checkPasswordHash(passwordEntered, passwordHash)

	if status {
		fmt.Println("Congrats, it's you!!")
	} else {
		fmt.Println("Invalid Password")
	}

}
