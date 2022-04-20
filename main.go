package main

import (
	"log"

	"github.com/aimericsr/auth-api-festivals/util"
)

func main() {
	_, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config variables:", err)
	}
}
