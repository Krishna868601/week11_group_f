package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFileServer(t *testing.T) {
	// Create a file server pointing to the "templates" directory
	fs := http.FileServer(http.Dir("./templates"))
	ts := httptest.NewServer(fs)
	defer ts.Close()

	// Test cases
	tests := []struct {
		description  string
		urlPath      string
		expectedCode int
		expectedMime string
	}{
		{
			description:  "Serve existing HTML file",
			urlPath:      "/index.html",
			expectedCode: http.StatusOK,
			expectedMime: "text/html",
		},
		{
			description:  "Handle non-existent file",
			urlPath:      "/nonexistent.html",
			expectedCode: http.StatusNotFound,
			expectedMime: "",
		},
		{
			description:  "Serve CSS file with correct MIME type",
			urlPath:      "/style.css",
			expectedCode: http.StatusOK,
			expectedMime: "text/css",
		},
		{
			description:  "Serve JavaScript file with correct MIME type",
			urlPath:      "/script.js",
			expectedCode: http.StatusOK,
			expectedMime: "application/javascript",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// Send GET request to the test server
			resp, err := http.Get(ts.URL + test.urlPath)
			if err != nil {
				t.Fatalf("Failed to send GET request: %v", err)
			}
			defer resp.Body.Close()

			// Validate HTTP status code
			if resp.StatusCode != test.expectedCode {
				t.Errorf("Expected status %d, got: %d", test.expectedCode, resp.StatusCode)
			}

			// Validate MIME type, if applicable
			if test.expectedMime != "" {
				contentType := resp.Header.Get("Content-Type")
				if !strings.HasPrefix(contentType, test.expectedMime) &&
					!(test.expectedMime == "application/javascript" && strings.HasPrefix(contentType, "text/javascript")) {
					t.Errorf("Expected MIME type %s, got: %s for URL %s", test.expectedMime, contentType, ts.URL+test.urlPath)
				}
			}

		})
	}
}
