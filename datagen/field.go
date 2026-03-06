package datagen

// Field describes a field in a record
type Field struct {
	name   string
	g      Generator
	hidden bool
}

// Name returns the field name
func (f Field) Name() string {
	return f.name
}

// NewField returns a new Field with the name and Generator set
func NewField(name string, g Generator) *Field {
	return &Field{
		name: name,
		g:    g,
	}
}

// NewHiddenField returns a new Field with the Generator set and the
// hidden flag set to true. Hidden fields do not have names
func NewHiddenField(g Generator) *Field {
	return &Field{
		g:      g,
		hidden: true,
	}
}
