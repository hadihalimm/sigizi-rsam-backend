package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.userService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved all users",
		"data":    users,
	})
}

func (h *UserHandler) GetByID(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	user, err := h.userService.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User retrieved successfully",
		"data":    user,
	})
}

func (h *UserHandler) Update(c *gin.Context) {
	var request request.UpdateUser
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	user, err := h.userService.Update(id, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"data":    user,
	})
}

func (h *UserHandler) Delete(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	err := h.userService.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	user, err := h.userService.ResetPassword(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully reset user password",
		"data":    user,
	})
}

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	var request request.UpdatePassword
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idUint64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	id := uint(idUint64)
	user, err := h.userService.UpdatePassword(id, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully reset user password",
		"data":    user,
	})
}
