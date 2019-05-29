package usecase

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"qlass-native/user"
	"qlass-native/models"
)


type userUsecase struct {
	userRepo    user.Repository
	contextTimeout time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of user.Usecase interface
func NewUserUsecase(repo user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: repo,
	}
}