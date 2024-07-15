package handler

import (
	"testing"

	"paper-test/storage/mocks"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	mockStorage := new(mocks.IStorage)
	handler := NewHandler(mockStorage)

	assert.NotNil(t, handler)
	assert.Equal(t, mockStorage, handler.Storage)
}
