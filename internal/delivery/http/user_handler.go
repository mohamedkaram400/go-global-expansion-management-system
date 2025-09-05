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

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service} 
}

func (h *UserHandler) Index(c *gin.Context) {

    skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    users, err := h.Service.GetAllUsers(c.Request.Context(), skip, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "Users Returned successfully",
		Data:    responses.FormatUsers(users),
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Show(c *gin.Context) {
    UserID := c.Param("id")


	user, err := h.Service.FindUserByID(c, UserID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "User Returned Successfully",
		Data:    responses.FormatUser(user),
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Create(c *gin.Context) {
    var req requests.UserRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.Service.InsertUser(c, &req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "User Created Successfully",
		Data:    responses.FormatUser(user),
	}

	c.JSON(http.StatusCreated, response)
} 

func (h *UserHandler) Update(c *gin.Context) {
    UserID := c.Param("id")
	
    var req requests.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	updates := &entities.User{
		Name:                req.Name,
		Email:               req.Email,
		Role:                req.Role,
    }

	user, err := h.Service.UpdateUserByID(c.Request.Context(), UserID, updates)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "User Update Successfully",
		Data:    responses.FormatUser(user),
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Destroy(c *gin.Context) {
    UserID := c.Param("id")

    _, err := h.Service.DeleteUserByID(c.Request.Context(), UserID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusNoContent, nil)
}
