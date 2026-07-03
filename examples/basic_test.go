package echo_test

import (
	gloo "github.com/gloo-foo/framework/patterns"

	command "github.com/gloo-foo/cmd-echo"
)

func ExampleEcho_basic() {
	// echo "Hello World"
	if err := gloo.Run(
		command.Echo("Hello World"),
	); err != nil {
		panic(err)
	}
	// Output:
	// Hello World
}
