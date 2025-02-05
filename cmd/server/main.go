package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rlnorthcutt/heal/internal/db"
	"github.com/rlnorthcutt/heal/internal/handlers"
)

func main() {
	e := echo.New()
	db.InitDB() // Initialize the database

	// Static file serving
	e.Static("/static", "static")

	// Routes
	e.GET("/", handlers.Home)
	e.GET("/content", handlers.ListContent)
	e.POST("/content", handlers.CreateContent)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
