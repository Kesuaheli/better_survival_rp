package item

// ModelType represents the type of an item model
type ModelType string

// Item model types
const (
	TYPE_COMPOSITE ModelType = "minecraft:composite"
	TYPE_MODEL     ModelType = "minecraft:model"
	TYPE_SELECT    ModelType = "minecraft:select"
)

// Model represents an item model
type Model interface {
	Type() ModelType
	MarshalJSON() ([]byte, error)
}
