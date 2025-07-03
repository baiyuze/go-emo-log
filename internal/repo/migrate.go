package repo

import (
	"emoLog/internal/model"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&model.User{},
		&model.Department{},
		&model.Role{},
		&model.Permission{},
		&model.Dict{},
		&model.DictItem{},
		&model.EmotionLog{},
		&model.PlusRecord{},
		&model.Resource{},
		&model.Version{},
		&model.Feedback{},
		&model.Device{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
