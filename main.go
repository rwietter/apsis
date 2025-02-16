package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
)

type AstroTerm struct {
	Name       string
	NameJP     string
	Etymology  string
	Definition string
	Poetic     string
	Icon       string
}

var astroTerms = map[string]AstroTerm{
	"aphelion": {
		Name:       "aphelion",
		NameJP:     "遠日点",
		Etymology:  "από (\"away\") + ἥλιος (\"sun\")",
		Definition: "The point in a celestial body's orbit\nfarthest from the Sun.",
		Poetic:     "farthest from light\nclosest to void",
		Icon: `
  ⁎ ───────── ⁎
╱     ◌        ╲
   ◯     ☉   
╲              ╱
  ⁎ ───────── ⁎
`,
	},
	"perihelion": {
		Name:       "perihelion",
		NameJP:     "近日点",
		Etymology:  "περί (\"near\") + ἥλιος (\"sun\")",
		Definition: "The point in a celestial body's orbit\nclosest to the Sun.",
		Poetic:     "closest to light\nembrace of fire",
		Icon: `
⁎ ───────── ⁎
╱             ╲
    ◯ ☉          
╲             ╱
⁎ ───────── ⁎
`,
	},
	"apogee": {
		Name:       "apogee",
		NameJP:     "遠地点",
		Etymology:  "από (\"away\") + γῆ (\"earth\")",
		Definition: "The point in the Moon's orbit\nfarthest from the Earth.",
		Poetic:     "dancing with stars\nfar from home",
		Icon: `
⁎ ───── ⊙
╱         ╲
◯           ●
╲         ╱
⁎ ───── ⁎
`,
	},
	"perigee": {
		Name:       "perigee",
		NameJP:     "近地点",
		Etymology:  "περί (\"near\") + γῆ (\"earth\")",
		Definition: "The point in the Moon's orbit\nclosest to the Earth.",
		Poetic:     "return to earth\nclose to home",
		Icon: `
  ⁎ ───── ⊙
  ╱         ╲
●    ◯      
  ╲         ╱
  ⁎ ───── ⁎
`,
	},
}

func centerText(text string, width int) string {
	padding := (width - runewidth.StringWidth(text)) / 2
	if padding < 0 {
		padding = 0
	}
	return fmt.Sprintf("%s%s", strings.Repeat(" ", padding), text)
}

func displayTerm(term AstroTerm, termWidth int) {
	iconColor := color.New(color.FgWhite)
	titleColor := color.New(color.FgHiCyan)
	subtitleColor := color.New(color.FgCyan)
	definitionColor := color.New(color.FgHiRed)
	italicColor := color.New(color.FgMagenta, color.Italic)

	// Display icon theme
	for _, line := range strings.Split(term.Icon, "\n") {
		iconColor.Println(centerText(line, termWidth))
	}

	// Display term name and JP name
	titleColor.Println(centerText(term.Name, termWidth))
	subtitleColor.Println(centerText(term.NameJP, termWidth))
	fmt.Println()

	// Display etymology
	definitionColor.Println(centerText(term.Etymology, termWidth))
	fmt.Println()

	// Display definition
	for _, line := range strings.Split(term.Definition, "\n") {
		definitionColor.Println(centerText(line, termWidth))
	}
	fmt.Println()

	// Display poetic description
	for _, line := range strings.Split(term.Poetic, "\n") {
		italicColor.Println(centerText(line, termWidth))
	}
	fmt.Println()
}

// Randomly select a term from the list
func getRandomTerm() string {
	terms := make([]string, 0, len(astroTerms))
	for k := range astroTerms {
		terms = append(terms, k)
	}
	return terms[rand.Intn(len(terms))]
}

func main() {
	termPtr := flag.String("term", "", "Astronomical term to display (aphelion, perihelion, apogee, perigee)")
	randomPtr := flag.Bool("random", false, "Display a random astronomical term")
	flag.Parse()

	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Get the term to show
	var termToShow string
	if *randomPtr {
		termToShow = getRandomTerm()
	} else if *termPtr != "" {
		termToShow = strings.ToLower(*termPtr)
	} else {
		termToShow = getRandomTerm()
	}

	term, exists := astroTerms[termToShow]
	if !exists {
		fmt.Printf("Term '%s' not found. Available terms: aphelion, perihelion, apogee, perigee\n", termToShow)
		return
	}

	displayTerm(term, 50)
}
