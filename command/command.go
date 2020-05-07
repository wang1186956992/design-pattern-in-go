package command

import (
	"fmt"
	"time"
)

type Command interface {
	Info() string
}

type TimeCommand struct {
	Start time.Time
}

func (t *TimeCommand) Info() string {
	return fmt.Sprintf("%v", time.Since(t.Start))
}

type HelloCommand struct{}

func (h *HelloCommand) Info() string {
	return "Hello, World"
}

func NewHelloCommand() *HelloCommand {
	return new(HelloCommand)
}

func NewTimeCommand(start time.Time) *TimeCommand {
	return &TimeCommand{
		Start: start,
	}
}

type CommandChain struct {
	Next *CommandChain
	Cmd  Command
}

func (c *CommandChain) Execute() {
	res := c.Cmd.Info()
	fmt.Printf("Command执行结果 %v\n", res)
	if c.Next != nil {
		c.Next.Execute()
	}
}
