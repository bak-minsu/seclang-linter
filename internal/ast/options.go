package ast

type DirectiveOption interface {
	Tokener

	// represents the regular expression for the option
	OptionRegex() string
}
