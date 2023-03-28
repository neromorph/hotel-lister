package main

import (
	"database/sql"
	"fmt"
	"hotel-lister/database"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// ENV
	err = godotenv.Load("configs/.env")
	if err != nil {
		fmt.Println("failed to load environment")
	} else {
		fmt.Println("success read file environment")
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
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	r := gin.Default()
	r.Run(":" + os.Getenv("PORT"))
}
