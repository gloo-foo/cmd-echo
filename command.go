package command

import (
	"context"
	"strings"

	gloo "github.com/gloo-foo/framework"
)

// Echo returns a Source that emits its arguments joined by spaces as a single line.
func Echo(args ...string) gloo.Source[[]byte] {
	return &echoSource{text: []byte(strings.Join(args, " "))}
}

type echoSource struct{ text []byte }

func (s *echoSource) Stream(_ context.Context) gloo.Stream[[]byte] {
	return gloo.StreamOf(s.text)
}
