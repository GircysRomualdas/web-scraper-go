package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetImagesFromHTML(t *testing.T) {
	cases := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "absolute image URL",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src="https://blog.boot.dev/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://blog.boot.dev/logo.png"},
		},
		{
			name:      "relative image URL",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://blog.boot.dev/logo.png"},
		},
		{
			name:     "multiple images",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<img src="/logo.png" alt="Logo">
				<img src="https://cdn.boot.dev/banner.jpg">
			</body></html>`,
			expected: []string{
				"https://blog.boot.dev/logo.png",
				"https://cdn.boot.dev/banner.jpg",
			},
		},
		{
			name:     "no images",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<p>No images here!</p>
			</body></html>`,
			expected: nil,
		},
		{
			name:     "invalid image src",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<img src=":\\invalidURL">
			</body></html>`,
			expected: nil,
		},
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Fatalf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
			}

			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected images %v, got %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
