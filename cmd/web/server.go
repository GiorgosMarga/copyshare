package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.con/GiorgosMarga/copyshare/cmd/internal/models"
	"github.con/GiorgosMarga/copyshare/cmd/internal/validator"
)

type Application struct {
	http.Server
	infoLog        *log.Logger
	snippet        *models.SnippetModel
	user           *models.UserModel
	sessionManager *scs.SessionManager
	validator      *validator.Validator
}

var errorLog *log.Logger = log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
var infoLog *log.Logger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

func NewServer(addr string, db *sql.DB) *Application {

	scsMng := scs.New()
	scsMng.Lifetime = 24 * time.Hour
	app := &Application{
		infoLog:        infoLog,
		snippet:        &models.SnippetModel{DB: db},
		user:           &models.UserModel{DB: db},
		sessionManager: scsMng,
		validator:      &validator.Validator{},
	}
	app.ErrorLog = errorLog
	app.Addr = ":" + addr // format it with :
	app.Handler = app.routes()
	return app
}

func newDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	infoLog.Printf("Conencted to database.")
	return db, nil
}
