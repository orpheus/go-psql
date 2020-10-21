package main

import (
	"database/sql"
	"github.com/gorilla/mux"
)

type Service struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *Service) Initialize(user, password, dbname string) { }

func (a *Service) Run(addr string) { }

func main() {
	s := Service{}
	s.Initialize("postgres","postgres", "wsdb")
	s.Run(":7777")

	//db := connectToPostgresDb()
	//startServer(db)
}
