package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/service"
)

type MealMenuHandler struct {
	mealMenuService service.MealMenuService
}

func NewMealMenuHandler(mealMenuService service.MealMenuService) *MealMenuHandler {
	return &MealMenuHandler{mealMenuService: mealMenuService}
}

func (h *MealMenuHandler) Create(c *gin.Context) {
	var request request.CreateMealMenu
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menu, err := h.mealMenuService.Create(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Meal menu created successfully",
		"data":    menu,
	})
}

func (h *MealMenuHandler) GetAll(c *gin.Context) {
	menus, err := h.mealMenuService.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All Meal menus retrieved successfully",
		"data":    menus,
	})
}

func (h *MealMenuHandler) GetByID(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	menu, err := h.mealMenuService.FindByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Meal menu retrieved successfully",
		"data":    menu,
	})
}

func (h *MealMenuHandler) Update(c *gin.Context) {
	var request request.UpdateMealMenu
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)

	menu, err := h.mealMenuService.Update(id, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Meal menu updated successfully",
		"data":    menu,
	})
}

func (h *MealMenuHandler) Delete(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)

	err := h.mealMenuService.Delete(id)
	fmt.Println(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Meal menu deleted successfully",
	})
}
