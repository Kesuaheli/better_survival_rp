package item

import (
	"bytes"
	"encoding/json"
)

type SelectPropertyComponent struct {
	Component   string
	SelectCases []SelectCase
	Fallback    Model
}

// Type implements Model
func (selComp SelectPropertyComponent) Type() ModelType {
	return TYPE_SELECT
}

// MarshalJSON implements json.Marshaler
func (selComp SelectPropertyComponent) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.WriteString(`{` +
		`"type":"` + string(TYPE_SELECT) + `",` +
		`"property":"` + string(SELECT_PROPERTY_COMPONENT) + `",` +
		`"component":"` + selComp.Component + `",` +
		`"cases":[`,
	)
	for i, selCase := range selComp.SelectCases {
		if i > 0 {
			buf.WriteString(",")
		}

		b, err := json.Marshal(selCase)
		if err != nil {
			return nil, err
		}
		buf.Write(b)
	}

	buf.WriteString(`],"fallback":`)
	b, err := json.Marshal(selComp.Fallback)
	if err != nil {
		return nil, err
	}
	buf.Write(b)
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// Property implements Select
func (selComp SelectPropertyComponent) Property() SelectProperty {
	return SELECT_PROPERTY_COMPONENT
}

// Cases implements Select
func (selComp SelectPropertyComponent) Cases() []SelectCase {
	return selComp.SelectCases
}
