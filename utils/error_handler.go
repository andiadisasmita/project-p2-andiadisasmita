package utils

import "github.com/labstack/echo/v4"

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// NewErrorResponse creates a standardized error response
func NewErrorResponse(message, details string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Details: details,
	}
}

// CustomHTTPErrorHandler handles HTTP errors for the Echo framework
func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := echo.ErrInternalServerError.Code
	message := "Internal Server Error"
	details := err.Error()

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if m, ok := he.Message.(string); ok {
			message = m
		}
	}

	c.JSON(code, NewErrorResponse(message, details))
}
