package hello

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
func TestGreet(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions" {
		t.Errorf("Greet() = %s; Expected Hello GitHub actions", result)
	}
}

func TestGet(t *testing.T) {
	Get()
}
