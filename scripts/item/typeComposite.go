package item

import (
	"bytes"
	"encoding/json"
)

// TypeComposite represents a composite item model
type TypeComposite struct {
	Models []Model
}

// Type implements Model
func (t TypeComposite) Type() ModelType {
	return TYPE_COMPOSITE
}

// MarshalJSON implements json.Marshaler
func (comp TypeComposite) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.WriteString(`{"type":"` + string(TYPE_COMPOSITE) + `","models":[`)
	for i, model := range comp.Models {
		if i > 0 {
			buf.WriteString(",")
		}

		b, err := json.Marshal(model)
		if err != nil {
			return nil, err
		}
		buf.Write(b)
	}
	buf.WriteString("]}")
	return buf.Bytes(), nil

}
