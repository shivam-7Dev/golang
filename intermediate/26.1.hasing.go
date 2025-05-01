package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func HasingInto() {
	passwodHash()
	//================Hasingh passw3ord with salt=======
	salt := generateSalt(8)
	password := "mypassword"
	hash := hashWithSalt(password, salt)

	fmt.Printf("Salt: %x\n", salt)
	fmt.Printf("Hash with salt: %x\n", hash)
	//================Hasingh passw3ord with salt=======

}

func passwodHash() {
	password := "passord!"
	hash := sha256.Sum256([]byte(password))
	fmt.Println("hash value of passowd:", hash)
	//printing hex value
	fmt.Printf("SH256: %x\n", hash)
}

func hashWithSalt(password string, salt []byte) []byte {
	data := append([]byte(password), salt...)
	hash := sha256.Sum256(data)
	return hash[:]
}

func generateSalt(lenght int) []byte {
	salt := make([]byte, lenght)
	rand.Read(salt) //secure random bytes
	return salt
}
