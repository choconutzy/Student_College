package config

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
)

func DB() *sql.DB {
	// Load env variables
	configDB, err := LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load env")
	}
	// check tcp connection
	ipAddress := configDB.DBHost
	port := configDB.DBPort
	target := ipAddress + ":" + port
	ln, err := net.Listen("tcp", target)
	if err != nil {
		fmt.Println("server, Listen", err)
	}
	fmt.Println(ln)

	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Println("Connection error:", err)
		return nil
	}
	defer conn.Close()

	// Connect to database
	db, err := ConnectDB(&configDB)

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	// print if success to connect
	fmt.Println("? Connected Successfully to the Database")

	fmt.Println("Connection successful!")
	return db
}
