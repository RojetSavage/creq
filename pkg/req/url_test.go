package req

import (
	"net/http"
	"net/url"
	"testing"
)

const (
	BASE_URL = "http://example.com/path?query#fragment"
)

func TestChangeUri(t *testing.T) {
	// Helper function to compare URLs
	compareURL := func(t *testing.T, expected, actual *url.URL) {
		if expected.String() != actual.String() {
			t.Errorf("Expected URL %s, got %s", expected, actual)
		}
	}

	tests := []struct {
		name     string
		portion  string
		s        string
		expected string
	}{
		{
			name:     "Change scheme",
			portion:  "scheme",
			s:        "https",
			expected: "https://example.com/path?query#fragment",
		},
		{
			name:     "Change host",
			portion:  "host",
			s:        "newhost",
			expected: "http://newhost/path?query#fragment",
		},
		{
			name:     "Change port",
			portion:  "port",
			s:        "8080",
			expected: "http://example.com:8080/path?query#fragment",
		},
		{
			name:     "Change path",
			portion:  "path",
			s:        "/newpath",
			expected: "http://example.com/newpath?query#fragment",
		},
		{
			name:     "Change query",
			portion:  "query",
			s:        "newquery",
			expected: "http://example.com/path?newquery#fragment",
		},
		{
			name:     "Change fragment",
			portion:  "fragment",
			s:        "newfragment",
			expected: "http://example.com/path?query#newfragment",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url, err := url.Parse(BASE_URL)
			expectedUrl, err := url.Parse(tc.expected)

			req := &http.Request{URL: url}

			err = ChangeUri(req, tc.portion, tc.s)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			compareURL(t, expectedUrl, req.URL)
		})
	}
}
