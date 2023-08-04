package migrations

import (
	"api_golang/models"

	"gorm.io/gorm"
)

func RodaMigration(db *gorm.DB) {
	db.AutoMigrate(models.Livro{})
}
