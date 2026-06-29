package echo_test

import (
	gloo "github.com/gloo-foo/framework/patterns"

	command "github.com/gloo-foo/cmd-echo"
)

func ExampleEcho_multiple() {
	// echo "one" "two" "three"
	gloo.MustRun(
		command.Echo("one", "two", "three"),
	)
	// Output:
	// one two three
}
