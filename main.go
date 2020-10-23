package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user, password, dbname)

	fmt.Println("DB Connection String:", connectionString)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Connection Success.")
	a.Router = mux.NewRouter()
	log.Println("Router created.")
}

func (a *App) Run(addr string) {}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file found")
	}
}

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	a.Run(":7777")
}
