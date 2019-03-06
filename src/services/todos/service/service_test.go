package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"services/todos/model"
	"services/todos/model/mocks"
	"services/utils/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	tcs := []struct {
		Description     string
		GivenBody       CreateTodoRequest
		ExpectedStatus  int
		VerifyTodoSaved bool
	}{
		{
			Description:    "invalid/body empty",
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "invalid/name empty",
			GivenBody: CreateTodoRequest{
				Until:       time.Now().Add(1 * time.Hour),
				Description: "Joe owes me money",
			},
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "invalid/name too long",
			GivenBody: CreateTodoRequest{
				Name:        rand.String(101),
				Until:       time.Now().Add(1 * time.Hour),
				Description: "Joe owes me money",
			},
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "invalid/description too long",
			GivenBody: CreateTodoRequest{
				Name:        rand.String(100),
				Until:       time.Now().Add(1 * time.Hour),
				Description: rand.String(301),
			},
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "valid",
			GivenBody: CreateTodoRequest{
				Name:        "Call Joe",
				Until:       time.Now().Add(1 * time.Hour).UTC(),
				Description: "Joe owes me money",
			},
			ExpectedStatus:  http.StatusOK,
			VerifyTodoSaved: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Description, func(t *testing.T) {
			storageMock := &mocks.Storage{}
			storageMock.On("Add", mock.Anything).Return(nil)
			sut := NewService(storageMock)

			data, err := json.Marshal(tc.GivenBody)
			require.NoError(t, err)

			rw := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(data))
			sut.handleCreateTodo(rw, r)

			require.Equal(t, tc.ExpectedStatus, rw.Result().StatusCode)
			if tc.VerifyTodoSaved {
				storageMock.AssertCalled(t, "Add", mock.MatchedBy(func(todo model.Todo) bool {
					return assert.Equal(t, tc.GivenBody.Name, string(todo.Name)) &&
						assert.Equal(t, tc.GivenBody.Description, string(todo.Description)) &&
						assert.Equal(t, tc.GivenBody.Until, todo.Until)
				}))
			}
		})
	}
}

func TestGet(t *testing.T) {
	tcs := []struct {
		Description      string
		GivenSavedTodos  []model.Todo
		ExpectedStatus   int
		ExpectedResponse GetTodosResponse
	}{
		{
			Description:    "empty",
			ExpectedStatus: http.StatusOK,
			ExpectedResponse: GetTodosResponse{
				Todos: []Todo{},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Description, func(t *testing.T) {
			storageStub := &mocks.Storage{}
			storageStub.On("GetAll").Return(tc.GivenSavedTodos, nil)
			sut := NewService(storageStub)

			r := httptest.NewRequest(http.MethodGet, "/todos", nil)
			rw := httptest.NewRecorder()

			sut.handleGetTodos(rw, r)

			require.Equal(t, tc.ExpectedStatus, rw.Result().StatusCode)

			bytes, err := ioutil.ReadAll(rw.Result().Body)
			require.NoError(t, err)

			var actualResponse GetTodosResponse
			err = json.Unmarshal(bytes, &actualResponse)
			require.NoError(t, err)

			require.Equal(t, tc.ExpectedResponse, actualResponse)
		})
	}
}
