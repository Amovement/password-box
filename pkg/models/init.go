package models

import (
	"fmt"

	"github.com/Amovement/password-box/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	cfg := config.GetConfig()

	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.Addr,
		cfg.MySQL.DBName,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	migrate()
}

func migrate() {
	err := db.AutoMigrate(
		&User{},
		&Box{},
	)

	if err != nil {
		panic(err)
	}

}
