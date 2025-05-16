package main

import "local/item"

const CARD_BASE = "minecraft:item/paper"

func makeDeckCardCases() (cases []item.SelectCase) {

	for _, color := range cardDeckColors {
		for _, number := range cardDeckNumbers {
			cases = append(cases, makeCardCase(number, color))
		}
	}
	return cases
}

func makeCardCase(number cardNumber, color cardColor) item.SelectCase {
	return item.SelectCase{
		When: string(number) + string(color),
		Model: item.TypeComposite{
			Models: []item.Model{
				item.TypeModel{
					Model: CARD_BASE,
				},
				item.TypeModel{
					Model: "better_survival:item/card/" + number.Name(),
				},
				item.TypeModel{
					Model: "better_survival:item/card/" + color.Name(),
				},
			},
		},
	}
}
