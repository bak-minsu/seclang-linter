package ast

// defines SecLang directive
type Directive interface {
	Tokener

	// Returns the possible options for the directive
	// in the order they should be defined
	Options() []DirectiveOption
}
