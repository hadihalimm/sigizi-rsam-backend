package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
	"gorm.io/gorm"
)

type RoomService interface {
	Create(request request.CreateRoom) (*model.Room, error)
	GetAll() ([]model.Room, error)
	GetByID(id uint) (*model.Room, error)
	Update(id uint, request request.UpdateRoom) (*model.Room, error)
	Delete(id uint) error
	FilterByRoomType(roomTypeID uint) ([]model.Room, error)
	SyncFromSIMRS(token *string, roomTypeID uint, roomTypeCode string) error
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
		Code:           request.Code,
		Name:           request.Name,
		ClassID:        request.ClassID,
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

	room.Name = request.Name
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

func (s *roomService) SyncFromSIMRS(token *string, roomTypeID uint, roomTypeCode string) error {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("http://192.168.2.20/wssimrs/index.php/simrs/referensi/kamar/%s", roomTypeCode), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Token", *token)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	type SIMRSKamarItem struct {
		Kode         string `json:"kode"`
		Nama         string `json:"nama_kamar"`
		KelasID      string `json:"kelas_id"`
		KelasLayanan string `json:"kelas_layanan"`
	}

	type SIMRSKamarResponse struct {
		Metadata struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"metadata"`
		Response []SIMRSKamarItem `json:"response"`
	}

	var result SIMRSKamarResponse
	err = json.NewDecoder(response.Body).Decode(&result)

	for _, kamar := range result.Response {
		existingRoom, err := s.roomRepo.FindByCode(kamar.Kode)
		fmt.Println(errors.Is(err, gorm.ErrRecordNotFound))
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			request := request.CreateRoom{
				Code:           kamar.Kode,
				Name:           kamar.Nama,
				ClassID:        kamar.KelasID,
				TreatmentClass: kamar.KelasLayanan,
				RoomTypeID:     roomTypeID}
			_, err = s.Create(request)
		} else {
			request := request.UpdateRoom{
				Code:           kamar.Kode,
				Name:           kamar.Nama,
				ClassID:        kamar.KelasID,
				TreatmentClass: kamar.KelasLayanan,
				RoomTypeID:     roomTypeID}
			_, err = s.Update(existingRoom.ID, request)
		}
	}
	return err
}
