package ast

// represents SecLang as an AST
type AST struct {
	// holds the directives in the order they are defined
	Directives []Directive
}
