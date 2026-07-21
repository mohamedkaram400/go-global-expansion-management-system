package end_2_end

import (
	"bytes"
	"encoding/json"
	// "fmt"

	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
	repositories "github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories/v1"
	services "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	handlers "github.com/mohamedkaram400/go-global-expansion-management-system/internal/delivery/http/v1"
	"github.com/stretchr/testify/assert"
)

func setupUserRouter() *gin.Engine {
	db := conn.SetupTestDB()
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()
	r.GET("/users", userHandler.Index)
	r.GET("/users/:id", userHandler.Show)
	r.POST("/users", userHandler.Create)
	r.PUT("/users/:id", userHandler.Update)
	r.DELETE("/users/:id", userHandler.Destroy)

	return r
}

func TestGetAllUsers(t *testing.T) {
	router := setupUserRouter()

	req, _ := http.NewRequest("GET", "/users?skip=0&limit=10", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetUserById(t *testing.T) {
	router := setupUserRouter()

	req, _ := http.NewRequest("GET", "/users/2", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCreateUser(t *testing.T) {
	router := setupUserRouter()

	body := map[string]string{
		"name":     "Ali",
		"email":    "ali_" + uuid.NewString() + "@example.com",
		"Role":     "Support",
		"password": "qazwsx123",
	}
	jsonBody, _ := json.Marshal(body)

	log.Println(jsonBody)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Ali")

}

func TestUpdateUser(t *testing.T) {
	router := setupUserRouter()

	// First, create a user to update
	createBody := map[string]interface{}{
		"name":     "Ahmed",
		"email":    "ali_" + uuid.NewString() + "@example.com",
		"Role":     "Support",
		"password": "secret",
	}

	jsonCreate, _ := json.Marshal(createBody)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonCreate))
	req.Header.Set("Content-Type", "application/json")

	respCreate := httptest.NewRecorder()
	router.ServeHTTP(respCreate, req)

	// fmt.Println(respCreate)

	assert.Equal(t, http.StatusCreated, respCreate.Code)
	// assert.Contains(t, respCreate.Body.String(), "Ahmed")

	// Now extract the ID from response
	var created map[string]interface{}
	_ = json.Unmarshal(respCreate.Body.Bytes(), &created)
	// t.Log("Create response:", respCreate.Body.String())

	data := created["data"].(map[string]interface{})
	id := int(data["id"].(float64))

	// Update request
	updateBody := map[string]interface{}{
		"name":     "Ali",
		"email":    "ali_" + uuid.NewString() + "@example.com",
		"Role":     "Support",
		"password": "secret",
	}

	jsonUpdate, _ := json.Marshal(updateBody)

	reqUpdate, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(id), bytes.NewBuffer(jsonUpdate))
	reqUpdate.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, reqUpdate)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Ali")
}

func TestDeleteUser(t *testing.T) {
	router := setupUserRouter()

	req, _ := http.NewRequest("DELETE", "/users/4", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}

func TestUserHandler_EndToEnd(t *testing.T) {
	router := setupUserRouter()

	// ---- 1. CREATE ----
	createBody := map[string]interface{}{
		"name":     "Ali",
		"email":    "ali_" + uuid.NewString() + "@example.com",
		"Role":     "Support",
		"password": "qazwsx123",
	}
	jsonBody, _ := json.Marshal(createBody)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Ali")

	var createResp map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &createResp)
	data := createResp["data"].(map[string]interface{})
	id := int(data["id"].(float64))

	// ---- 2. INDEX ----
	req, _ = http.NewRequest("GET", "/users?skip=0&limit=10", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Ali")

	// ---- 3. SHOW ----
	req, _ = http.NewRequest("GET", "/users/"+strconv.Itoa(id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Ali")

	// ---- 4. UPDATE ----
	updateBody := map[string]interface{}{
		"name":     "Ali Updated",
		"email":    "ali_" + uuid.NewString() + "@example.com",
		"Role":     "Support",
		"password": "qazwsx123",
	}
	jsonBody, _ = json.Marshal(updateBody)
	req, _ = http.NewRequest("PUT", "/users/"+strconv.Itoa(id), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Ali Updated")

	// ---- 5. DELETE ----
	req, _ = http.NewRequest("DELETE", "/users/"+strconv.Itoa(id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNoContent, resp.Code)
}
