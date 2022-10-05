package datagen

// Record describes a record
type Record struct {
	name   string
	fields []*Field
}

// NewRecord constructs and returns a new Record
func NewRecord(name string, f ...*Field) *Record {
	return &Record{name: name, fields: f}
}

// AddFields adds the passed fields to the record
func (r *Record) AddFields(f ...*Field) {
	r.fields = append(r.fields, f...)
}

// Generate will return a slice of strings generated from the fields
func (r Record) Generate() []string {
	rval := make([]string, 0, len(r.fields))
	for _, f := range r.fields {
		rval = append(rval, f.g.Generate())
	}
	return rval
}

// GenerateTitles will return a slice of strings generated from the field names
func (r Record) GenerateTitles() []string {
	rval := make([]string, 0, len(r.fields))
	for _, f := range r.fields {
		rval = append(rval, f.Name())
	}
	return rval
}

// GenerateAsMap will return a map of field names to strings generated from
// the fields.
func (r Record) GenerateAsMap() map[string]string {
	rval := make(map[string]string, len(r.fields))
	for _, f := range r.fields {
		rval[f.Name()] = f.g.Generate()
	}
	return rval
}

// Next moves all the fields to their next value
func (r Record) Next() {
	for _, f := range r.fields {
		f.g.Next()
	}
}
