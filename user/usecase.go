package user

import (
	"qlass-native/models"
)

// Usecase represent the users's usecases
type Usecase interface {
	// GetByID(id int64) (*models.User, error)
	Store(*models.User) error
}