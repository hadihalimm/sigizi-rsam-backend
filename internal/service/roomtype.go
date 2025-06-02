package service

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
	"gorm.io/gorm"
)

type RoomTypeService interface {
	Create(request request.CreateRoomType) (*model.RoomType, error)
	GetAll() ([]model.RoomType, error)
	GetByID(id uint) (*model.RoomType, error)
	Update(id uint, request request.UpdateRoomType) (*model.RoomType, error)
	Delete(id uint) error
}

type roomTypeService struct {
	roomTypeRepo repo.RoomTypeRepo
	validate     *validator.Validate
}

func NewRoomTypeService(roomTypeRepo repo.RoomTypeRepo, validate *validator.Validate) RoomTypeService {
	return &roomTypeService{roomTypeRepo: roomTypeRepo, validate: validate}
}

func (s *roomTypeService) Create(request request.CreateRoomType) (*model.RoomType, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newRoomType := &model.RoomType{Name: request.Name, Code: request.Code}
	return s.roomTypeRepo.Create(newRoomType)
}

func (s *roomTypeService) GetAll() ([]model.RoomType, error) {
	return s.roomTypeRepo.FindAll()
}

func (s *roomTypeService) GetByID(id uint) (*model.RoomType, error) {
	return s.roomTypeRepo.FindByID(id)
}

func (s *roomTypeService) Update(id uint, request request.UpdateRoomType) (*model.RoomType, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	rt, err := s.roomTypeRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("room type not found")
	}

	rt.Name = request.Name
	rt.Code = request.Code
	return s.roomTypeRepo.Update(rt)
}

func (s *roomTypeService) Delete(id uint) error {
	return s.roomTypeRepo.Delete(id)
}

func (s *roomTypeService) SyncFromSIMRS() error {
	response, err := http.Get("/tes")
	if err != nil {
		return err
	}
	defer response.Body.Close()
	var result []model.SIMRSRoomType
	err = json.NewDecoder(response.Body).Decode(&result)

	for _, rt := range result {
		existingRoomType, err := s.roomTypeRepo.FindByCode(rt.Code)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			request := request.CreateRoomType{Code: rt.Code, Name: rt.Name}
			_, err = s.Create(request)
		} else {
			request := request.UpdateRoomType{Code: rt.Code, Name: rt.Name}
			_, err = s.Update(existingRoomType.ID, request)
		}
	}
	return err
}
