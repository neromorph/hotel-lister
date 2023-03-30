package routers

import (
	"hotel-lister/controllers"
	"hotel-lister/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/auth")
	v1 := r.Group("/v1")

	v1.Use(middleware.Auth)

	//user
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
	v1.GET("/users", controllers.GetAllUsers)
	v1.DELETE("/user/:id", controllers.DeleteUser)

	//countries
	v1.GET("/countries", controllers.GetCountries)
	v1.POST("/country", controllers.InsertCountries)
	v1.PUT("/country/:id", controllers.UpdateCountries)
	v1.DELETE("country/:id", controllers.DeleteCountries)
	v1.GET("/country/:id/hotels", controllers.GetHotelByCountries)

	//cities
	v1.GET("/cities", controllers.GetCities)
	v1.POST("/city", controllers.InsertCities)
	v1.PUT("/city/:id", controllers.UpdateCities)
	v1.DELETE("city/:id", controllers.DeleteCities)
	v1.GET("/city/:id/hotels", controllers.GetHotelByCities)

	//hotels
	v1.GET("/hotels", controllers.GetAllHotels)
	v1.POST("/hotel", controllers.InsertHotels)
	v1.PUT("/hotel/:id", controllers.UpdateHotels)
	v1.DELETE("hotel/:id", controllers.DeleteHotels)

	//rooms
	v1.GET("/hotel/:id/rooms", controllers.GetRoomByHotels)
	v1.POST("/room", controllers.InsertRooms)
	v1.PUT("/room/:id", controllers.UpdateRooms)
	v1.DELETE("room/:id", controllers.DeleteRooms)

	//review
	v1.GET("/reviews", controllers.GetAllReviews)
	v1.POST("/review", controllers.InsertReviews)
	v1.PUT("/review/:id", controllers.UpdateReviews)
	v1.DELETE("review/:id", controllers.DeleteReviews)
	v1.GET("/hotel/:id/reviews", controllers.GetReviewByHotel)

	r.Run(":" + os.Getenv("PORT"))

	return r
}
