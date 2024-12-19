package utils

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/go-resty/resty/v2"
)

// CreateInvoice creates an invoice using Xendit's API
func CreateInvoice(rentalID uint, amount float64, description, callbackURL string) (string, error) {
	client := resty.New()

	// Define the Xendit API endpoint
	apiURL := "https://api.xendit.co/v2/invoices"

	// Prepare the request payload
	payload := map[string]interface{}{
		"external_id":          rentalID,
		"amount":               amount,
		"description":          description,
		"success_redirect_url": callbackURL,
	}

	// Execute the request
	resp, err := client.R().
		SetBasicAuth(os.Getenv("XENDIT_SECRET_API_KEY"), "").
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(apiURL)

	if err != nil {
		return "", err
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 201 {
		return "", errors.New("failed to create invoice: " + resp.String())
	}

	// Parse response and extract the invoice URL
	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", err
	}

	// Extract and return the invoice URL
	if invoiceURL, ok := result["invoice_url"].(string); ok {
		return invoiceURL, nil
	}

	return "", errors.New("unable to extract invoice URL from response")
}

// CheckInvoiceStatus retrieves the status of an invoice
func CheckInvoiceStatus(invoiceID string) (string, error) {
	client := resty.New()

	apiURL := "https://api.xendit.co/v2/invoices/" + invoiceID

	resp, err := client.R().
		SetBasicAuth(os.Getenv("XENDIT_SECRET_API_KEY"), "").
		Get(apiURL)

	if err != nil {
		return "", err
	}

	if resp.StatusCode() != 200 {
		return "", errors.New("failed to fetch invoice status: " + resp.String())
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", err
	}

	if status, ok := result["status"].(string); ok {
		return status, nil
	}

	return "", errors.New("unable to extract status from response")
}
