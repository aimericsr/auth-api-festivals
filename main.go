package main

import (
	"log"

	"github.com/aimericsr/auth-api-festivals/db"
	"github.com/aimericsr/auth-api-festivals/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config variables:", err)
	}

	db.Init(config.DBSource)

}
