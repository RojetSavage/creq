package req

import (
	"net/url"
	"testing"
)

var tests = []struct {
	name      string
	component string
	s         string
	original  string
	expected  string
}{
	{
		name:      "Change scheme",
		component: "scheme",
		s:         "https",
		original:  "http://example.com/path?query#fragment",
		expected:  "https://example.com/path?query#fragment",
	},
	{
		name:      "Change host",
		component: "host",
		s:         "localhost",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://localhost/path?query#fragment",
	},
	{
		name:      "Change host",
		component: "host",
		s:         "localhost.com.au",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://localhost.com.au/path?query#fragment",
	},
	{
		name:      "Add port",
		component: "port",
		s:         "8080",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://example.com:8080/path?query#fragment",
	},
	{
		name:      "Change port",
		component: "port",
		s:         "8080",
		original:  "http://example.com:8081/path?query#fragment",
		expected:  "http://example.com:8080/path?query#fragment",
	},
	{
		name:      "Change to invalid port",
		component: "port",
		s:         "808o",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://example.com/path?query#fragment",
	},
	{
		name:      "Change to invalid port (removing existing port)",
		component: "port",
		s:         "808o",
		original:  "http://example.com:8080/path?query#fragment",
		expected:  "http://example.com/path?query#fragment",
	}, {
		name:      "Add Path",
		component: "path",
		s:         "/newpath",
		original:  "http://example.com/?query#fragment",
		expected:  "http://example.com/newpath?query#fragment",
	},
	{
		name:      "Change path",
		component: "path",
		s:         "/newpath",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://example.com/newpath?query#fragment",
	},
	{
		name:      "Change path without providing leading slash",
		component: "path",
		s:         "newpath/user/id",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://example.com/newpath/user/id?query#fragment",
	}, {
		name:      "Change path",
		component: "path",
		s:         "/newpath/user/id/example",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://example.com/newpath/user/id/example?query#fragment",
	},
	{
		name:      "Change query",
		component: "query",
		s:         "user=1",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://example.com/path?user=1#fragment",
	},
	{
		name:      "Change query",
		component: "query",
		s:         "user=1&foo=bar",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://example.com/path?user=1&foo=bar#fragment",
	},
	{
		name:      "Remove query",
		component: "query",
		s:         "",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://example.com/path#fragment",
	},
	{
		name:      "Change fragment",
		component: "fragment",
		s:         "newfragment",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://example.com/path?query#newfragment",
	},
	{
		name:      "Remove fragment",
		component: "fragment",
		s:         "",
		original:  "http://example.com/path?query#fragment",
		expected:  "http://example.com/path?query",
	},
}

func compareURL(t *testing.T, expected, actual *url.URL) {
	if expected.String() != actual.String() {
		t.Errorf("Expected URL %s, got %s", expected.String(), actual.String())
	}
}

func TestChangeUri(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url, err := url.Parse(tc.original)
			expectedUrl, err := url.Parse(tc.expected)

			r := NewRequest()
			r.setUrl(tc.original)
			err = r.changeUri(tc.component, tc.s)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			compareURL(t, expectedUrl, r.URL)
		})
	}
}
