package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/service"
)

type AllergyHandler struct {
	allergyService service.AllergyService
}

func NewAllergyHandler(allergyService service.AllergyService) *AllergyHandler {
	return &AllergyHandler{allergyService: allergyService}
}

func (h *AllergyHandler) Create(c *gin.Context) {
	var request request.CreateAllergy
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allergy, err := h.allergyService.Create(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Allergy created successfully",
		"data":    allergy,
	})
}

func (h *AllergyHandler) GetAll(c *gin.Context) {
	allergies, err := h.allergyService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved all allergy",
		"data":    allergies,
	})
}

func (h *AllergyHandler) GetByID(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	allergy, err := h.allergyService.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Allergy retrieved successfully",
		"data":    allergy,
	})
}

func (h *AllergyHandler) Update(c *gin.Context) {
	var request request.UpdateAllergy
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	allergy, err := h.allergyService.Update(id, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Allergy updated successfully",
		"data":    allergy,
	})
}

func (h *AllergyHandler) Delete(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	err := h.allergyService.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Allergy deleted successfully",
	})
}
