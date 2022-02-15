package solution

import "testing"

func TestSolution(t *testing.T) {
	want := "Hello 🗺️ !"

	if got := GetMessage(); got != want {
		t.Errorf("GetMessage() = %q, want %q", got, want)
	}
}
