package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
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
		"message": "Daily patient meal retrieved successfully",
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

func (h *DailyPatientMealHandler) FilterByDateAndRoomType(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomTypeUint64, _ := strconv.ParseUint(c.Query("roomType"), 10, 64)
	roomType := uint(roomTypeUint64)

	meals, err := h.dailyPatientMealService.FilterByDateAndRoomType(date, roomType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Daily patient meals for RoomType %d %s retrieved successfully", roomType, dateString),
		"data":    meals,
	})
}

func (h *DailyPatientMealHandler) CountByDateAndRoomType(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomTypeUint64, _ := strconv.ParseUint(c.Query("roomType"), 10, 64)
	roomType := uint(roomTypeUint64)

	matrix, err := h.dailyPatientMealService.CountByDateAndRoomType(date, roomType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if matrix == nil {
		matrix = []repo.MealMatrixEntry{}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Count Daily patient meals for RoomType %d %s retrieved successfully", roomType, dateString),
		"data":    matrix,
	})
}

func (h *DailyPatientMealHandler) ExportToExcel(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := h.dailyPatientMealService.ExportToExcel(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to export Excel"})
		return
	}

	filename := fmt.Sprintf("permintaan_makanan_%s.xlsx", date.Format("02_01_2006"))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	c.Header("Content-Transfer-Encoding", "binary")

	if err := file.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file to response"})
	}
}

func (h *DailyPatientMealHandler) FilterLogsByDate(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomTypeUint64, _ := strconv.ParseUint(c.Query("roomType"), 10, 64)
	roomType := uint(roomTypeUint64)

	logs, err := h.dailyPatientMealService.FilterLogsByDateAndRoomType(date, roomType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Daily patient meals log for %s retrieved successfully", dateString),
		"data":    logs,
	})
}

func (h *DailyPatientMealHandler) CountDietCombinationsByDate(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dietCombinationsCount,
		complicationCount,
		nonComplicationCount,
		err := h.dailyPatientMealService.CountDietCombinationsByDate(date)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Diet combinations count for %s retrieved successfully", dateString),
		"data": gin.H{
			"combinationsCount":    dietCombinationsCount,
			"complicationCount":    complicationCount,
			"nonComplicationCount": nonComplicationCount,
		},
	})
}
