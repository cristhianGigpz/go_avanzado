package main

import (
	"errors"
	"fmt"
	"go-avanzado/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const dsn = "host=localhost user=postgres password=gigpz dbname=bd_tests port=5434 sslmode=disable"

type Post struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	UserID uint
}

type Role struct {
	ID   uint
	Name string
}

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100"`
	Email string `gorm:"uniqueIndex;size:100"`
	Age   int    `gorm:"default:18"`
	Posts []Post `gorm:"foreignKey:UserID"`
	//Roles     []Role `gorm:"many2many:user_roles"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type user_roles struct {
	UserID uint
	RoleID uint
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	fmt.Println("Creando usuario")

	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {

	fmt.Println("Usuario creado")

	return nil
}

func Adults(db *gorm.DB) *gorm.DB {
	return db.Where("age >= ?", 18)
}

func Init() {
	r := gin.New()

	//r.Use(MyMiddleware())
	r.Use(
		middleware.LoggerMiddleware(),
		middleware.RecoveryMiddleware(),
		middleware.CORSMiddleware(),
		middleware.RateLimitMiddleware(),
	)
	//r.Use(ErrorMiddleware())
	// r.Use(RecoveryMiddleware())
	// r.Use(CORSMiddleware())
	// r.Use(RateLimitMiddleware())

	protected := r.Group("/api")

	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Perfil privado",
			"userID":  c.GetString("userID"),
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Go Avanzado !",
		})
	})

	r.GET("/users", middleware.HeadersMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Esta es la ruta para ver los usuarios",
		})
	})

	r.GET("/error", func(c *gin.Context) {

		c.Error(errors.New("error interno"))
	})

	r.Run(":8080")
}

