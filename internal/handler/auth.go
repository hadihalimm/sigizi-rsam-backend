package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/sigizi-rsam/internal/api/request"
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)

	var request request.Register
	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		logger.Warnw("Bad request", "err", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.authService.Register(request)
	if err != nil {
		logger.Errorw("Failed to register user", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Infow("User registered successfully", "userID", user.ID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)
	var request request.SignIn
	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		logger.Warnw("Bad request", "err", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.authService.SignIn(request)
	if err != nil {
		logger.Errorw("Failed to sign in", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	session, err := config.SessionStore.Get(c.Request, "sigizi-rsam")
	if err != nil {
		logger.Errorw("Failed to retrieve session", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	session.Values["userID"] = user.ID
	session.Values["username"] = user.Username
	session.Values["name"] = user.Name
	session.Values["role"] = user.Role
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		logger.Errorw("Failed to create session", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Infow("User signed-in successfully", "userID", user.ID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Sign in successful",
		"data": gin.H{
			"userID":   session.Values["userID"],
			"username": session.Values["username"],
			"name":     session.Values["name"],
			"role":     session.Values["role"],
		},
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	logger := config.WithRequestContext(config.Logger, c)
	session, err := config.SessionStore.Get(c.Request, "sigizi-rsam")
	if err != nil {
		logger.Errorw("Failed to retrieve session", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	session.Options.MaxAge = -1
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		logger.Errorw("Failed to save session", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Infow("User signed-out successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

func (h *AuthHandler) CheckSession(c *gin.Context) {
	session, err := config.SessionStore.Get(c.Request, "sigizi-rsam")
	if err != nil || session.Values["userID"] == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Session invalid or expired",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully check the session",
		"data": gin.H{
			"userID":   session.Values["userID"],
			"username": session.Values["username"],
			"name":     session.Values["name"],
			"role":     session.Values["role"],
		},
	})
}
