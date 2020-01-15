package cli

type Argument interface {
	Name() string
	Description() string
	Validators() []Validator
}
