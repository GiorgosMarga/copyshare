package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dsn := os.Getenv("DSN")
	port := os.Getenv("PORT")
	db, err := newDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	app := NewServer(port, db)
	defer db.Close()
	app.infoLog.Printf("Starting server on port %s", app.Addr)
	log.Fatal(app.ListenAndServe())

}
