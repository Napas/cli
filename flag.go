package cli

type Flag interface {
	Name() string
	Alias() string
	Description() string
	Validators() []Validator
}
