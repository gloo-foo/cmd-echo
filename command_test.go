package command_test

import (
	"context"
	"testing"

	command "github.com/gloo-foo/cmd-echo"
	gloo "github.com/gloo-foo/framework"
)

func TestEcho_SingleArg(t *testing.T) {
	src := command.Echo("hello")
	got, err := gloo.Collect(context.Background(), src.Stream(context.Background()))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || string(got[0]) != "hello" {
		t.Errorf("got %q, want [\"hello\"]", got)
	}
}

func TestEcho_MultipleArgs(t *testing.T) {
	src := command.Echo("hello", "world")
	got, err := gloo.Collect(context.Background(), src.Stream(context.Background()))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || string(got[0]) != "hello world" {
		t.Errorf("got %q, want [\"hello world\"]", got)
	}
}

func TestEcho_NoArgs(t *testing.T) {
	src := command.Echo()
	got, err := gloo.Collect(context.Background(), src.Stream(context.Background()))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || string(got[0]) != "" {
		t.Errorf("got %q, want [\"\"]", got)
	}
}
