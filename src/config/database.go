package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("livraria.db"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}
}
