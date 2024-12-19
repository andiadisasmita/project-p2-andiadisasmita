package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/config"
	"github.com/a-andiadisasmita/project-p2-andiadisasmita/controllers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateReview(t *testing.T) {
	// Setup Echo and Database
	e := echo.New()
	err := config.InitializeDatabase()
	assert.NoError(t, err)

	// Mock HTTP Request
	payload := `{"boardgame_id":1,"user_id":1,"rating":5,"comment":"Great game!"}`
	req := httptest.NewRequest(http.MethodPost, "/reviews", bytes.NewReader([]byte(payload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test CreateReview Controller
	err = controllers.CreateReview(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		// Parse and Validate the Response
		var response map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Review created successfully", response["message"])
	}
}

func TestGetReviews(t *testing.T) {
	// Setup Echo and Database
	e := echo.New()
	err := config.InitializeDatabase()
	assert.NoError(t, err)

	// Mock HTTP Request
	req := httptest.NewRequest(http.MethodGet, "/reviews/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test GetReviews Controller
	err = controllers.GetReviews(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Parse and Validate the Response
		var response []map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	}
}
