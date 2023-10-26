package models

import (
	"fmt"

	"github.com/cajami/go-mariadb/db"
)

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
`

// Construir usuario
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user

}

// Crear usuario e insertar a la bd
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

// Insertar Registro
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)

	user.Id, _ = result.LastInsertId()
}

// listar todos los registros
func ListUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)

	}
	return users
}

func GetUser(id int) *User {
	user := NewUser("", "", "")

	sql := "SELECT id, username, password, email FROM users WHERE id=?"

	rows, err := db.Query(sql, id)

	if err != nil{
		fmt.Println("Se produjo un error al buscar un usuario por su id: ",err.Error())
	}

	for rows.Next() {
		fmt.Println("entró por aqui")
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}
	return user
}

func (user *User) update() {

	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"

	_, err := db.Exec(sql, user.Username, user.Password, user.Email, user.Id)

	if err != nil {
		fmt.Println("Error al modificar registro:", err.Error())
	} else {
		fmt.Println("Registro fue modificado con éxito")
	}

}

func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}

}

func (user *User) Eliminar(){
	sql := "DELETE FROM users WHERE id=?"

	_, err := db.Exec(sql, user.Id)

	if err != nil {
		fmt.Println("Error al eliminar registro:", err.Error())
	} else {
		fmt.Println("Registro fue eliminado con éxito")
	}
}
