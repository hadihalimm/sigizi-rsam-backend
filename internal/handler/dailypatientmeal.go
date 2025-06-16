package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/config"
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
	logger := config.WithRequestContext(config.Logger, c)

	var request request.CreateDailyPatientMeal
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		logger.Warnw("Bad request", "err", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	meal, err := h.dailyPatientMealService.Create(request)
	if err != nil {
		logger.Errorw("Failed to create daily patient meal", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("Daily patient meal created successfully", "objectID", meal.ID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Daily patient meal created successfully",
		"data":    meal,
	})
}

func (h *DailyPatientMealHandler) GetAll(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)

	meals, err := h.dailyPatientMealService.GetAll()
	if err != nil {
		logger.Errorw("Failed to retrieve all daily patient meal", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved all daily patient meals",
		"data":    meals,
	})
}

func (h *DailyPatientMealHandler) GetByID(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)

	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	meal, err := h.dailyPatientMealService.GetByID(id)
	if err != nil {
		logger.Errorw("Failed to retrieve daily patient meal", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("Daily patient meal retrieved successfully", "objectID", meal.ID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Daily patient meal retrieved successfully",
		"data":    meal,
	})
}

func (h *DailyPatientMealHandler) Update(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)

	var request request.UpdateDailyPatientMeal
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		logger.Warnw("Bad request", "err", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	meal, err := h.dailyPatientMealService.Update(id, request)
	if err != nil {
		logger.Errorw("Failed to updated daily patient meal", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("Daily patient meal updated successfully", "objectID", meal.ID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Daily patient meal retrieved successfully",
		"data":    meal,
	})
}

func (h *DailyPatientMealHandler) Delete(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)

	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	err := h.dailyPatientMealService.Delete(id)
	if err != nil {
		logger.Errorw("Failed to delete daily patient meal", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("Daily patient meal deleted successfully", "objectID", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Daily patient meal deleted successfully",
	})
}

func (h *DailyPatientMealHandler) FilterByDateAndRoomType(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)

	dateString := c.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		logger.Warnw("Bad request", "err", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomTypeUint64, _ := strconv.ParseUint(c.Query("roomType"), 10, 64)
	roomType := uint(roomTypeUint64)

	meals, err := h.dailyPatientMealService.FilterByDateAndRoomType(date, roomType)
	if err != nil {
		logger.Errorw("Failed to filter daily patient meal", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("Daily patient meals filtered successfully", "objectCount", len(meals))
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Daily patient meals for RoomType %d %s retrieved successfully", roomType, dateString),
		"data":    meals,
	})
}

func (h *DailyPatientMealHandler) CountByDateAndRoomType(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)

	dateString := c.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		logger.Warnw("Bad request", "err", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomTypeUint64, _ := strconv.ParseUint(c.Query("roomType"), 10, 64)
	roomType := uint(roomTypeUint64)

	matrix, err := h.dailyPatientMealService.CountByDateAndRoomType(date, roomType)
	if err != nil {
		logger.Errorw("Failed to count daily patient meal", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if matrix == nil {
		matrix = []repo.MealMatrixEntry{}
	}

	logger.Infow("Daily patient meals counted successfully", "objectCount", len(matrix))
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Count Daily patient meals for RoomType %d %s retrieved successfully", roomType, dateString),
		"data":    matrix,
	})
}

func (h *DailyPatientMealHandler) ExportToExcel(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)

	dateString := c.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := h.dailyPatientMealService.ExportToExcel(date)
	if err != nil {
		logger.Errorw("Failed to export daily patient meal to spreadsheet", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to export Excel"})
		return
	}

	filename := fmt.Sprintf("permintaan_makanan_%s.xlsx", date.Format("02_01_2006"))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	c.Header("Content-Transfer-Encoding", "binary")

	if err := file.Write(c.Writer); err != nil {
		logger.Errorw("Failed to write file to response", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file to response"})
	}
}

func (h *DailyPatientMealHandler) FilterLogsByDate(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)

	dateString := c.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		logger.Warnw("Bad request", "err", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logs, err := h.dailyPatientMealService.FilterLogsByDate(date)
	if err != nil {
		logger.Errorw("Failed to filter daily patient meal logs", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Infow("Daily patient meals counted successfully", "objectCount", len(logs))
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Daily patient meals log for %s retrieved successfully", dateString),
		"data":    logs,
	})
}

func (h *DailyPatientMealHandler) CountDietCombinationsByDate(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)

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
		logger.Errorw("Failed to count diet combinations", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Infow("Diet combinations count retrieved successfully", "objectCount", len(dietCombinationsCount))
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Diet combinations count for %s retrieved successfully", dateString),
		"data": gin.H{
			"combinationsCount":    dietCombinationsCount,
			"complicationCount":    complicationCount,
			"nonComplicationCount": nonComplicationCount,
		},
	})
}

func (h *DailyPatientMealHandler) CopyFromYesterday(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomTypeUint64, _ := strconv.ParseUint(c.Query("roomType"), 10, 64)
	roomTypeID := uint(roomTypeUint64)

	err = h.dailyPatientMealService.CopyFromYesterday(date, roomTypeID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully copied meals for RoomType %d %s", roomTypeID, dateString),
		"data":    "",
	})
}
