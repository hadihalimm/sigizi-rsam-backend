package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
)

type RoomService interface {
	Create(request request.CreateRoom) (*model.Room, error)
	GetAll() ([]model.Room, error)
	GetByID(id uint) (*model.Room, error)
	Update(id uint, request request.UpdateRoom) (*model.Room, error)
	Delete(id uint) error
	FilterByRoomType(roomTypeID uint) ([]model.Room, error)
}

type roomService struct {
	roomRepo     repo.RoomRepo
	roomTypeRepo repo.RoomTypeRepo
	validate     *validator.Validate
}

func NewRoomService(
	roomRepo repo.RoomRepo,
	roomTypeRepo repo.RoomTypeRepo,
	validate *validator.Validate) RoomService {
	return &roomService{
		roomRepo:     roomRepo,
		roomTypeRepo: roomTypeRepo,
		validate:     validate}
}

func (s *roomService) Create(request request.CreateRoom) (*model.Room, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	newRoom := &model.Room{
		RoomNumber:     request.RoomNumber,
		TreatmentClass: request.TreatmentClass,
		RoomTypeID:     request.RoomTypeID,
	}

	return s.roomRepo.Create(newRoom)
}

func (s *roomService) GetAll() ([]model.Room, error) {
	return s.roomRepo.FindAll()
}

func (s *roomService) GetByID(id uint) (*model.Room, error) {
	return s.roomRepo.FindByID(id)
}

func (s *roomService) Update(id uint, request request.UpdateRoom) (*model.Room, error) {
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	room, err := s.roomRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	room.RoomNumber = request.RoomNumber
	room.TreatmentClass = request.TreatmentClass
	room.RoomTypeID = request.RoomTypeID

	return s.roomRepo.Update(room)
}

func (s *roomService) Delete(id uint) error {
	return s.roomRepo.Delete(id)
}

func (s *roomService) FilterByRoomType(roomTypeID uint) ([]model.Room, error) {
	_, err := s.roomTypeRepo.FindByID(roomTypeID)
	if err != nil {
		return nil, errors.New("RoomType not found")
	}
	return s.roomRepo.FilterByRoomType(roomTypeID)
}
