package oscalls

import (
	"strings"
	"testing"
)

/* func TestGetData(t *testing.T) {
	got := GetData()
	want := "HAPPY NEW YEAR!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
} */
func TestGetDataIntegration(t *testing.T) {
	got := GetData(getXMLFromCommand())
	want := "HAPPY NEW YEAR!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetData(t *testing.T) {
	input := strings.NewReader(`<payload>
	<message>what a load of lovely!</message>
</payload>`)
	got := GetData(input)
	want := "WHAT A LOAD OF LOVELY!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
