package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"

	"tt/model"
)

var Conn *gorm.DB

func init() {
	dbSqlInmem, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Conn = dbSqlInmem

	err = Conn.Migrator().CreateTable(&model.User{})

	// initial fill for testing
	for i := 0; i < 50; i++ {
		phone := fmt.Sprintf("+1%06d", i)
		age := 35 - i
		user := model.User{
			FullName: fmt.Sprintf("Full Name%d", i),
			Email:    fmt.Sprintf("email%d@company.com", i),
			Phone:    &phone,
			Age:      &age,
		}
		if err = Conn.Create(&user).Error; err != nil {
			log.Fatal(err)
		}
	}
}
