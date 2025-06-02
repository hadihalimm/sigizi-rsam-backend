package repo

import (
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
)

type RoomRepo interface {
	Create(room *model.Room) (*model.Room, error)
	FindAll() ([]model.Room, error)
	FindByID(id uint) (*model.Room, error)
	FindByCode(code string) (*model.Room, error)
	Update(room *model.Room) (*model.Room, error)
	Delete(id uint) error
	FilterByRoomType(roomTypeID uint) ([]model.Room, error)
}

type roomRepo struct {
	db *config.Database
}

func NewRoomRepo(db *config.Database) RoomRepo {
	return &roomRepo{db: db}
}

func (r *roomRepo) Create(room *model.Room) (*model.Room, error) {
	tx := r.db.Gorm.Create(&room)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = r.db.Gorm.Preload("RoomType").First(&room, room.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return room, nil
}

func (r *roomRepo) FindAll() ([]model.Room, error) {
	var rooms []model.Room
	tx := r.db.Gorm.Preload("RoomType").Find(&rooms)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rooms, nil
}

func (r *roomRepo) FindByID(id uint) (*model.Room, error) {
	var room model.Room
	tx := r.db.Gorm.Preload("RoomType").First(&room, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &room, nil
}

func (r *roomRepo) FindByCode(code string) (*model.Room, error) {
	var room model.Room
	tx := r.db.Gorm.Preload("RoomType").Where("code = ?", code).First(&room)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &room, nil
}

func (r *roomRepo) Update(room *model.Room) (*model.Room, error) {
	tx := r.db.Gorm.Save(room)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return room, nil
}

func (r *roomRepo) Delete(id uint) error {
	tx := r.db.Gorm.Delete(&model.Room{}, id)
	return tx.Error
}

func (r *roomRepo) FilterByRoomType(roomTypeID uint) ([]model.Room, error) {
	var rooms []model.Room
	tx := r.db.Gorm.Preload("RoomType").Where("room_type_id = ?", roomTypeID).Find(&rooms)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rooms, nil
}
