package usecase

import (
	models "github.com/maulidihsan/flashdeal-webservice/pkg/models"
	repository "github.com/maulidihsan/flashdeal-webservice/pkg/user/repository"
)

type UserUsecase interface {
	AddUser(newUser *models.User) error
	GetByEmail(email string) (*models.User, error)
	Login(cred *models.Credential) (*models.User, error)
}

type userUsecase struct {
	userRepos *repository.UserCollection
}

func (u *userUsecase) AddUser(newUser *models.User) error {
	return u.userRepos.AddUser(newUser)
}

func (u *userUsecase) GetByEmail(email string ) (*models.User, error) {
	return u.userRepos.GetByEmail(email)
}

func (u *userUsecase) Login(cred *models.Credential) (*models.User, error) {
	return u.userRepos.Login(cred)
}

func NewUserUseCase(u *repository.UserCollection) UserUsecase {
	return &userUsecase{u}
}