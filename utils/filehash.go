package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func Sha256file(filename string) string {
	hasher := sha256.New()

	fileHandler, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandler.Close()
	if _, err := io.Copy(hasher, fileHandler); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}
