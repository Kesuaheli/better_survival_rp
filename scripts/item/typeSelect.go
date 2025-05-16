package item

type SelectProperty string

const (
	SELECT_PROPERTY_COMPONENT SelectProperty = "minecraft:component"
)

// TypeSelect represents a select item model
type TypeSelect interface {
	Model
	SelectProperty() SelectProperty
	SelectCases() []SelectCase
}

type SelectCase struct {
	When  string `json:"when"`
	Model Model  `json:"model"`
}
