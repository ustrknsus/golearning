package greeting

import "testing"

func TestHello(t *testing.T) {
	if got := Hello("Go"); got != "Hello, Go!" {
		t.Fatalf("unexpected greeting: %q", got)
	}
}
