package database

import (
	"context"
	"fmt"
	"log"
	"todo/ent"
	"todo/internal/config"
)

// DBClient global database context to use
var DBClient *ent.Client
var DBCtx context.Context

// Connect connects to the database
func Connect(migrate bool) {

	connectionStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		config.Settings.Database.Host, config.Settings.Database.Port, config.Settings.Database.Name,
		config.Settings.Database.User, config.Settings.Database.Pwd)

	log.Println(connectionStr)

	var err error
	// instantiate db context and client
	DBCtx = context.Background()
	DBClient, err = ent.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("error connecting to database \n%v", err)
	}
	log.Println("connected to database")

	if migrate {
		runMigration()
	}

}

func runMigration() {
	log.Println("running migration...")
	if err := DBClient.Schema.Create(DBCtx); err != nil {
		log.Fatalf("error creating database schema. \n%v", err)
	}
}
