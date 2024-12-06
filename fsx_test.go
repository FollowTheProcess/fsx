package fsx_test

import (
	"testing"

	"github.com/FollowTheProcess/fsx"
)

func TestHello(t *testing.T) {
	got := fsx.Hello()
	want := "Hello fsx"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
