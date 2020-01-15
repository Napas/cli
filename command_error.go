package cli

type CommandError interface {
	error
	ExitCode() int
}
