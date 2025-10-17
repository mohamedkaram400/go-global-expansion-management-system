package end_2_end

import (
	"bytes"
	"encoding/json"

	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
	repositories "github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories/v1"
	services "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	handlers "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/v1"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	db := conn.SetupTestDB()
	clientRepo := repositories.NewClientRepo(db)
	clientService := services.NewClientService(clientRepo)
	clientHandler := handlers.NewClientHandler(clientService) 

	r := gin.Default()
	r.GET("/clients", clientHandler.Index)
	r.GET("/clients/:id", clientHandler.Show)
	r.POST("/clients", clientHandler.Create)
	r.PUT("/clients/:id", clientHandler.Update)
	r.DELETE("/clients/:id", clientHandler.Destroy)

	return r
}

func TestGetAllClients(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/clients?skip=0&limit=10", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetClientById(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/clients/2", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCreateClient(t *testing.T) {
	router := setupRouter()

	email := fmt.Sprintf("test_%d@example.com", time.Now().UnixNano())

	body := map[string]string{
		"company_name":  "Test Company",
		"contact_email": email,
		"password": "qazwsx123",
	}
	jsonBody, _ := json.Marshal(body)

	log.Println(jsonBody)

	req, _ := http.NewRequest("POST", "/clients", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test Company")
	
}

func TestUpdateClient(t *testing.T) {
	router := setupRouter()

	// First, create a client to update
	createBody := map[string]interface{}{
		"company_name":  "Original Company",
		"contact_email": "unique_" + uuid.NewString() + "@example.com",
		"password":      "secret",
	}
	jsonCreate, _ := json.Marshal(createBody)


	req, _ := http.NewRequest("POST", "/clients", bytes.NewBuffer(jsonCreate))
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
		"company_name":  "Updated Company",
		"contact_email": "update_" + uuid.NewString() + "@example.com",
		"password":      "secret",
	}

	jsonUpdate, _ := json.Marshal(updateBody)

	reqUpdate, _ := http.NewRequest("PUT", "/clients/"+strconv.Itoa(id), bytes.NewBuffer(jsonUpdate))
	reqUpdate.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, reqUpdate)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Updated Company")
}

func TestDeleteClient(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/clients/4", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	
	assert.Equal(t, http.StatusNoContent, resp.Code)
}

func TestClientHandler_EndToEnd(t *testing.T) {
	router := setupRouter()

	// ---- 1. CREATE ----
	createBody := map[string]interface{}{
		"company_name":  "Test Company",
		"contact_email": "test_" + uuid.NewString() + "@example.com",
		"password": "qazwsx123",

	}
	jsonBody, _ := json.Marshal(createBody)
	req, _ := http.NewRequest("POST", "/clients", bytes.NewBuffer(jsonBody))
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
	req, _ = http.NewRequest("GET", "/clients?skip=0&limit=10", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test Company")


	// ---- 3. SHOW ----
	req, _ = http.NewRequest("GET", "/clients/"+strconv.Itoa(id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test Company")


	// ---- 4. UPDATE ----
	updateBody := map[string]interface{}{
		"company_name":  "Updated Company",
		"contact_email": "updated@example.com",
		"password": "qazwsx123",
	}
	jsonBody, _ = json.Marshal(updateBody)
	req, _ = http.NewRequest("PUT", "/clients/"+strconv.Itoa(id), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Updated Company")

	
	// ---- 5. DELETE ----
	req, _ = http.NewRequest("DELETE", "/clients/"+strconv.Itoa(id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}