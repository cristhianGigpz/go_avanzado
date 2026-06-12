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
	type User struct {
		ID    int
		Name  string
		Email string
	}

	_, err = conn.Prepare(
		context.Background(),
		"get-users",
		"SELECT * FROM users WHERE email LIKE '%gigpz%'",
	)
	if err != nil {
		panic(err)
	}

	rows, _ := conn.Query(
		context.Background(),
		"get-users",
	)

	for rows.Next() {
		var user User

		rows.Scan(&user.ID, &user.Name, &user.Email)
		fmt.Println(user)
	}

	// SELECT (DQL)//

	// var user User

	// _, err = conn.Prepare(
	// 	context.Background(),
	// 	"get-user-by-id",
	// 	"SELECT id,name,email FROM users WHERE email=$1",
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// conn.QueryRow(
	// 	context.Background(),
	// 	"get-user-by-id",
	// 	"gigpz@example.com",
	// ).Scan(&user.ID, &user.Name, &user.Email)

	// fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)

	defer conn.Close(context.Background())

	//fmt.Println("Conectado correctamente, tabla creada")
	//fmt.Println("Conectado correctamente, fila insertada")
}
