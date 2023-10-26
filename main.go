package main

import (
	"fmt"

	"github.com/cajami/go-mariadb/db"
	"github.com/cajami/go-mariadb/models"
)

func main() {
	db.Connect()

	db.TruncateTable("users")

	models.CreateUser("fabrizio", "fabrizio.javier4423", "fabrizio.javier1618@gmail.com")
	models.CreateUser("Javier", "javier4423", "javier2315@gmail.com")

	db.Close()
	return


	// models.CreateUser("fabrizio", "fabrizio.javier4423", "fabrizio.javier1618@gmail.com")

	// fmt.Println(user)
	// users := models.ListUsers()
	// fmt.Println(users)

	user :=models.GetUser(2)

	fmt.Println(user)

	// user.Username = "Fabrizio Javier Huiñocana Sama"
	// user.Save()

	// user.Eliminar()

	// db.TruncateTable("users")


	// db .TruncateTable("users")
	//db.Ping()
	db.Close()
}
