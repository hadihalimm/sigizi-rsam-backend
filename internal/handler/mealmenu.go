package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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

func (h *MealMenuHandler) CreateMealMenuTemplate(c *gin.Context) {
	var request request.CreateMealMenuTemplate
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.mealMenuService.CreateMealMenuTemplate(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Meal menu template created successfully",
		"data":    nil,
	})
}

func (h *MealMenuHandler) GetAllMealMenuTemplate(c *gin.Context) {
	menus, err := h.mealMenuService.FindAllMealMenuTemplate()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All Meal menu templates retrieved successfully",
		"data":    menus,
	})
}

func (h *MealMenuHandler) GetByIDMealMenuTemplate(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("template-id"), 10, 64)
	id := uint(idUint64)
	menu, err := h.mealMenuService.FindByIDMealMenuTemplate(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Meal menu template retrieved successfully",
		"data":    menu,
	})
}

func (h *MealMenuHandler) UpdateMealMenuTemplate(c *gin.Context) {
	var request request.UpdateMealMenuTemplate
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idUint64, _ := strconv.ParseUint(c.Param("template-id"), 10, 64)
	id := uint(idUint64)

	menu, err := h.mealMenuService.UpdateMealMenuTemplate(id, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Meal menu template updated successfully",
		"data":    menu,
	})
}

func (h *MealMenuHandler) DeleteMealMenuTemplate(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("template-id"), 10, 64)
	id := uint(idUint64)

	err := h.mealMenuService.DeleteMealMenuTemplate(id)
	fmt.Println(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Meal menu template deleted successfully",
	})
}

func (h *MealMenuHandler) CreateMenuTemplateSchedule(c *gin.Context) {
	var request request.CreateMenuTemplateSchedule
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schedule, err := h.mealMenuService.CreateMenuTemplateSchedule(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Menu template schedule created successfully",
		"data":    schedule,
	})
}

func (h *MealMenuHandler) GetMenuTemplateScheduleByID(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("schedule-id"), 10, 64)
	id := uint(idUint64)

	schedule, err := h.mealMenuService.FindMenuTemplateScheduleByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Menu template schedule retrieved successfully",
		"data":    schedule,
	})
}

func (h *MealMenuHandler) FilterMenuTemplateScheduleByDate(c *gin.Context) {
	dateString := c.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	schedules, err := h.mealMenuService.FilterMenuTemplateScheduleByDate(date)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Menu template schedule filtered successfully",
		"data":    schedules,
	})
}

func (h *MealMenuHandler) UpdateMenuTemplateSchedule(c *gin.Context) {
	var request request.UpdateMenuTemplateSchedule
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idUint64, _ := strconv.ParseUint(c.Param("schedule-id"), 10, 64)
	id := uint(idUint64)

	schedule, err := h.mealMenuService.UpdateMenuTemplateSchedule(id, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Menu template schedule updated successfully",
		"data":    schedule,
	})
}
