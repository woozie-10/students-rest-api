package tests

import (
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestDeleteStudentByUsername(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Delete("http://localhost:8081/students/username")
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}

	if resp.StatusCode() != 200 {
		t.Fatalf("Received status code %d, expected 200", resp.StatusCode())
	}

	t.Log("Test completed successfully")
}