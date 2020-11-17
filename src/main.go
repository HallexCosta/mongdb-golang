package main

import (
	"fmt"

	"github.com/HallexCosta/sqlite-golang/src/configs"
	"github.com/HallexCosta/sqlite-golang/src/database"
	"github.com/HallexCosta/sqlite-golang/src/database/migrations"
	"github.com/HallexCosta/sqlite-golang/src/entities"
)

func init() {
	migrate := migrations.NewMigrate()
	migrate.Run()
}

func main() {
	var connection database.IConnection = database.NewConnection()
	connection.SetConfig(configs.GetConnectionConfig())

	db := connection.GetConnection()

	var user *entities.User = &entities.User{
		Name:     "Hállex",
		Email:    "hallex.costa@hotmail.com",
		Password: "hallex123",
	}

	//db.Create(user)

	db.Select("*").Create(&user)

	fmt.Println("Database connected with successfully!!")
}
