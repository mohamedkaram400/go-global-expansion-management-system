package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	middlewares "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/middlewares/auth"
	services "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/auth"
	requests "github.com/mohamedkaram400/go-global-expansion-management-system/requests/auth"
	responses "github.com/mohamedkaram400/go-global-expansion-management-system/responses/auth"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/generic_api_response"
)

type ClientAuthHandler struct {
	service *services.ClientAuthService
}

func NewClientAuthHandler(service *services.ClientAuthService) *ClientAuthHandler {
	return &ClientAuthHandler{service: service} 
}

func (h *ClientAuthHandler) Register(c *gin.Context) {
    var req requests.ClientRegisterRequest

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

func (h *ClientAuthHandler) Login(c *gin.Context) {
	var req requests.ClientLoginRequest

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

func (h *ClientAuthHandler) Logout(c *gin.Context) {
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
