package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/service"
)

type SnackHandler struct {
	snackService service.SnackService
}

func NewSnackHandler(snackService service.SnackService) *SnackHandler {
	return &SnackHandler{snackService: snackService}
}

func (h *SnackHandler) Create(c *gin.Context) {
	var request request.CreateSnack
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	snack, err := h.snackService.Create(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Snack created successfully",
		"data":    snack,
	})
}

func (h *SnackHandler) GetAll(c *gin.Context) {
	snacks, err := h.snackService.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully retrieved all snacks",
		"data":    snacks,
	})
}

func (h *SnackHandler) GetByID(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("snack-id"), 10, 64)
	id := uint(idUint64)
	snack, err := h.snackService.FindByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Snack retrieved successfully",
		"data":    snack,
	})
}

func (h *SnackHandler) Update(c *gin.Context) {
	var request request.UpdateSnack
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idUint64, _ := strconv.ParseUint(c.Param("snack-id"), 10, 64)
	id := uint(idUint64)
	snack, err := h.snackService.Update(id, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Snack updated successfully",
		"data":    snack,
	})
}

func (h *SnackHandler) Delete(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("snack-id"), 10, 64)
	id := uint(idUint64)
	err := h.snackService.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Snack deleted successfully",
	})
}

func (h *SnackHandler) CreateVariant(c *gin.Context) {
	var request request.CreateSnackVariant
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint64, _ := strconv.ParseUint(c.Param("snack-id"), 10, 64)
	snackID := uint(idUint64)

	variant, err := h.snackService.CreateVariant(snackID, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Snack variant created successfully",
		"data":    variant,
	})
}

func (h *SnackHandler) GetAllVariant(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("snack-id"), 10, 64)
	snackID := uint(idUint64)

	variants, err := h.snackService.FindAllVariant(snackID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved all snack variants",
		"data":    variants,
	})
}

func (h *SnackHandler) GetVariantByID(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("variant-id"), 10, 64)
	id := uint(idUint64)

	variant, err := h.snackService.FindVariantByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Snack variant retrieved successfully",
		"data":    variant,
	})
}

func (h *SnackHandler) UpdateVariant(c *gin.Context) {
	var request request.UpdateSnackVariant
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idUint64, _ := strconv.ParseUint(c.Param("variant-id"), 10, 64)
	id := uint(idUint64)

	variant, err := h.snackService.UpdateVariant(id, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Snack updated successfully",
		"data":    variant,
	})
}

func (h *SnackHandler) DeleteVariant(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("variant-id"), 10, 64)
	id := uint(idUint64)
	err := h.snackService.DeleteVariant(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Snack variant deleted successfully",
	})
}
