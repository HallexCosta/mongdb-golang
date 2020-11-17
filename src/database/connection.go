package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// IConnection ...
type IConnection interface {
	SetConfig(connectionConfig *ConnectionConfig)
	GetConnection() *gorm.DB
}

// Connection ...
type Connection struct {
	ConnectionConfig *ConnectionConfig
}

// Config ...
type Config = gorm.Config

// ConnectionConfig ...
type ConnectionConfig struct {
	Path   string
	Config Config
}

// NewConnection ...
func NewConnection() IConnection {
	return &Connection{}
}

// SetConfig ...
func (c *Connection) SetConfig(connectionConfig *ConnectionConfig) {
	c.ConnectionConfig = connectionConfig
}

// GetConnection ...
func (c *Connection) GetConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(c.ConnectionConfig.Path), &c.ConnectionConfig.Config)

	if err != nil {
		panic("Error at create database")
	}

	return db
}
