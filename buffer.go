package main

import (
	"io"
)

type CtrlBuffer struct {
	Ctrl     bool
	Positive io.Writer
	Negative io.Writer
}

func (c *CtrlBuffer) Write(p []byte) (n int, err error) {
	if c.Ctrl {
		return c.Positive.Write(p)
	}
	return c.Negative.Write(p)
}
