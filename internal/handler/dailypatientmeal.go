package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/service"
)

type DailyPatientMealHandler struct {
	dailyPatientMealService service.DailyPatientMealService
}

func NewDailyPatientMealHandler(dailyPatientMealService service.DailyPatientMealService) *DailyPatientMealHandler {
	return &DailyPatientMealHandler{dailyPatientMealService: dailyPatientMealService}
}

func (h *DailyPatientMealHandler) Create(c *gin.Context) {
	var request request.CreateDailyPatientMeal
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	meal, err := h.dailyPatientMealService.Create(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Daily patient meal created successfully",
		"data":    meal,
	})
}

func (h *DailyPatientMealHandler) GetAll(c *gin.Context) {
	meals, err := h.dailyPatientMealService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved all daily patient meals",
		"data":    meals,
	})
}

func (h *DailyPatientMealHandler) GetByID(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	meal, err := h.dailyPatientMealService.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Daily patient meal retrieved successfully",
		"data":    meal,
	})
}

func (h *DailyPatientMealHandler) Update(c *gin.Context) {
	var request request.UpdateDailyPatientMeal
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	meal, err := h.dailyPatientMealService.Update(id, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Daily patient meal successfully",
		"data":    meal,
	})
}

func (h *DailyPatientMealHandler) Delete(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	err := h.dailyPatientMealService.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Daily patient meal deleted successfully",
	})
}
