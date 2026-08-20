package security

import "golang.org/x/crypto/bcrypt"

//Cria um Hash da senha inserida
func Hash(passwordHash string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passwordHash), bcrypt.DefaultCost)
}

// Compara a senha e o Hashs e retorna se são iguais
func PassVerifier(passHash, passString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passHash), []byte(passString))
}