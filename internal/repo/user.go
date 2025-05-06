package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type UserRepo interface {
	Create(user *model.User) (*model.User, error)
	FindAll() ([]model.User, error)
	FindByID(id uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id uint) error
}

type userRepo struct {
	db *config.Database
}

func NewUserRepo(db *config.Database) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *model.User) (*model.User, error) {
	tx := r.db.Gorm.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (r *userRepo) FindAll() ([]model.User, error) {
	var users []model.User
	tx := r.db.Gorm.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (r *userRepo) FindByID(id uint) (*model.User, error) {
	var user model.User
	tx := r.db.Gorm.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (r *userRepo) FindByUsername(username string) (*model.User, error) {
	var user model.User
	tx := r.db.Gorm.Where("username = ?", username).Limit(1).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (r *userRepo) Update(user *model.User) (*model.User, error) {
	tx := r.db.Gorm.Save(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (r *userRepo) Delete(id uint) error {
	tx := r.db.Gorm.Delete(&model.User{}, id)
	return tx.Error
}
