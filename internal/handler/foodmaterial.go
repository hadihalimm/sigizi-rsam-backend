package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/service"
)

type FoodMaterialHandler struct {
	foodMaterialService service.FoodMaterialService
}

func NewFoodMaterialHandler(foodMaterialService service.FoodMaterialService) *FoodMaterialHandler {
	return &FoodMaterialHandler{foodMaterialService: foodMaterialService}
}
func (h *FoodMaterialHandler) Create(c *gin.Context) {
	var request request.CreateFoodMaterial
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	food, err := h.foodMaterialService.Create(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Food material created successfully",
		"data":    food,
	})
}

func (h *FoodMaterialHandler) GetAll(c *gin.Context) {
	foods, err := h.foodMaterialService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved all food materials",
		"data":    foods,
	})
}

func (h *FoodMaterialHandler) GetByID(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	food, err := h.foodMaterialService.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food material retrieved successfully",
		"data":    food,
	})
}

func (h *FoodMaterialHandler) Update(c *gin.Context) {
	var request request.UpdateFoodMaterial
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	food, err := h.foodMaterialService.Update(id, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food material updated successfully",
		"data":    food,
	})
}

func (h *FoodMaterialHandler) Delete(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	err := h.foodMaterialService.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food material deleted successfully",
	})
}
