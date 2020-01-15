package cli

import (
	"context"
	"io"
)

type Application struct {
	commands commands
	output   io.Writer
}

func (app *Application) Run(args []string) error {
	return app.RunCtx(context.Background(), args)
}

func (app *Application) RunCtx(ctx context.Context, args []string) error {

}
