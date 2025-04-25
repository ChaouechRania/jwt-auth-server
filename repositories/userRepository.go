package repositories

import (
	"jwt-auth-server/api"
	"jwt-auth-server/config"
	"jwt-auth-server/encrypt"
	"jwt-auth-server/errors"
	"jwt-auth-server/models"
	token "jwt-auth-server/token"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	Authenticate(login *api.Login) (string, error)
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUsertRepository() UserRepository {
	return &userRepository{db: config.DB}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Authenticate(login *api.Login) (string, error) {
	var err error
	var u models.User
	err = r.db.Model(models.User{}).Where("email = ?", login.Email).Take(&u).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return "", errors.ErrNoSuchEntity
		}
		return "", err
	}
	err = encrypt.VerifyPassword(u.Password, login.Password)
	if err != nil {
		return "", errors.ErrNoSuchPassword
	}
	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (u *userRepository) PrepareGive() {
	var user models.User
	user.Password = ""
}
