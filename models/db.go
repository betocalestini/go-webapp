package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// comentar para produção
const (
	USER   = "postgres"
	PASS   = "postgres"
	DBNAME = "postgres"
)

func Connect() *sql.DB {
	//comentar para produção
	// URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASS, DBNAME)
	// db, err := sql.Open("postgres", URL)
	//comentar para desenvolvimento
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Banco de dados conectado com sucesso!")
}
