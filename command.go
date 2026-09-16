package command

import (
	"context"
	"strings"

	gloo "github.com/gloo-foo/framework"
)

// Text is a single echo argument — one word of the output line.
type Text string

// Echo returns a Source that emits its arguments joined by spaces as a single line.
func Echo(args ...Text) gloo.Source[[]byte] {
	words := make([]string, len(args))
	for i, arg := range args {
		words[i] = string(arg)
	}
	return echoSource{text: []byte(strings.Join(words, " "))}
}

type echoSource struct{ text []byte }

func (s echoSource) Stream(_ context.Context) gloo.Stream[[]byte] {
	return gloo.StreamOf(s.text)
}
