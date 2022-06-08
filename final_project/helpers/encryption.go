package helpers

import "golang.org/x/crypto/bcrypt"

func PasswordEncryption(password string) (string, error) {
	encPass, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(encPass), err
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
