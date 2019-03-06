package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	s := new(Service)

	r := httptest.NewRequest(http.MethodPost, "/todos", nil)
	rw := httptest.NewRecorder()

	s.handleCreateTodo(rw, r)

	require.Equal(t, http.StatusOK, rw.Result().StatusCode)
}
