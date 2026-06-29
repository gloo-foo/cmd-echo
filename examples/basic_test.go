package echo_test

import (
	gloo "github.com/gloo-foo/framework/patterns"

	command "github.com/gloo-foo/cmd-echo"
)

func ExampleEcho_basic() {
	// echo "Hello World"
	gloo.MustRun(
		command.Echo("Hello World"),
	)
	// Output:
	// Hello World
}
