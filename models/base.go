package models

import (
	"ardidas/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(config.MainConfig.Sql.Database, config.MainConfig.Sql.URI)
	if err != nil {
		log.Fatal("MySQL connection fail")
	}
	db.AutoMigrate(Item{}, Request{}, User{})

}
