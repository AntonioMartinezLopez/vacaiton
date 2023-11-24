package repository

import (
	"backend/services/userService/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	// GetUserById(id string) (res interface{}, err error)
	// GetUserByMail(id string) (res interface{}, err error)
	RegisterUser(registerUserInput *models.RegisterUserInput) (res int, err error)
}

func (r *GormRepository) GetUserById(id string) (user *models.User, err error) {
	err = r.db.Database.Find(&user).Error
	return
}

func (r *GormRepository) GetUserByMail(email string) (user *models.User, err error) {
	err = r.db.Database.First(&user, "Email = ?", email).Error
	return
}

func (r *GormRepository) RegisterUser(registerUserInput *models.RegisterUserInput) (userId int, err error) {

	// First check whether user already exists
	if _, error := r.GetUserByMail(registerUserInput.Email); error == nil {
		return 0, errors.New("ERROR: User with given email already exists")
	}

	// Hash password
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(registerUserInput.Password), bcrypt.DefaultCost)

	if error != nil {
		return 0, error
	}

	newUser := models.User{
		Email:     registerUserInput.Email,
		Firstname: registerUserInput.Firstname,
		Lastname:  registerUserInput.Lastname,
		Password:  string(hashedPassword),
	}
	result := r.db.Database.Create(&newUser)

	return newUser.Id, result.Error
}
