package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"services/todos/storage"
	"services/utils/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	tcs := []struct {
		Description     string
		GivenBody       CreateTodoDto
		ExpectedStatus  int
		VerifyTodoSaved bool
	}{
		{
			Description: "invalid/name empty",
			GivenBody: CreateTodoDto{
				Until:       time.Now().Add(1 * time.Hour),
				Description: "Joe owes me money",
			},
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "invalid/name too long",
			GivenBody: CreateTodoDto{
				Name:        rand.String(101),
				Until:       time.Now().Add(1 * time.Hour),
				Description: "Joe owes me money",
			},
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "invalid/description too long",
			GivenBody: CreateTodoDto{
				Name:        rand.String(100),
				Until:       time.Now().Add(1 * time.Hour),
				Description: rand.String(301),
			},
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Description: "valid",
			GivenBody: CreateTodoDto{
				Name:        "Call Joe",
				Until:       time.Now().Add(1 * time.Hour),
				Description: "Joe owes me money",
			},
			ExpectedStatus:  http.StatusOK,
			VerifyTodoSaved: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Description, func(t *testing.T) {
			storage := storage.New()
			sut := NewService(storage)

			data, err := json.Marshal(tc.GivenBody)
			require.NoError(t, err)

			r := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(data))

			rw := httptest.NewRecorder()

			sut.handleCreateTodo(rw, r)

			require.Equal(t, tc.ExpectedStatus, rw.Result().StatusCode)

			if tc.VerifyTodoSaved {
				all, err := storage.GetAll()
				require.NoError(t, err)
				require.Len(t, all, 1)
			}
		})
	}
}
