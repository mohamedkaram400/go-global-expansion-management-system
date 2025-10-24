package end_2_end

import (
	"bytes"
	"encoding/json"

	// "log"
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

func setupProjectRouter() *gin.Engine {
	db := conn.SetupTestDB()
	projectRepo := repositories.NewProjectRepo(db)
	projectService := services.NewProjectService(projectRepo)
	projectHandler := handlers.NewProjectHandler(projectService) 

	r := gin.Default()
	r.GET("/projects", projectHandler.Index)
	r.GET("/projects/:id", projectHandler.Show)
	r.POST("/projects", projectHandler.Create)
	r.PUT("/projects/:id", projectHandler.Update)
	r.DELETE("/projects/:id", projectHandler.Destroy)

	return r
}

func TestGetAllProjects(t *testing.T) {
	router := setupProjectRouter()

	req, _ := http.NewRequest("GET", "/projects?skip=0&limit=10", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetProjectById(t *testing.T) {
	router := setupProjectRouter()

	req, _ := http.NewRequest("GET", "/projects/2", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCreateProject(t *testing.T) {
	router := setupProjectRouter()

	body := map[string]interface{}{
		"country":          "Egypt",
		"service_needed":  []string{"legal", "hiring"},
		"budget":           40000,
		"client_id":        1,
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/projects", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	t.Log("Response Body:", resp.Body.String())

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Egypt")
}

func TestUpdateProject(t *testing.T) {
	router := setupProjectRouter()

	createBody := map[string]interface{}{
		"country":         "Egypt",
		"service_needed": []string{"legal", "hiring"},
		"budget":          40000,
		"client_id":       1,
	}
	jsonCreate, _ := json.Marshal(createBody)

	req, _ := http.NewRequest("POST", "/projects", bytes.NewBuffer(jsonCreate))
	req.Header.Set("Content-Type", "application/json")

	respCreate := httptest.NewRecorder()
	router.ServeHTTP(respCreate, req)

	assert.Equal(t, http.StatusCreated, respCreate.Code)

	var created map[string]interface{}
	_ = json.Unmarshal(respCreate.Body.Bytes(), &created)

	if created["data"] == nil {
		t.Fatalf("Create failed: %v", created)
	}

	data := created["data"].(map[string]interface{})
	id := int(data["id"].(float64))

	updateBody := map[string]interface{}{
		"country":         "KSA",
		"services_needed": []string{"legal", "hiring"},
		"budget":          55000,
		"client_id":       1,
	}
	jsonUpdate, _ := json.Marshal(updateBody)

	reqUpdate, _ := http.NewRequest("PUT", "/projects/"+strconv.Itoa(id), bytes.NewBuffer(jsonUpdate))
	reqUpdate.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, reqUpdate)

	t.Log("Update response:", resp.Body.String())

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "KSA")
}

func TestDeleteProject(t *testing.T) {
	router := setupProjectRouter()

	req, _ := http.NewRequest("DELETE", "/projects/4", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	
	assert.Equal(t, http.StatusNoContent, resp.Code)
}


func TestProjectHandler_EndToEnd(t *testing.T) {
	router := setupProjectRouter()

	// ---- 1. CREATE ----
	createBody := map[string]interface{}{
		"country":         "Egypt",
		"service_needed": []string{"legal", "hiring"},
		"budget":          5000,
		"client_id":       1,
	}
	jsonBody, _ := json.Marshal(createBody)
	req, _ := http.NewRequest("POST", "/projects", bytes.NewBuffer(jsonBody))
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
	req, _ = http.NewRequest("GET", "/projects?skip=0&limit=10", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Egypt")


	// ---- 3. SHOW ----
	req, _ = http.NewRequest("GET", "/projects/"+strconv.Itoa(id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Egypt")


	// ---- 4. UPDATE ----
	updateBody := map[string]interface{}{
		"country":         "KSA",
		"service_needed": []string{"legal", "hiring"},
		"budget":          55000,
		"client_id":       1,
	}
	jsonBody, _ = json.Marshal(updateBody)
	req, _ = http.NewRequest("PUT", "/projects/"+strconv.Itoa(id), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "KSA")

	
	// ---- 5. DELETE ----
	req, _ = http.NewRequest("DELETE", "/projects/"+strconv.Itoa(id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}