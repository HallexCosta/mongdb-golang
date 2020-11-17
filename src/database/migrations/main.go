package migrations

import (
	"github.com/HallexCosta/sqlite-golang/src/configs"
	"github.com/HallexCosta/sqlite-golang/src/database"
	"github.com/HallexCosta/sqlite-golang/src/entities"
)

// IMigrate ..
type IMigrate interface {
	Run()
}

// Migrate ...
type Migrate struct {
}

// NewMigrate ...
func NewMigrate() IMigrate {
	return &Migrate{}
}

// Run ...
func (m *Migrate) Run() {
	var connection database.IConnection = database.NewConnection()
	connection.SetConfig(configs.GetConnectionConfig())

	db := connection.GetConnection()

	err := db.AutoMigrate(entities.User{})

	if err != nil {
		panic("Error at create migrate")
	}
}
