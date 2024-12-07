package todo

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"
)

func TestJsonResponse(t *testing.T) {
	expectedContentType := "application/json"
	type Expected struct{
		body []byte
		code int
	}
	type test struct{
		input []Task
		expected Expected
	}

	tests := []test{
		{
			input: []Task{
				{
					Detail: "Number 1",
				},
				{
					Detail: "Number 2",
				},
			},
			expected: Expected{
				body: []byte(`[{"Detail":"Number 1"},{"Detail":"Number 2"}]`),
				code: 200,
			},
		},
		{
			input: []Task{},
			expected: Expected{
				body: []byte(`[]`),
				code: 200,
			},
		},
		
	}

	for _, test := range tests {

		w := httptest.NewRecorder()
	
		jsonResponse(w, test.input)
	
		response := w.Result()
	
		if response.StatusCode != test.expected.code {
			t.Errorf("expected status code %d, got %d", test.expected.code, response.StatusCode)
		}

		if response.Header.Get("Content-Type") != expectedContentType {
			t.Errorf("expected response content type %s, got %s", expectedContentType, response.Header.Get("Content-Type"))
		}

		responseBody, _ := io.ReadAll(response.Body)
		if !bytes.Equal(responseBody, test.expected.body){
			t.Errorf("expected body %s, got %s", test.expected.body, responseBody)
		}
	}

}