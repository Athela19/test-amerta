package auth

import "golang.org/x/crypto/bcrypt"

func Verify(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
