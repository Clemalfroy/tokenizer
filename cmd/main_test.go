package main

import (
	"testing"

	"git.vpgrp.io/vp-labs/lab-42/front-search/tokenizer"
)

const texte string = "La larve est le premier stade de développement de l'individu après l'éclosion de l'œuf ou la naissance chez un grand nombre d'espèces animales ayant un développement post-embryonnaire appelé « indirect ». C'est le cas dans la plupart des embranchements, notamment chez les arthropodes (insectes, crustacés), les mollusques, les annélides et les Chordés (urochordés, « poissons », amphibiens, marsupiaux). Chez certaines espèces pondant de très nombreux œufs, la plupart des larves seront mangées par des prédateurs avant d'atteindre le stade adulte."

func BenchmarkTokenalamano(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, word := range tokenizer.Tokenize(texte) {
			_ = word
		}
	}
}
