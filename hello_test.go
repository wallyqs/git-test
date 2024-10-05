package hello

import "testing"

func TestGreet(t *testing.T) {
	t.Logf("Example test")
	p := &Person{Name: "Beto"}

	got := p.Greet()
	expected := "Hello Beto!!!"
	if got != expected {
		t.Errorf("Unexpected result: got: %s, expected: %s", got, expected)
	}
}
