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

	sql := `
	CREATE TABLE IF NOT EXISTS users2(
		id SERIAL PRIMARY KEY,
		name VARCHAR(100),
		email VARCHAR(100)
	)
	`

	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())
	fmt.Println("Conectado correctamente, tabla creada")
}
