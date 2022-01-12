package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"

	"github.com/google/uuid"
)

func HashGen(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func idGenarator() string {
	return uuid.New().String()
}
func log_excepts(_log string) {
	path := func() string {
		if GetProccessMode() == "development" {
			return "./logs/user-logs.log"
		} else {
			return "/data/logs/user-logs.log"
		}
	}()
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	defer f.Close()
	_, err = f.WriteString(_log)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatalf("sentry.Init: %s", e)
	}
}
