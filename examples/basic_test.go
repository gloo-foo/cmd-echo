package echo_test

import (
	. "github.com/gloo-foo/cmd-echo"
	gloo "github.com/gloo-foo/framework/patterns"
)

func ExampleEcho_basic() {
	// echo "Hello World"
	gloo.MustRun(
		Echo("Hello World"),
	)
	// Output:
	// Hello World
}
