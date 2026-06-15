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

	// type User struct {
	// 	ID    int
	// 	Name  string
	// 	Email string
	// }

	// // SELECT (DQL)//
	// _, err = conn.Prepare(
	// 	context.Background(),
	// 	"get-users",
	// 	"SELECT * FROM users WHERE email LIKE '%gigpz%'",
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// rows, _ := conn.Query(
	// 	context.Background(),
	// 	"get-users",
	// )

	// for rows.Next() {
	// 	var user User

	// 	rows.Scan(&user.ID, &user.Name, &user.Email)
	// 	fmt.Println(user)
	// }

	// transacciones //
	tx, err := conn.Begin(context.Background())
	if err != nil {
		panic(err)
	}

	defer tx.Rollback(context.Background())

	_, err = tx.Exec(
		context.Background(),
		"UPDATE accounts SET balance=balance-100 WHERE id=1",
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.Exec(
		context.Background(),
		"UPDATE accounts SET balance=balance+100 WHERE id=2",
	)

	if err != nil {
		panic(err)
	}

	tx.Commit(context.Background())
	fmt.Println("Transacciones completada correctamente")

	// type UserPost struct {
	// 	UserName  string
	// 	PostTitle string
	// }

	// // SELECT (DQL)//
	// _, err = conn.Prepare(
	// 	context.Background(),
	// 	"get-users-posts",
	// 	"SELECT users.name, posts.title FROM users INNER JOIN posts ON users.id = posts.user_id;",
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// rows, _ := conn.Query(
	// 	context.Background(),
	// 	"get-users-posts",
	// )

	// for rows.Next() {
	// 	var userPost UserPost

	// 	rows.Scan(&userPost.UserName, &userPost.PostTitle)
	// 	fmt.Println(userPost)
	// }

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

	// var total int
	// _, err = conn.Prepare(
	// 	context.Background(),
	// 	"get-total",
	// 	"SELECT COUNT(*) FROM users;",
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// conn.QueryRow(
	// 	context.Background(),
	// 	"get-total",
	// ).Scan(&total)

	// fmt.Printf("Total users: %d\n", total)

	defer conn.Close(context.Background())

	//fmt.Println("Conectado correctamente, tabla creada")
	//fmt.Println("Conectado correctamente, fila insertada")
}
