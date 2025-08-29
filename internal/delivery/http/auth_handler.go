package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/middlewares"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/generic_api_response"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service} 
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req requests.RegisterRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newClient, err := h.service.Register(c, &req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "Client Registerd Successfully",
		Data: responses.RegisterClientResponse{
			ID:           newClient.ID,
			CompanyName:  newClient.CompanyName,
			ContactEmail: newClient.ContactEmail,
		},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req requests.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, accessToken, refreshToken, err := h.service.Login(c, &req)
	if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	response := generic_api_response.APIResponse{
		Message: "Client Login Successfully",
		Data: responses.LoginClientResponse{
			ID:           client.ID,
			CompanyName:  client.CompanyName,
			ContactEmail: client.ContactEmail,
			AccessToken:  accessToken,
			RefrashToken: refreshToken,
		},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// 1️⃣ Try to get clientID from Gin Context
	clientIDVal, exists := c.Get(string(middlewares.ClientIDKey))
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// 2️⃣ Check if clientID is a valid string
	clientID, ok := clientIDVal.(uint)

	if !ok || clientID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid client ID"})
		return
	}

	// 3️⃣ Call service.Logout to remove refresh token from Redis
	if err := h.service.Logout(clientID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 4️⃣ Return success response
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
