package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/service"
)

type FoodHandler struct {
	foodService service.FoodService
}

func NewFoodHandler(foodService service.FoodService) *FoodHandler {
	return &FoodHandler{foodService: foodService}
}

func (h *FoodHandler) Create(c *gin.Context) {
	var request request.CreateFood
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	food, err := h.foodService.Create(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Food created successfully",
		"data":    food,
	})
}

func (h *FoodHandler) GetAll(c *gin.Context) {
	foods, err := h.foodService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved all room types",
		"data":    foods,
	})
}

func (h *FoodHandler) GetByID(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	food, err := h.foodService.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food retrieved successfully",
		"data":    food,
	})
}

func (h *FoodHandler) Update(c *gin.Context) {
	var request request.UpdateFood
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	food, err := h.foodService.Update(id, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food updated successfully",
		"data":    food,
	})
}

func (h *FoodHandler) Delete(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	err := h.foodService.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food deleted successfully",
	})
}
