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
	logs []ILog
}

func (c *LogChain) Handle() {
	for _, log := range c.logs {
		log.Log()
	}
}

func (c *LogChain) AddLog(log ILog) {
	c.logs = append(c.logs, log)
}
