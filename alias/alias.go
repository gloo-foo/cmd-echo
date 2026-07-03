// Package alias provides an unprefixed name for the echo command.
//
//	import echo "github.com/gloo-foo/cmd-echo/alias"
//	echo.Echo("hello", "world")
package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-echo"
)

// Text is a single echo argument — one word of the output line.
type Text = command.Text

// Echo emits its arguments joined by spaces as a single line; see the command
// package for the semantics.
func Echo(args ...Text) gloo.Source[[]byte] { return command.Echo(args...) }
