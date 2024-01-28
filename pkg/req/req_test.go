package req

import (
	"net/http"
	"testing"
)

var baseRequest = NewRequest()

var requestTests = []struct {
	name      string
	testField interface{}
	expected  interface{}
}{
	{
		name:      "Test default method",
		testField: baseRequest.Method,
		expected:  http.MethodGet,
	},
	{
		name:      "Test default URL",
		testField: baseRequest.URL.String(),
		expected:  "http://localhost:3001",
	},
	{
		name:      "Test default body",
		testField: baseRequest.Body,
		expected:  nil,
	},
}

func TestNewRequest(t *testing.T) {
	for _, tc := range requestTests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.testField != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, tc.testField)
			}
		})
	}
}
