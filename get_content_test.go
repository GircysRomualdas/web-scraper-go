package main

import (
	"testing"
)

func TestGetH1FromHTML(t *testing.T) {
	cases := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "basic H1",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "empty H1",
			inputBody: "<html><body><p>No h1 here</p></body></html>",
			expected:  "",
		},
		{
			name:      "empty body",
			inputBody: "<html><body><p></p></body></html>",
			expected:  "",
		},
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected %q, got %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	cases := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name: "main paragraph preferred",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>Main paragraph.</p>
				</main>
			</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "fallback to first paragraph",
			inputBody: `<html><body>
				<p>First paragraph outside main.</p>
				<p>Second paragraph outside main.</p>
			</body></html>`,
			expected: "First paragraph outside main.",
		},
		{
			name:      "empty paragraph",
			inputBody: `<html><body><p></p></body></html>`,
			expected:  "",
		},
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected %q, got %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}
