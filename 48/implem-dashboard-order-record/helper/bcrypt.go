package helper

import (
	// "crypto/rand"
	// "encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(inputPassword, storedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func HashPasswordSeed(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func CompareHashAndPassword(pass1, pass2 string) error {
	return bcrypt.CompareHashAndPassword([]byte(pass1), []byte(pass2))
}

// func GenerateToken() string {
// 	token := make([]byte, 32)
// 	_, err := rand.Read(token)
// 	if err != nil {
// 		panic("failed to generate token")
// 	}
// 	return hex.EncodeToString(token)
// }
