package cmd

import (
	"bytes"
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewURLComponents_normal(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected urlComponents
	}{
		{
			name: "scheme and host",
			url:  "https://example.com",
			expected: urlComponents{
				Scheme:   "https",
				Host:     "example.com",
				User:     &userInfo{},
				Port:     "",
				Path:     "",
				Fragment: "",
				Queries:  map[string]interface{}{},
			},
		},
		{
			name: "full",
			url:  "https://user:pass@example.com:1234/path1/path2/?q1=v1&q2=v2-1&q2=v2-2#frag",
			expected: urlComponents{
				Scheme: "https",
				User: &userInfo{
					Username: "user",
					Password: "pass",
				},
				Host:     "example.com",
				Port:     "1234",
				Path:     "/path1/path2/",
				Fragment: "frag",
				Queries: map[string]interface{}{
					"q1": "v1",
					"q2": []string{"v2-1", "v2-2"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputURL, _ := url.Parse(tt.url)
			actual := newURLComponents(inputURL)
			assert.Equal(t, &tt.expected, actual)
		})
	}
}

func TestExecute_normal(t *testing.T) {
	inputRawURL := "https://example.com"
	resultJson := `
  {
    "scheme": "https",
    "user": {
      "username": "",
      "password": ""
    },
    "host": "example.com",
    "port": "",
    "path": "",
    "fragment": "",
    "queries": {}
  }
  `

	// Test multiple argument passing.
	tests := []struct {
		name      string
		arguments []string
		stdin     string
		expected  string
	}{
		{
			name:      "run by argument",
			arguments: []string{inputRawURL},
			stdin:     "",
			expected:  resultJson,
		},
		{
			name:      "run by stdin (no argument)",
			arguments: []string{},
			stdin:     inputRawURL,
			expected:  resultJson,
		},
		{
			name:      "run by stdin (with '-' argument)",
			arguments: []string{"-"},
			stdin:     inputRawURL,
			expected:  resultJson,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBuffer := bytes.NewBufferString("")
			resultBuffer := bytes.NewBufferString("")
			cmd := newRootCmd(inputBuffer, resultBuffer)

			cmd.SetArgs(tt.arguments)
			inputBuffer.Write([]byte(tt.stdin))
			cmd.Execute()

			actual, _ := ioutil.ReadAll(resultBuffer)
			require.JSONEq(t, tt.expected, string(actual))
		})
	}
}
