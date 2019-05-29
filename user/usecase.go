package user

import (
	"qlass-native/models"
)

// Usecase represent the users's usecases
type Usecase interface {
	GetByID(id uuid) (*models.User, error)
	Store(*models.User) error
}