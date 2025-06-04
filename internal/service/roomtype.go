package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

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
	GetSIMRSToken() (*string, error)
	SyncFromSIMRS(token *string) ([]model.RoomType, error)
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

type TokenResponse struct {
	Metadata struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"metadata"`
	Response struct {
		Token string `json:"token"`
	} `json:"response"`
}

func (s *roomTypeService) GetSIMRSToken() (*string, error) {
	req, err := http.NewRequest("GET", "http://192.168.2.20/wssimrs/index.php/auth/token", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Username", os.Getenv("SIMRS_USERNAME"))
	req.Header.Set("X-Password", os.Getenv("SIMRS_PASSWORD"))
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var tokenResponse TokenResponse
	err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	return &tokenResponse.Response.Token, nil
}

func (s *roomTypeService) SyncFromSIMRS(token *string) ([]model.RoomType, error) {

	req, err := http.NewRequest("GET", "http://192.168.2.20/wssimrs/index.php/simrs/referensi/bangsal", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Token", *token)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	type SIMRSBangsalItem struct {
		Kode  string `json:"kode"`
		Ruang string `json:"ruang"`
	}

	type SIMRSBangsalResponse struct {
		Metadata struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"metadata"`
		Response []SIMRSBangsalItem `json:"response"`
	}

	var result SIMRSBangsalResponse
	err = json.NewDecoder(response.Body).Decode(&result)

	for _, rt := range result.Response {
		existingRoomType, err := s.roomTypeRepo.FindByCode(rt.Kode)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			request := request.CreateRoomType{Code: rt.Kode, Name: rt.Ruang}
			_, err = s.Create(request)
		} else {
			request := request.UpdateRoomType{Code: rt.Kode, Name: rt.Ruang}
			_, err = s.Update(existingRoomType.ID, request)
		}
	}
	return s.roomTypeRepo.FindAll()
}
