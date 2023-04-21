package main

import (
	"Jobhun_Mahasiswa/src/config"
	"Jobhun_Mahasiswa/src/routers"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configEnv, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	config.DB()

	fmt.Println("Connection successful!")

	PORT := fmt.Sprintf(":%s", configEnv.ServerPort)
	routers.StartServer().Run(PORT)
}
