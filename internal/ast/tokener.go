package ast

type Tokener interface {
	// returns the token for a given AST node
	Token() string
}
