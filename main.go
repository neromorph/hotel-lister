package main

import (
	"database/sql"
	"fmt"
	"hotel-lister/database"
	"hotel-lister/routers"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// ENV
	err = godotenv.Load("configs/.env")
	if err != nil {
		fmt.Println("environment failed to load")
	} else {
		fmt.Println("environment loaded")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Database failed to connect")
		panic(err)
	} else {
		fmt.Println("Database connected")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	//run app
	routers.Server()
}
