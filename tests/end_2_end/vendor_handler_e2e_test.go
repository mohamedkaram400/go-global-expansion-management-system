package end_2_end

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
	repositories "github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories/v1"
	services "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	handlers "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/v1"
	"github.com/stretchr/testify/assert"
)

func setupVendorRouter() *gin.Engine {
	db := conn.SetupTestDB()
	vendorRepo := repositories.NewVendorRepo(db)
	vendorService := services.NewVendorService(vendorRepo)
	vendorHandler := handlers.NewVendorHandler(vendorService)

	r := gin.Default()
	r.GET("/vendors", vendorHandler.Index)
	r.GET("/vendors/:id", vendorHandler.Show)
	r.POST("/vendors", vendorHandler.Create)
	r.PUT("/vendors/:id", vendorHandler.Update)
	r.DELETE("/vendors/:id", vendorHandler.Destroy)

	return r
}

func TestGetAllVendors(t *testing.T) {
	router := setupVendorRouter()

	req, _ := http.NewRequest("GET", "/vendors?skip=0&limit=10", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCreateVendor(t *testing.T) {
	router := setupVendorRouter()

	body := map[string]interface{}{
		"name":                "ITHelper",
		"countries_supported": []string{"KSA", "UAE"},
		"services_offered":    []string{"hiring"},
		"rating":              4.5,
		"response_sla_hours":  3,
	}

	jsonBody, _ := json.Marshal(body)

	// log.Println(jsonBody)

	req, _ := http.NewRequest("POST", "/vendors", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	fmt.Println(resp)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "ITHelper")

}

func TestGetVendorById(t *testing.T) {
	router := setupVendorRouter()

	createBody := map[string]interface{}{
		"name":                "ITWorks",
		"countries_supported": []string{"Egypt", "UAE"},
		"services_offered":    []string{"hr"},
		"rating":              4.5,
		"response_sla_hours":  3,
	}

	jsonCreate, _ := json.Marshal(createBody)

	req, _ := http.NewRequest("POST", "/vendors", bytes.NewBuffer(jsonCreate))
	req.Header.Set("Content-Type", "application/json")

	respCreate := httptest.NewRecorder()
	router.ServeHTTP(respCreate, req)

	t.Log("Create response:", respCreate.Body.String())
	assert.Equal(t, http.StatusCreated, respCreate.Code)

	// Now GET
	req, _ = http.NewRequest("GET", "/vendors/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateVendor(t *testing.T) {
	router := setupVendorRouter()

	// First, create a Vendor to update
	createBody := map[string]interface{}{
		"name":                "ITWorks",
		"countries_supported": []string{"Egypt", "UAE"},
		"services_offered":    []string{"hr"},
		"rating":              4.5,
		"response_sla_hours":  3,
	}
	jsonCreate, _ := json.Marshal(createBody)

	req, _ := http.NewRequest("POST", "/vendors", bytes.NewBuffer(jsonCreate))
	req.Header.Set("Content-Type", "application/json")

	respCreate := httptest.NewRecorder()
	router.ServeHTTP(respCreate, req)

	assert.Equal(t, http.StatusCreated, respCreate.Code)

	// Now extract the ID from response
	var created map[string]interface{}
	_ = json.Unmarshal(respCreate.Body.Bytes(), &created)
	// t.Log("Create response:", respCreate.Body.String())

	data := created["data"].(map[string]interface{})
	id := int(data["id"].(float64))

	// Update request
	updateBody := map[string]interface{}{
		"name":                "ITWorks",
		"countries_supported": []string{"Egypt", "UAE"},
		"services_offered":    []string{"hr"},
		"rating":              2,
		"response_sla_hours":  5,
	}

	jsonUpdate, _ := json.Marshal(updateBody)

	reqUpdate, _ := http.NewRequest("PUT", "/vendors/"+strconv.Itoa(id), bytes.NewBuffer(jsonUpdate))
	reqUpdate.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, reqUpdate)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "ITWorks")
}

func TestDeleteVendor(t *testing.T) {
	router := setupVendorRouter()

	req, _ := http.NewRequest("DELETE", "/vendors/4", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}

func TestVendorHandler_EndToEnd(t *testing.T) {
	router := setupVendorRouter()

	// ---- 1. CREATE ----
	createBody := map[string]interface{}{
		"name":                "ITWorks",
		"countries_supported": []string{"Egypt", "UAE"},
		"services_offered":    []string{"hr"},
		"rating":              4.5,
		"response_sla_hours":  3,
	}

	jsonBody, _ := json.Marshal(createBody)
	req, _ := http.NewRequest("POST", "/vendors", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	t.Log("Create response:", resp.Code, resp.Body.String())

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createResp map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &createResp)
	data := createResp["data"].(map[string]interface{})
	id := int(data["id"].(float64))

	// ---- 2. INDEX ----
	req, _ = http.NewRequest("GET", "/vendors?skip=0&limit=10", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "ITWorks")

	// ---- 3. SHOW ----
	req, _ = http.NewRequest("GET", "/vendors/"+strconv.Itoa(id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "ITWorks")

	// ---- 4. UPDATE ----
	updateBody := map[string]interface{}{
		"name":                "ITHelper",
		"countries_supported": []string{"KSA", "UAE"},
		"services_offered":    []string{"hiring"},
		"rating":              4.5,
		"response_sla_hours":  3,
	}
	jsonBody, _ = json.Marshal(updateBody)
	req, _ = http.NewRequest("PUT", "/vendors/"+strconv.Itoa(id), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "ITHelper")

	// ---- 5. DELETE ----
	req, _ = http.NewRequest("DELETE", "/vendors/"+strconv.Itoa(id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}
