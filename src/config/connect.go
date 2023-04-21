package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(config *Config) (*sql.DB, error) {

	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName))
	fmt.Println(config.DBHost, config.DBUserName, config.DBPort, config.DBName)
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	// print if success to connect
	// defer the close till after the main function has finished
	// executing
	// defer DB.Close()
	fmt.Println("? Connected Successfully to the Database")
	return DB, err
}
