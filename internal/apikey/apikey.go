package apikey

import (
	"errors"
	"log"
	"os"
)

func Read() (string, string, error) {
	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		log.Fatalln("API_KEY environment variable not set")
		return "", "", errors.New("empty API_KEY")
	}
	SECRET_KEY := os.Getenv("SECRET_KEY")
	if SECRET_KEY == "" {
		log.Fatalln("SECRET_KEY environment variable not set")
		return "", "", errors.New("empty SECRET_KEY")
	}

	return API_KEY, SECRET_KEY, nil
}
