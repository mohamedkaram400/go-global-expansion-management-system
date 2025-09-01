package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/generic_api_response"
)

type ClientHandler struct {
	Service *services.ClientService
}

func NewClientHandler(service *services.ClientService) *ClientHandler {
	return &ClientHandler{Service: service} 
}

func (h *ClientHandler) Index(c *gin.Context) {

    skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    clients, err := h.Service.GetAllClients(c.Request.Context(), skip, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	var clientResponses []responses.ClientResponse
    for _, client := range clients {
        clientResponses = append(clientResponses, responses.ClientResponse{
            ID:           client.ID,
            CompanyName:  client.CompanyName,
            ContactEmail: client.ContactEmail,
        })
    }

	response := generic_api_response.APIResponse{
		Message: "Clients Returned Successfully",
		Data: clientResponses,
	}

	c.JSON(http.StatusOK, response)
}

func (h *ClientHandler) Show(c *gin.Context) {
    clientID := c.Param("id")


	newClient, err := h.Service.FindClientByID(c, clientID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "Client Returned Successfully",
		Data: responses.ClientResponse{
			ID:           newClient.ID,
			CompanyName:  newClient.CompanyName,
			ContactEmail: newClient.ContactEmail,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *ClientHandler) Create(c *gin.Context) {
    var req requests.ClientRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newClient, err := h.Service.InsertClient(c, &req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "Client Created Successfully",
		Data: responses.ClientResponse{
			ID:           newClient.ID,
			CompanyName:  newClient.CompanyName,
			ContactEmail: newClient.ContactEmail,
		},
	}

	c.JSON(http.StatusCreated, response)
} 

func (h *ClientHandler) Update(c *gin.Context) {
    clientID := c.Param("id")
	
    var req requests.ClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	updates := &entities.Client{
        CompanyName:  req.CompanyName,
        ContactEmail: req.ContactEmail,
    }

	client, err := h.Service.UpdateClientByID(c.Request.Context(), clientID, updates)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "Client Update Successfully",
		Data: responses.ClientResponse{
			ID:           client.ID,
			CompanyName:  client.CompanyName,
			ContactEmail: client.ContactEmail,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *ClientHandler) Destroy(c *gin.Context) {
    clientID := c.Param("id")

    _, err := h.Service.DeleteClientByID(c.Request.Context(), clientID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusNoContent, nil)
}
