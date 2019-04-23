package sample

import "testing"

func TestHello(t *testing.T) {
	actual := Hello("World")
	expected := "Hello, World"
	if actual != expected {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}
