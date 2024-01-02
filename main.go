package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JakubPluta/gocrud/config"
	"github.com/JakubPluta/gocrud/helpers"
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

	fmt.Printf("Listening on port %s\n", port)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())

}
