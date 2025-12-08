package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     int
	Available bool
}

func main() {
	connStr := "postgres://postgres:Admin@127.0.0.1:5433/gopgtest?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	createProductTable(db)

}

func createProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(6,2) NOT NULL,
    available BOOLEAN,
    created timestamp DEFAULT NOW()
    )`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	product := Product{"Boo", 14, false}
	pk := insertProduct(db, product)
	fmt.Printf("ID:%d", pk)

}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product(name , price ,available)
		VALUES ($1,$2,$3) RETURNING id
		`
	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
