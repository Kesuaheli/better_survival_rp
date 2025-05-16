package item

// TypeModel represents the simplest form of an item model
type TypeModel struct {
	Model string
}

// Type implements Model
func (m TypeModel) Type() ModelType {
	return TYPE_MODEL
}

// MarshalJSON implements json.Marshaler
func (model TypeModel) MarshalJSON() ([]byte, error) {
	return []byte(`{"type":"` + string(TYPE_MODEL) + `","model":"` + model.Model + `"}`), nil
}
