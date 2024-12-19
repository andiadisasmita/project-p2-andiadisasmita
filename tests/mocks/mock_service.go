package tests

import (
	"errors"
)

// MockService simulates behaviors of a service layer for testing
type MockService struct {
	UserExists      bool
	UserCreateError error
	BoardgameError  error
	PaymentError    error
	RentalError     error
	ReviewError     error
}

func (ms *MockService) CreateUser(username, password string) error {
	if ms.UserExists {
		return errors.New("user already exists")
	}
	if ms.UserCreateError != nil {
		return ms.UserCreateError
	}
	return nil
}

func (ms *MockService) CreateBoardgame(title string) error {
	if ms.BoardgameError != nil {
		return ms.BoardgameError
	}
	return nil
}

func (ms *MockService) ProcessPayment(amount float64) error {
	if ms.PaymentError != nil {
		return ms.PaymentError
	}
	return nil
}

func (ms *MockService) RentBoardgame(userID, gameID uint) error {
	if ms.RentalError != nil {
		return ms.RentalError
	}
	return nil
}

func (ms *MockService) SubmitReview(userID, gameID uint, content string) error {
	if ms.ReviewError != nil {
		return ms.ReviewError
	}
	return nil
}
