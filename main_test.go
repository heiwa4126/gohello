package main

import (
	"testing"

	"github.com/kami-zh/go-capturer"
)

func Test1(t *testing.T) {

	//--- test1
	want := "Hello world!\n"
	out := capturer.CaptureStdout(func() {
		sayAnything()
	})

	if out != want {
		t.Errorf("Wrong word...\n")
	}
}
