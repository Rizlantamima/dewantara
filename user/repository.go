package user

import (
	"qlass-native/models"
)

// Repository represent the user's repository contract
type Repository interface {
	GetByID(id int64) (*models.User, error)
	Store(a *models.User) error	
}