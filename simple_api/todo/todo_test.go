package todo

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestJsonResponse(t *testing.T) {
	expectedContentType := "application/json"
	type Expected struct {
		body []byte
		code int
	}
	type test struct {
		input    []Task
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
		if !bytes.Equal(responseBody, test.expected.body) {
			t.Errorf("expected body %s, got %s", test.expected.body, responseBody)
		}
	}
}

func TestGetTaskByIndex(t *testing.T) {
	type test struct {
		name             string
		requestPathvalue string
		expectedCode     int
		expectedBody     string
	}
	baseURL := "/task/"

	Tasks = []Task{
		{
			Detail: "Number 1",
		},
		{
			Detail: "Number 2",
		},
	}

	tests := []test{
		{
			name:             "Valid index",
			requestPathvalue: "0",
			expectedCode:     http.StatusOK,
			expectedBody:     `{"Detail":"Number 1"}`,
		},
		{
			name:             "Valid index",
			requestPathvalue: "1",
			expectedCode:     http.StatusOK,
			expectedBody:     `{"Detail":"Number 2"}`,
		},
		{
			name:             "Negative index",
			requestPathvalue: "-1",
			expectedCode:     http.StatusBadRequest,
			expectedBody:     ``,
		},
		{
			name:             "Invalid index",
			requestPathvalue: "3",
			expectedCode:     http.StatusNotFound,
			expectedBody:     ``,
		},
		{
			name:             "Non integer index",
			requestPathvalue: "hello",
			expectedCode:     http.StatusBadRequest,
			expectedBody:     ``,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, baseURL+test.requestPathvalue, nil)
			request.SetPathValue("id", test.requestPathvalue)

			w := httptest.NewRecorder()

			GetTaskByIndex(w, request)

			response := w.Result()

			if response.StatusCode != test.expectedCode {
				t.Errorf("expected status code %d but received %d", test.expectedCode, response.StatusCode)
			}

			responseBody, _ := io.ReadAll(response.Body)
			if string(responseBody) != test.expectedBody {
				t.Errorf("expected status code %s but received %s", test.expectedBody, string(responseBody))
			}
		})
	}
}

func TestCreateTask(t *testing.T) {
	type test struct {
		name            string
		requestMethod   string
		requestPayload  string
		expectedCode    int
		currentTaskList []Task
	}
	baseURL := "/task/"

	Tasks = []Task{}

	tests := []test{
		{
			name:            "Initial task",
			requestMethod:   http.MethodPost,
			requestPayload:  `{"Detail": "first task"}`,
			expectedCode:    http.StatusCreated,
			currentTaskList: []Task{{Detail: "first task"}},
		},
		{
			name:            "Empty task",
			requestMethod:   http.MethodPost,
			requestPayload:  `{"Detail": ""}`,
			expectedCode:    http.StatusNotAcceptable,
			currentTaskList: []Task{{Detail: "first task"}},
		},
		{
			name:            "Invalid task struct",
			requestMethod:   http.MethodPost,
			requestPayload:  `{"Detl": ""}`,
			expectedCode:    http.StatusBadRequest,
			currentTaskList: []Task{{Detail: "first task"}},
		},
		{
			name:            "Add more task",
			requestMethod:   http.MethodPost,
			requestPayload:  `{"Detail": "2nd task"}`,
			expectedCode:    http.StatusCreated,
			currentTaskList: []Task{{Detail: "first task"},{Detail: "2nd task"}},
		},
	}

	for _, test := range tests {
		var payload = []byte(test.requestPayload)

		request := httptest.NewRequest(test.requestMethod, baseURL, bytes.NewBuffer(payload))
		request.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		CreateTask(w, request)

		response := w.Result()

		if response.StatusCode != test.expectedCode {
			t.Errorf("%s expected status code %d but received %d", test.name, test.expectedCode, response.StatusCode)
		}
		if !reflect.DeepEqual(Tasks, test.currentTaskList) {
			t.Errorf("%s expected current task list %v but is %v", test.name, test.currentTaskList, Tasks)
		}
	}
}
