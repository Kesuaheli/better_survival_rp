package model

// Model represents a model
type Model struct {
	Parent   string            `json:"parent"`
	Textures map[string]string `json:"textures"`
}

// SimpleModel returns a model definition with a single texture layer.
func SimpleModel(texture string) Model {
	return Model{
		Parent:   "minecraft:item/generated",
		Textures: map[string]string{"layer0": texture},
	}
}
