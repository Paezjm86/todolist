package todolist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileStore(t *testing.T) {
	assert := assert.New(t)
	store := &FileStore{FileLocation: "todos.json"}
	store.Load()
	assert.Equal(store.Data[0].Subject, "this is the first subject", "")
}
