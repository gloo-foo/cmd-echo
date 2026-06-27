package echo_test

import (
	. "github.com/gloo-foo/cmd-echo"
	gloo "github.com/gloo-foo/framework/patterns"
)

func ExampleEcho_multiple() {
	// echo "one" "two" "three"
	gloo.MustRun(
		Echo("one", "two", "three"),
	)
	// Output:
	// one two three
}
