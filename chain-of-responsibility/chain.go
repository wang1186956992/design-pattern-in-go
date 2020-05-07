package chain

import (
	"fmt"
	"io"
)

var (
	LOG_MSG = "%v is logging\n"
)

type ILog interface {
	Log()
}

type Log struct {
	Name string
	W    io.Writer
}

func (l *Log) Log() {
	str := fmt.Sprintf(LOG_MSG, l.Name)
	l.W.Write([]byte(str))
}

type LogChain struct {
	Next *LogChain
	L    ILog
}

func (c *LogChain) Handle() {
	c.L.Log()
	if c.Next != nil {
		c.Next.Handle()
	}
}
