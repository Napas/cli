package cli

import (
	"context"
	"io"
)

type Command interface {
	Name() string
	Run(ctx context.Context, in Input, out io.Writer) error
}

type CommandWithCategory interface {
	Category() string
}

type CommandWithDescription interface {
	Description() string
}

type CommandWithFlags interface {
	Flags() []Flag
}

type CommandWithArguments interface {
	Arguments() []Argument
}

type commands []Command

func (c commands) GetByName(name string) Command {
	for _, cmd := range c {
		if cmd.Name() == name {
			return cmd
		}
	}

	return nil
}
