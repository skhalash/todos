package storage

import (
	"services/todos/model"
	"services/utils/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	sut := New()

	for i := 0; i < 10; i++ {
		todo, _ := model.NewTodo(rand.String(50), rand.String(100), time.Now(), time.Now())
		sut.Add(*todo)
	}

	got, err := sut.GetAll()
	require.NoError(t, err)
	require.Len(t, got, 10)
}
