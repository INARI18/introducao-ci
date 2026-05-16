package greeting

import "testing"

func TestMessage(t *testing.T) {
	got := Message("estudantes")
	want := "Olá, estudantes!"

	if got != want {
		t.Fatalf("Message() = %q, want %q", got, want)
	}
}