func main() {
	fmt.Println("Hello World, GO avanzado !")

	Init()

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// println("Conectado correctamente a la base de datos")

	//////////////////////////////////////////////////////////////////

	// errT := db.Transaction(
	// 	func(tx *gorm.DB) error {

	// 		user := User{
	// 			Name:  "Luisa",
	// 			Email: "luisa@gmail.com",
	// 		}

	// 		if err := tx.Create(&user).Error; err != nil {
	// 			return err
	// 		}

	// 		post := Post{
	// 			Title:  "Primer Post de Luisa",
	// 			UserID: user.ID,
	// 		}

	// 		if err := tx.Create(&post).Error; err != nil {
	// 			return err
	// 		}

	// 		return nil
	// 	},
	// )

	// if errT != nil {
	// 	fmt.Println("Rollback")
	// }

	//////////////////////////////////////////////////////////////////

	// Migración de la tabla 'users' //
	// err = db.AutoMigrate(&User{}, &Post{}, &Role{})
	// if err != nil {
	// 	panic("failed to migrate database")
	// }
	// println("Migración de las tablas 'users' y 'posts' completada correctamente")

	////////////////////////////////////////
	// Insertar un nuevo usuario CREATE//
	// user := User{
	// 	Name:  "Juan",
	// 	Email: "juan@gmail.com",
	// }
	// result := db.Create(&user)
	// if result.Error != nil {
	// 	panic("failed to insert user")
	// }
	// fmt.Printf("Usuario insertado con ID: %d\n", user.ID)

	////////////////////////////////////////
	// Insertar un nuevo usuario con posts y roles CREATE//
	// admin := Role{Name: "Admin"}
	// editor := Role{Name: "Editor"}

	// user := User{
	// 	Name:  "Juan",
	// 	Email: "juan@gmail.com",
	// 	Posts: []Post{
	// 		{Title: "Post 1"},
	// 		{Title: "Post 2"},
	// 		{Title: "Post 3"},
	// 	},
	// 	Roles: []Role{admin, editor},
	// }
	// db.Create(&user)
	////////////////////////////////////////

	// var userWithRoles User
	// db.Preload("Roles").
	// 	First(&userWithRoles, 3)

	// fmt.Printf("Usuario: %s, Roles: ", userWithRoles.Name)
	// for _, role := range userWithRoles.Roles {
	// 	fmt.Printf("%s ", role.Name)
	// }
	////////////////////////////////////////

	// Insertar un nuevo usuario con posts CREATE//
	// user := User{
	// 	Name:  "cristhian",
	// 	Email: "cristhian@gmail.com",
	// 	Posts: []Post{
	// 		{Title: "Post 1"},
	// 		{Title: "Post 2"},
	// 	},
	// }
	// db.Create(&user)

	// var user User
	// db.Preload("Posts").
	// 	First(&user, 2)
	// fmt.Println(user.Posts)
	////////////////////////////////////////

	// agregar roles a usuario existente CREATE//
	// userWithRoles := user_roles{
	// 	UserID: 3,
	// 	RoleID: 2,
	// }
	// db.Create(&userWithRoles)

	////////////////////////CONSULTAS PREPARADAS////////////////////
	// var users []User
	// db.Scopes(Adults).Find(&users)
	// fmt.Println("Usuarios adultos encontrados:")
	// for _, u := range users {
	// 	fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", u.ID, u.Name, u.Email, u.Age)
	// }

	///////////////////////////////////////////////////////

	//// LEER DATOS READ //
	// SELECT * FROM users
	// WHERE id = 1
	// LIMIT 1;
	// db.First(&user, 3)
	// fmt.Printf("Usuario encontrado: ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)

	// db.Where(
	// 	"email = ?",
	// 	"juan@gmail.com",
	// ).First(&user)
	// fmt.Printf("Usuario encontrado: ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)

	// var users []User
	// db.Find(&users)
	// fmt.Println("Usuarios encontrados:")
	// for _, u := range users {
	// 	fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	// }

	///////////////////////////////////////////////////////////
	// db.Limit(10).
	// 	Find(&users)
	// fmt.Println("Usuarios encontrados:")
	// for _, u := range users {
	// 	fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	// }

	// db.Offset(1).
	// 	Limit(10).
	// 	Find(&users)
	// fmt.Println("Usuarios encontrados:")
	// for _, u := range users {
	// 	fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	// }

	// db.Order("id desc").
	// 	Find(&users)
	// fmt.Println("Usuarios encontrados:")
	// for _, u := range users {
	// 	fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	// }

	// db.Select("name,email").
	// 	Find(&users)
	// fmt.Println("Usuarios encontrados:")
	// for _, u := range users {
	// 	fmt.Printf("Name: %s, Email: %s\n", u.Name, u.Email)
	// }

	// var total int64
	// db.Model(&User{}).
	// 	Count(&total)
	// fmt.Printf("Total de usuarios: %d\n", total)
	///////////////////////////////////////////////////////////

	// UPDATE //
	//actualiza un campo específico
	//db.Model(&user).Where("id = ?", 1).Update("Name", "Gigpz")
	//actualiza varios campos
	//db.Model(&user).Where("id = ?", 1).Updates(User{Name: "Gigpz", Email: "gigpz@example.com", Age: 30})

	//DELETE //
	//db.Where("id = ?", 3).Delete(&User{})

	// db.Delete(
	// 	&User{},
	// 	3,
	// )

	/////////////////////////////////////////////////////////////
	// conn, err := pgx.Connect(
	// 	context.Background(),
	// 	"postgres://postgres:gigpz@localhost:5434/bd_tests",
	// )

	// if err != nil {
	// 	panic(err)
	// }

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
	// tx, err := conn.Begin(context.Background())
	// if err != nil {
	// 	panic(err)
	// }

	// defer tx.Rollback(context.Background())

	// _, err = tx.Exec(
	// 	context.Background(),
	// 	"UPDATE accounts SET balance=balance-100 WHERE id=1",
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = tx.Exec(
	// 	context.Background(),
	// 	"UPDATE accounts SET balance=balance+100 WHERE id=2",
	// )

	// if err != nil {
	// 	panic(err)
	// }

	// tx.Commit(context.Background())
	// fmt.Println("Transacciones completada correctamente")
	///////////////////////////////////////////////////////

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

	//defer conn.Close(context.Background())

	//fmt.Println("Conectado correctamente, tabla creada")
	//fmt.Println("Conectado correctamente, fila insertada")
}
