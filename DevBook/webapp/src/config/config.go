package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//APIURL representa a URL para comunicação com a API
	APIURL 	 = ""
	//Port onde a aplicação Web está rodando
	Port 	 = 0
	//HashKey é utulizado para autenticar o cookie
	HashKey	 []byte
	//BlockKey é utilizado para criptografar os dados do cookie
	BlockKey []byte
)

//Load inicializa as variáveis de ambiente
func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}