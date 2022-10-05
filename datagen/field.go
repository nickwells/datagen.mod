package datagen

// Field describes a field in a record
type Field struct {
	name string
	g    Generator
}

// Name returns the field name
func (f Field) Name() string {
	return f.name
}

// NewField returns a new Field with the name and Generator set
func NewField(name string, g Generator) *Field {
	return &Field{name: name, g: g}
}
