package usecase_test

import (
	"github.com/TolkinSL/go-practice/dev02/internal/usecase"
	"testing"
)

func TestHello(t *testing.T) {
	hello := usecase.NewHello()

	actual := hello.Say()
	expected := "Hello!"

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
