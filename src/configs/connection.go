package configs

import (
	"github.com/HallexCosta/sqlite-golang/src/database"
	"gorm.io/gorm"
)

// GetConnectionConfig ...
func GetConnectionConfig() *database.ConnectionConfig {
	return &database.ConnectionConfig{
		Path:   "./src/database/database.sqlite",
		Config: gorm.Config{},
	}
}
