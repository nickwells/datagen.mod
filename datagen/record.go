package datagen

import "fmt"

// Record describes a record
type Record struct {
	name   string
	fields []*Field
}

// NewRecord constructs and returns a new Record. The Fields given should be
// in the order wanted in the final record. If two fields have the same name
// a nil pointer is returned with an error describing which fields have the
// same name; hidden fields, having no name, are not checked.
func NewRecord(name string, flds ...*Field) (*Record, error) {
	checkForUnique := make(map[string]int, len(flds))
	for i, f := range flds {
		if f.hidden {
			continue
		}

		if idx, ok := checkForUnique[f.Name()]; ok {
			return nil, fmt.Errorf(
				"fields %d and %d share the same name: %q",
				i, idx, f.Name())
		}

		checkForUnique[f.Name()] = i
	}

	return &Record{
		name:   name,
		fields: flds,
	}, nil
}

// AddFields adds the passed fields to the record. It returns an error if any
// of the new fields have a non-unique name; hidden fields, having no name,
// are not checked.
func (r *Record) AddFields(flds ...*Field) error {
	checkForUniqueOrig := make(map[string]int, len(r.fields))
	checkForUniqueNew := make(map[string]int, len(flds))

	for i, f := range r.fields {
		if f.hidden {
			continue
		}

		checkForUniqueOrig[f.Name()] = i
	}

	for i, f := range flds {
		if f.hidden {
			continue
		}

		if idx, ok := checkForUniqueOrig[f.Name()]; ok {
			return fmt.Errorf(
				"new field %d shares the same name as original field %d: %q",
				i, idx, f.Name())
		}

		if idx, ok := checkForUniqueNew[f.Name()]; ok {
			return fmt.Errorf(
				"new fields %d and %d share the same name: %q",
				i, idx, f.Name())
		}

		checkForUniqueNew[f.Name()] = i
	}

	r.fields = append(r.fields, flds...)

	return nil
}

// Generate will return a slice of strings generated from the fields. Hidden
// fields are ignored.
func (r Record) Generate() []string {
	rval := make([]string, 0, len(r.fields))
	for _, f := range r.fields {
		if f.hidden {
			continue
		}

		rval = append(rval, f.g.String())
	}

	return rval
}

// GenerateTitles will return a slice of strings generated from the field
// names. Hidden fields are ignored.
func (r Record) GenerateTitles() []string {
	rval := make([]string, 0, len(r.fields))
	for _, f := range r.fields {
		if f.hidden {
			continue
		}

		rval = append(rval, f.Name())
	}

	return rval
}

// GenerateAsMap will return a map of field names to strings generated from
// the fields.
func (r Record) GenerateAsMap() map[string]string {
	rval := make(map[string]string, len(r.fields))
	for _, f := range r.fields {
		if f.hidden {
			continue
		}

		rval[f.Name()] = f.g.String()
	}

	return rval
}

// Next moves all the fields to their next value including any hidden fields
func (r Record) Next() {
	for _, f := range r.fields {
		f.g.Next()
	}
}
