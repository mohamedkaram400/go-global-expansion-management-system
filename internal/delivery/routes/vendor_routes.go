package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/middlewares"
)

func RegisterVendorRoutes(rg *gin.RouterGroup, vendorHandler *http.VendorHandler) {
	vendor := rg.Group("/vendor")
	vendor.Use(middlewares.JWTAuth())
	// vendor.Use(middlewares.JWTAuth())

	{
		vendor.POST("/create-vendor", vendorHandler.Create)
		vendor.GET("/all-vendors", vendorHandler.Index)
        vendor.GET("/show-vendor/:id", vendorHandler.Show)
        vendor.PUT("/update-vendor/:id", vendorHandler.Update)
        vendor.DELETE("/delete-vendor/:id", vendorHandler.Destroy)
	}
}
