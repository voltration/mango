package db

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error

	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Could not start database: %e", err))
	}

	DB.AutoMigrate(&User{})
}
