package api_test

import (
	"testing"

	"github.com/Shenon69/cryptomasters/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Error("Expected error, got nil")
	}
}