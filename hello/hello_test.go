package hello_test

import (
	"testing"

	"github.com/alemelomeza/cli-actions-example/hello"
)

func TestHello(t *testing.T) {
	got := hello.Hello()
	want := "Hello, world!"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
