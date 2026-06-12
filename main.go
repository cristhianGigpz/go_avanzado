package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func main() {
	fmt.Println("Hello World, GO avanzado !")
	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:gigpz@localhost:5434/bd_tests",
	)

	if err != nil {
		panic(err)
	}

	// CREATE TABLE (DDL)//
	// sql := `
	// CREATE TABLE IF NOT EXISTS users(
	// 	id SERIAL PRIMARY KEY,
	// 	name VARCHAR(100),
	// 	email VARCHAR(100)
	// )
	// `

	// INSERT (DML)//
	// insert := `INSERT INTO users(id, name, email) VALUES (2,'Joel', 'joel@example.com')`

	// _, err = conn.Exec(context.Background(), insert)
	// if err != nil {
	// 	panic(err)
	// }

	// UPDATE (DML)//
	// var id int = 1

	// _, err = conn.Exec(context.Background(), "UPDATE users SET name = 'Gigpz' WHERE id = $1", id)
	// if err != nil {
	// 	panic(err)
	// }

	// DELETE (DML)//
	// _, err = conn.Exec(
	// 	context.Background(),
	// 	"DELETE FROM users WHERE id=$1",
	// 	2,
	// )

	// SELECT (DQL)//
	// rows, err := conn.Query(context.Background(), "SELECT id, name, email FROM users")
	// if err != nil {
	// 	panic(err)
	// }

	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	var email string

	// 	rows.Scan(&id, &name, &email)
	// 	fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	// }

	// SELECT (DQL)//
	var id int
	var name string
	var email string

	err = conn.QueryRow(
		context.Background(),
		"SELECT id,name,email FROM users WHERE id=$1",
		2,
	).Scan(&id, &name, &email)

	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)

	defer conn.Close(context.Background())

	//fmt.Println("Conectado correctamente, tabla creada")
	//fmt.Println("Conectado correctamente, fila insertada")
}
