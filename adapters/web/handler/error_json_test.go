package handler_test

import (
	"github.com/stretchr/testify/require"
	"github.com/filipedtristao/hexagonal-architecture/adapters/web/handler"
	"testing"
)

func TestJsonError(t *testing.T) {
	msg := "Test error"
	expected := `{"error":"Test error"}`
	actual := handler.JsonError(msg)

	require.Equal(t, expected, string(actual))
}