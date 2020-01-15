package cli

type Validator interface {
	Validate(value string) error
}
