package initializers

import (
	"jwt-auth-tejdeep/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
