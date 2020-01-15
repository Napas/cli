package cli

type Input interface {
	HasOption(name string) bool
	OptionString(name string) string
	OptionInt(name string) int

	HasArgument(name string) bool
	ArgumentString(name string) string
	ArgumentInt(name string) int
}

type input struct {
	options   argsMap
	arguments argsMap
}

func newFromArgs(args []string) Input {
	return &input{
		options:   argsMap{},
		arguments: argsMap{},
	}
}

func (i *input) HasOption(name string) bool {
	return i.options.Has(name)
}

func (i *input) OptionString(name string) string {
	return i.options.String(name)
}

func (i *input) OptionInt(name string) int {
	return i.options.Int(name)
}

func (i *input) HasArgument(name string) bool {
	return i.arguments.Has(name)
}

func (i *input) ArgumentString(name string) string {
	return i.arguments.String(name)
}

func (i *input) ArgumentInt(name string) int {
	return i.arguments.Int(name)
}

type argsMap map[string]interface{}

func (m argsMap) Has(name string) bool {
	_, ok := m[name]

	return ok
}

func (m argsMap) String(name string) string {
	return m[name].(string)
}

func (m argsMap) Int(name string) int {
	return m[name].(int)
}
