package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JakubPluta/gocrud/config"
	"github.com/JakubPluta/gocrud/controller"
	"github.com/JakubPluta/gocrud/helpers"
	"github.com/JakubPluta/gocrud/repository"
	"github.com/JakubPluta/gocrud/router"
	"github.com/JakubPluta/gocrud/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	db, err := config.ConnectDB()
	helpers.ErrorPanic(err)
	defer db.Prisma.Disconnect()

	// repo
	postsRepository := repository.NewPostRepository(db)

	// service
	postsService := service.NewPostServiceImpl(postsRepository)

	// controller
	postsController := controller.NewPostController(postsService)

	// router
	routes := router.NewRouter(postsController)

	fmt.Printf("Listening on port %s\n", port)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())

}
