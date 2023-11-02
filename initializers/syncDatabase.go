package initializers

import (
	"github.com/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
