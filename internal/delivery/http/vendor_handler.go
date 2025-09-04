package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/generic_api_response"
)

type VendorHandler struct {
	Service *services.VendorService
}

func NewVendorHandler(service *services.VendorService) *VendorHandler {
	return &VendorHandler{Service: service} 
}

func (h *VendorHandler) Index(c *gin.Context) {

    skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    vendors, err := h.Service.GetAllVendors(c.Request.Context(), skip, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "Vendors returned successfully",
		Data:    responses.FormatVendors(vendors),
	}

	c.JSON(http.StatusOK, response)
}

func (h *VendorHandler) Show(c *gin.Context) {
    VendorID := c.Param("id")


	newVendor, err := h.Service.FindVendorByID(c, VendorID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "Vendor Returned Successfully",
		Data:    responses.FormatVendor(newVendor),
	}

	c.JSON(http.StatusOK, response)
}

func (h *VendorHandler) Create(c *gin.Context) {
    var req requests.VendorRequest

    fmt.Printf("📥 Handler received request: %+v\n", req) // debug

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newVendor, err := h.Service.InsertVendor(c, &req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
		Message: "Vendor Created Successfully",
		Data:    responses.FormatVendor(newVendor),
	}

	c.JSON(http.StatusCreated, response)
} 

func (h *VendorHandler) Update(c *gin.Context) {
    VendorID := c.Param("id")
	
    var req requests.VendorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// Marshal []string → datatypes.JSON
    countriesJSON, _ := json.Marshal(req.CountriesSupported)
    servicesJSON, _ := json.Marshal(req.ServicesOffered)


	updates := &entities.Vendor{
        Name:               req.Name,
        CountriesSupported: countriesJSON,
        ServicesOffered:    servicesJSON,
        Rating:             req.Rating,
        ResponseSlaHours:   req.ResponseSlaHours,
    }

	vendor, err := h.Service.UpdateVendorByID(c.Request.Context(), VendorID, updates)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	response := generic_api_response.APIResponse{
        Message: "Vendor updated successfully",
		Data:    responses.FormatVendor(vendor),
    }

	c.JSON(http.StatusOK, response)
}

func (h *VendorHandler) Destroy(c *gin.Context) {
    VendorID := c.Param("id")

    _, err := h.Service.DeleteVendorByID(c.Request.Context(), VendorID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusNoContent, nil)
}
