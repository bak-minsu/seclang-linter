package parse

// Represents the root of the parse tree
type File struct {
	// name represents the path of the file
	name string
	// list of directives found in the file
	Directives []*Directive
}

func (f *File) Name() string {
	return f.name
}
