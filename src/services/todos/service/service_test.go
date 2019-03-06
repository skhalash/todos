package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	tcs := []struct {
		Description    string
		GivenBody      CreateTodoDto
		ExpectedStatus int
	}{
		{
			Description: "valid",
			GivenBody: CreateTodoDto{
				Name:        "Call Joe",
				Until:       time.Now().Add(1 * time.Hour),
				Description: "Joe owes me money",
			},
			ExpectedStatus: http.StatusOK,
		},
	}

	for _, tc := range tcs {
		s := new(Service)

		data, err := json.Marshal(tc.GivenBody)
		require.NoError(t, err)

		r := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(data))

		rw := httptest.NewRecorder()

		s.handleCreateTodo(rw, r)

		require.Equal(t, tc.ExpectedStatus, rw.Result().StatusCode)
	}
}
