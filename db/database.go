package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3306)//database

const url = "root:Interbank2025%@tcp(localhost:3306)/goweb_db"

// guarda la conexioni
var db *sql.DB

// realiza la conexión
func Connect() {

	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexión exitosa")
	db = conection

}

// Cerrar la conexión
func Close() {
	db.Close()
}

// Verificar la conexión
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//verifica si una tabla existe o no

func ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)

	rows, err := Query(sql)

	if err != nil {

		fmt.Println("Error: ", err)
	}
	return rows.Next()

}

// CREA UNA TABLA
func CreateTable(schema string, name string) {

	if !ExistTable(name) {
		_, err := Exec(schema)

		if err != nil {
			fmt.Println("Error: ", err)

		}
		fmt.Println("Se creó tabla")
	} else {
		fmt.Println("Tabla ya estaba creada")

	}

}

// Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

//reiniciar el registro de una tabla

func TruncateTable (tableName string){
	sql := fmt.Sprintf("TRUNCATE %s", tableName)

	Exec((sql))
}
// Polimorfismo de Query

func Query(query string, args ...any) (*sql.Rows, error) {
	row, err := db.Query(query, args...)

	if err != nil {
		fmt.Println(err)
	}

	return row, err

}
