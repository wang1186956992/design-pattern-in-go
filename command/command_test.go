package command

import (
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	helloCmd := NewHelloCommand()
	timeCmd := NewTimeCommand(time.Now())
	time.Sleep(time.Second)

	helloChain := CommandChain{
		Cmd: helloCmd,
	}
	timeChain := CommandChain{
		Next: &helloChain,
		Cmd:  timeCmd,
	}
	timeChain.Execute()
}
