package alias_test

import (
	"context"
	"slices"
	"testing"

	echo "github.com/gloo-foo/cmd-echo/alias"
	gloo "github.com/gloo-foo/framework"
)

// The alias package re-exports the Echo constructor under an unprefixed name. A
// mis-wired re-export (Echo bound to the wrong function) compiles cleanly, so
// only behavior can prove the wiring. This test exercises the re-export and
// asserts the echo output it must produce: the arguments joined by single
// spaces, emitted as a single line.

func collect(t *testing.T, src gloo.Source[[]byte]) []string {
	t.Helper()
	ctx := context.Background()
	items, err := gloo.Collect(ctx, src.Stream(ctx))
	if err != nil {
		t.Fatal(err)
	}
	lines := make([]string, len(items))
	for i, item := range items {
		lines[i] = string(item)
	}
	return lines
}

func TestAlias_EchoJoinsArgsWithSpaces(t *testing.T) {
	got := collect(t, echo.Echo("hello", "world"))
	if !slices.Equal(got, []string{"hello world"}) {
		t.Fatalf("got %q, want %q", got, []string{"hello world"})
	}
}
