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
			request.SetPathValue("taskId", test.requestPathvalue)

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
		requestPayload  string
		requestPathvalue string
		expectedCode    int
		currentTaskList []Task
	}
	baseURL := "/task/"

	Tasks = []Task{}

	tests := []test{
		{
			name:            "Initial task",
			requestPayload:  `{"Detail": "first task"}`,
			requestPathvalue: "0",
			expectedCode:    http.StatusCreated,
			currentTaskList: []Task{{Detail: "first task"}},
		},
		{
			name:            "Empty task",
			requestPayload:  `{"Detail": ""}`,
			requestPathvalue: "0",
			expectedCode:    http.StatusNotAcceptable,
			currentTaskList: []Task{{Detail: "first task"}},
		},
		{
			name:            "Invalid task struct",
			requestPayload:  `{"Detl": ""}`,
			requestPathvalue: "0",
			expectedCode:    http.StatusBadRequest,
			currentTaskList: []Task{{Detail: "first task"}},
		},
		{
			name:            "Add more task",
			requestPayload:  `{"Detail": "2nd task"}`,
			requestPathvalue: "1",
			expectedCode:    http.StatusCreated,
			currentTaskList: []Task{{Detail: "first task"},{Detail: "2nd task"}},
		},
		{
			name:            "Tasks in array",
			requestPayload:  `[{"Detail": "2nd task"}]`,
			requestPathvalue: "0",
			expectedCode:    http.StatusBadRequest,
			currentTaskList: []Task{{Detail: "first task"},{Detail: "2nd task"}},
		},
		{
			name:            "Disallow multiple tasks",
			requestPayload:  `[{"Detail": "1nd task"},{"Detail": "2nd task"}]`,
			requestPathvalue: "0",
			expectedCode:    http.StatusBadRequest,
			currentTaskList: []Task{{Detail: "first task"},{Detail: "2nd task"}},
		},
		{
			name:            "Add task at index larger than length",
			requestPayload:  `{"Detail": "big task"}`,
			requestPathvalue: "10",
			expectedCode:    http.StatusCreated,
			currentTaskList: []Task{{Detail: "first task"},{Detail: "2nd task"},{Detail: "big task"}},
		},
		{
			name:            "Add task at negative index",
			requestPayload:  `{"Detail": "negative task"}`,
			requestPathvalue: "-1",
			expectedCode:    http.StatusBadRequest,
			currentTaskList: []Task{{Detail: "first task"},{Detail: "2nd task"},{Detail: "big task"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var payload = []byte(test.requestPayload)

			request := httptest.NewRequest(http.MethodPost, baseURL+test.requestPathvalue, bytes.NewBuffer(payload))
			request.SetPathValue("taskId", test.requestPathvalue)
			request.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			CreateTaskAtIndex(w, request)

			response := w.Result()

			if response.StatusCode != test.expectedCode {
				t.Errorf("%s expected status code %d but received %d", test.name, test.expectedCode, response.StatusCode)
			}
			// check if task is updated correctly
			if !reflect.DeepEqual(Tasks, test.currentTaskList) {
				t.Errorf("%s expected current task list %v but is %v", test.name, test.currentTaskList, Tasks)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	type test struct {
		name             string
		requestPathvalue string
		expectedCode     int
		currentTaskList []Task
	}
	baseURL := "/task/"

	Tasks = []Task{
		{
			Detail: "Number 1",
		},
		{
			Detail: "Number 2",
		},
		{
			Detail: "Number 3",
		},
	}

	tests := []test{
		{
			name:             "Valid index",
			requestPathvalue: "0",
			expectedCode:     http.StatusOK,
			currentTaskList: []Task{{Detail: "Number 2"},{Detail: "Number 3"}},
		},
		{
			name:             "Valid index",
			requestPathvalue: "1",
			expectedCode:     http.StatusOK,
			currentTaskList: []Task{{Detail: "Number 2"}},
		},
		{
			name:             "Negative index",
			requestPathvalue: "-1",
			expectedCode:     http.StatusBadRequest,
			currentTaskList: []Task{{Detail: "Number 2"}},
		},
		{
			name:             "Invalid index",
			requestPathvalue: "3",
			expectedCode:     http.StatusNotFound,
			currentTaskList: []Task{{Detail: "Number 2"}},
		},
		{
			name:             "Non integer index",
			requestPathvalue: "hello",
			expectedCode:     http.StatusBadRequest,
			currentTaskList: []Task{{Detail: "Number 2"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodDelete, baseURL+test.requestPathvalue, nil)
			request.SetPathValue("taskId", test.requestPathvalue)

			w := httptest.NewRecorder()

			DeleteTask(w, request)

			response := w.Result()

			if response.StatusCode != test.expectedCode {
				t.Errorf("expected status code %d but received %d", test.expectedCode, response.StatusCode)
			}
			
			// check if task is updated correctly
			if !reflect.DeepEqual(Tasks, test.currentTaskList) {
				t.Errorf("%s expected current task list %v but is %v", test.name, test.currentTaskList, Tasks)
			}
		})
	}
}


