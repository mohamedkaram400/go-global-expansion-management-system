package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1/auth"
	middlewares "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/middlewares/v1/auth"
	requests "github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1/auth"
	responses "github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1/auth"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1/generic_api_response"
)

type UserAuthHandler struct {
	service *services.UserAuthService
}

func NewUserAuthHandler(service *services.UserAuthService) *UserAuthHandler {
	return &UserAuthHandler{service: service}
}

func (h *UserAuthHandler) Login(c *gin.Context) {
	var req requests.UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, accessToken, refreshToken, err := h.service.Login(c, &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	response := generic_api_response.APIResponse{
		Message: "User Login Successfully",
		Data: responses.LoginUserResponse{
			ID:           user.ID,
			Email:        user.Email,
			AccessToken:  accessToken,
			RefrashToken: refreshToken,
		},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *UserAuthHandler) Logout(c *gin.Context) {
	// 1️⃣ Try to get userID from Gin Context
	fmt.Println(c)

	userIDVal, exists := c.Get(string(middlewares.UserIDKey))
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// 2️⃣ Check if userID is a valid string
	userID, ok := userIDVal.(string)

	if !ok || userID == "0" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID"})
		return
	}

	// 3️⃣ Call service.Logout to remove refresh token from Redis
	if err := h.service.Logout(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 4️⃣ Return success response
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
