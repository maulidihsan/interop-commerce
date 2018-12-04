package repository

import models "github.com/maulidihsan/interop-commerce/pkg/models"

type UserService interface {
	AddUser(newUser *models.User) error
	GetByEmail(email string) (*models.User, error)
	Login(cred *models.Credential) (*models.User, error)
}
