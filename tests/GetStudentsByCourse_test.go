package tests

import (
	"github.com/go-resty/resty/v2"
	"testing"
)

func TestGetStudentsByCourse(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8081/course/course")
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}

	if resp.StatusCode() != 200 {
		t.Fatalf("Received status code %d, expected 200", resp.StatusCode())
	}

	t.Logf("Test completed successfully")
}
