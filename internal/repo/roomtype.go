package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type RoomTypeRepo interface {
	Create(roomType *model.RoomType) (*model.RoomType, error)
	FindAll() ([]model.RoomType, error)
	FindByID(id uint) (*model.RoomType, error)
	FindByCode(code string) (*model.RoomType, error)
	Update(rt *model.RoomType) (*model.RoomType, error)
	Delete(id uint) error
}

type roomTypeRepo struct {
	db *config.Database
}

func NewRoomTypeRepo(db *config.Database) RoomTypeRepo {
	return &roomTypeRepo{db: db}
}

func (r *roomTypeRepo) Create(roomType *model.RoomType) (*model.RoomType, error) {
	tx := r.db.Gorm.Create(roomType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return roomType, nil
}

func (r *roomTypeRepo) FindAll() ([]model.RoomType, error) {
	var roomTypes []model.RoomType
	tx := r.db.Gorm.Find(&roomTypes)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return roomTypes, nil
}

func (r *roomTypeRepo) FindByID(id uint) (*model.RoomType, error) {
	var rt model.RoomType
	tx := r.db.Gorm.First(&rt, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &rt, nil
}

func (r *roomTypeRepo) FindByCode(code string) (*model.RoomType, error) {
	var rt model.RoomType
	tx := r.db.Gorm.Where("code = ?", code).First(&rt)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &rt, nil
}

func (r *roomTypeRepo) Update(rt *model.RoomType) (*model.RoomType, error) {
	tx := r.db.Gorm.Save(rt)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rt, nil
}

func (r *roomTypeRepo) Delete(id uint) error {
	tx := r.db.Gorm.Delete(&model.RoomType{}, id)
	return tx.Error
}
