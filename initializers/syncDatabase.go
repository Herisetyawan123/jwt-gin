package initializers

import "github.com/heirsetyawan233/jwt-gin/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
