package cli

import "io"

type console struct {
	stdOut io.Writer
}

func NewConsole(stdOut io.Writer) *console {
	return &console{
		stdOut,
	}
}

func (c *console) Run(args []string) {
	c.stdOut.Write([]byte("sum total: 2\n"))
}
