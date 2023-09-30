package main

import (
	"log"
)

// func init() {
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func main() {
	dsn := "pp2y8t0tscjkn1sw1wop:pscale_pw_1qn9J7sFP9bFBXG6VzaZWJ7JKnuLNJwaZCvmoZGVQ45@tcp(aws.connect.psdb.cloud)/copyshare?tls=true&interpolateParams=true&parseTime=true"
	port := "3000"
	db, err := newDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	app := NewServer(port, db)
	defer db.Close()
	app.infoLog.Printf("Starting server on port %s", app.Addr)
	log.Fatal(app.ListenAndServe())

}
