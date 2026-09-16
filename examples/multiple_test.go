package echo_test

import (
	gloo "github.com/gloo-foo/framework/patterns"

	command "github.com/gloo-foo/cmd-echo"
)

func ExampleEcho_multiple() {
	// echo "one" "two" "three"
	if err := gloo.Run(
		command.Echo("one", "two", "three"),
	); err != nil {
		panic(err)
	}
	// Output:
	// one two three
}
