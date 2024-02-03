package req

import (
	"net/http"
	"testing"
)

var r = NewRequest()

var SetHttpMethodTests = []struct {
	name      string
	testField string
	expected  string
}{
	{
		name:      "Set method to GET",
		testField: http.MethodGet,
		expected:  http.MethodGet,
	},
	{
		name:      "Set method to POST",
		testField: http.MethodPost,
		expected:  http.MethodPost,
	},
	{
		name:      "Set method to PUT",
		testField: http.MethodPut,
		expected:  http.MethodPut,
	},
	{
		name:      "Set method to DELETE",
		testField: http.MethodDelete,
		expected:  http.MethodDelete,
	},
	{
		name:      "Set method to PATCH",
		testField: http.MethodPatch,
		expected:  http.MethodPatch,
	},
	{
		name:      "Set method to HEAD",
		testField: http.MethodHead,
		expected:  http.MethodHead,
	},
	{
		name:      "Set method to OPTIONS",
		testField: http.MethodOptions,
		expected:  http.MethodOptions,
	},
	{
		name:      "Set method to TRACE",
		testField: http.MethodTrace,
		expected:  http.MethodTrace,
	},
}

var setHeaderTests = []struct {
	name      string
	testField interface{}
	expected  interface{}
}{
	{
		name:      "Test default method",
		testField: baseRequest.Method,
		expected:  http.MethodGet,
	},
}
var setCookieTests = []struct {
	name      string
	testField interface{}
	expected  interface{}
}{
	{
		name:      "Test default method",
		testField: baseRequest.Method,
		expected:  http.MethodGet,
	},
}

func TestSettingHttpMethod(t *testing.T) {
	for _, tc := range SetHttpMethodTests {
		t.Run(tc.name, func(t *testing.T) {
			r.setHttpMethod(tc.testField)
			if tc.expected != tc.testField {
				t.Errorf("Expected %s, got %s", tc.expected, tc.testField)
			}
		})
	}
}
