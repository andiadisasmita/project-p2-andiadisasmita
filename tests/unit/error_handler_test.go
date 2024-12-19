package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-andiadisasmita/project-p2-andiadisasmita/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCustomHTTPErrorHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := func(err error, c echo.Context) {
		utils.CustomHTTPErrorHandler(err, c)
	}

	// Trigger a 400 error
	handler(echo.NewHTTPError(http.StatusBadRequest, "Bad Request"), c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
