package main

import (
	"database/sql"
	"fmt"
	"hotel-lister/controllers"
	"hotel-lister/database"
	"hotel-lister/middleware"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	DB  *sql.DB
	err error
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
	v1 := r.Group("/v1")

	v1.Use(middleware.Auth)

	//user
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	v1.GET("/users", controllers.GetAllUsers)
	v1.DELETE("/user/:id", controllers.DeleteUser)

	//countries
	v1.GET("/countries", controllers.GetCountries)        // v
	v1.POST("/country", controllers.InsertCountries)      // v
	v1.PUT("/country/:id", controllers.UpdateCountries)   // v
	v1.DELETE("country/:id", controllers.DeleteCountries) // v
	v1.GET("/country/:id/hotels", controllers.GetHotelByCountries)

	//cities
	v1.GET("/cities", controllers.GetCities)        // v
	v1.POST("/city", controllers.InsertCities)      // v
	v1.PUT("/city/:id", controllers.UpdateCities)   // v
	v1.DELETE("city/:id", controllers.DeleteCities) // v
	v1.GET("/city/:id/hotels", controllers.GetHotelByCities)

	//hotels
	v1.GET("/hotels", controllers.GetAllHotels)      // v
	v1.POST("/hotel", controllers.InsertHotels)      // v
	v1.PUT("/hotel/:id", controllers.UpdateHotels)   // v
	v1.DELETE("hotel/:id", controllers.DeleteHotels) // v

	//rooms
	v1.GET("/hotel/:id/rooms", controllers.GetRoomByHotels) // v
	v1.POST("/room", controllers.InsertRooms)               // v validate duplicate id
	v1.PUT("/room/:id", controllers.UpdateRooms)            // v
	v1.DELETE("room/:id", controllers.DeleteRooms)          // v

	//review
	v1.GET("/reviews", controllers.GetAllReviews)      // v validate duplicate id
	v1.POST("/review", controllers.InsertReviews)      // v limit 1-5
	v1.PUT("/review/:id", controllers.UpdateReviews)   // v
	v1.DELETE("review/:id", controllers.DeleteReviews) // v
	v1.GET("/hotel/:id/reviews", controllers.GetReviewByHotel)

	//run app
	r.Run(":" + os.Getenv("PORT"))
}
