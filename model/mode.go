package model

import (
	"log"

	"gorm.io/gorm"
)

func InitTables(db *gorm.DB, model interface{}) {
	rh := model
	if !db.Migrator().HasTable(rh) {
		_ = db.Migrator().CreateTable(rh)
	}
	err := db.AutoMigrate(rh)
	if err != nil {
		log.Fatal("InitTables error:", err)
	}
}
