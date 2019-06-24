package usecase

import (
	"qlass-native/user"
	"qlass-native/models"
)


type userUsecase struct {
	userRepo    user.Repository
}

// NewUserUsecase will create new an userUsecase object representation of user.Usecase interface
func NewUserUsecase(repo user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: repo,
	}
}


func (u *userUsecase) Store(c context.Context, m *models.User) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, _ := a.GetByTitle(ctx, m.Title)
	if existedArticle != nil {
		return models.ErrConflict
	}

	err := a.articleRepo.Store(ctx, m)
	if err != nil {
		return err
	}
	return nil
	
}