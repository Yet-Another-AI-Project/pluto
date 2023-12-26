package main

import (
	"log"
	"os"

	"pluto/config"

	"pluto/database"

	"pluto/utils/migrate"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.NewConfig(os.Args, "")

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDatabase(cfg)

	if err != nil {
		log.Fatal(err)
	}

	if err := migrate.Migrate(db); err != nil {
		log.Fatal(err)
	}

	log.Println("migrate success")
}
