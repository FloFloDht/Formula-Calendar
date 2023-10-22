package main

import "time"

const (
	Formula1    string = "Formula 1"
	Formula2    string = "Formula 2"
	Formula3    string = "Formula 3"
	FormulaAcad string = "Formula 1 Academy"
)

type Tracks struct {
	Date         time.Time
	Name         string
	Localisation string
	Category     []CategoryAuto
}

type CategoryAuto struct {
	Name string
}


// Il y a 4 catégorie de monoplace possible sur un circuit
// Chaque circuit contiendra au moins la F1
// Un circuit c'est une date, une lieu, un nom et un liste de catégorie de monoplace

