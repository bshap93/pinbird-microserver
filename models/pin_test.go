package models

import (
	"testing"
	"time"
)

func TestTimeCreated(t *testing.T) {
	myPin := Pin{
		Href:        "www.google.com",
		Description: "Google",
		Time:        "2022-11-21T01:49:01Z"}

	got := myPin.timeCreated()
	want := time.Date(2022, 11, 21, 01, 49, 01, 0, time.UTC)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
