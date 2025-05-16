package main

import (
	"encoding/json"
	"fmt"
	"os"

	"local/item"
	"local/model"
)

const (
	HOME_DIRECTORY = "../../"

	OUTPUT_DIRECTORY_MC    = HOME_DIRECTORY + "assets/minecraft/"
	OUTPUT_DIRECTORY_BS    = HOME_DIRECTORY + "assets/better_survival/"
	OUTPUT_ITEM_DRECTORY   = OUTPUT_DIRECTORY_MC + "items/"
	OUTPUT_ITEM_FILE       = OUTPUT_ITEM_DRECTORY + "paper.json"
	OUTPUT_MODEL_DIRECTORY = OUTPUT_DIRECTORY_BS + "models/item/card/"

	RESOURCE_LOCATION_TEXTURES = "better_survival:item/card/"
)

func main() {
	var err error
	if err = buildItemFile(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = buildModelFiles(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func buildItemFile() (err error) {
	cardDeckModel := item.SelectPropertyComponent{
		Component:   "minecraft:custom_name",
		SelectCases: makeDeckCardCases(),
		Fallback: item.TypeModel{
			Model: "minecraft:item/paper",
		},
	}

	buf, err := json.MarshalIndent(item.Item{Model: cardDeckModel}, "", "\t")
	if err != nil {
		return err
	}
	err = os.MkdirAll(OUTPUT_ITEM_DRECTORY, 0755)
	if err != nil {
		return err
	}
	err = os.WriteFile(OUTPUT_ITEM_FILE, buf, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Wrote card deck item definition to '%s' with %d cases\n", OUTPUT_ITEM_FILE, len(cardDeckModel.SelectCases))
	return nil
}

func buildModelFiles() (err error) {
	for _, color := range cardDeckColors {
		buf, err := json.MarshalIndent(model.SimpleModel(RESOURCE_LOCATION_TEXTURES+color.Name()), "", "\t")
		if err != nil {
			return err
		}

		err = os.MkdirAll(OUTPUT_MODEL_DIRECTORY, 0755)
		if err != nil {
			return err
		}

		err = os.WriteFile(OUTPUT_MODEL_DIRECTORY+color.Name()+".json", buf, 0644)
		if err != nil {
			return err
		}
	}

	for _, number := range cardDeckNumbers {
		buf, err := json.MarshalIndent(model.SimpleModel(RESOURCE_LOCATION_TEXTURES+number.Name()), "", "\t")
		if err != nil {
			return err
		}
		err = os.WriteFile(OUTPUT_MODEL_DIRECTORY+number.Name()+".json", buf, 0644)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Wrote %d card deck model files (%d colors, %d numbers) to '%s'\n",
		len(cardDeckColors)+len(cardDeckNumbers),
		len(cardDeckColors), len(cardDeckNumbers),
		OUTPUT_MODEL_DIRECTORY,
	)

	return nil
}
