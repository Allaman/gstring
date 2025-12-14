package main

import (
	"strings"
	"unicode"
)

type SpellingAlphabet map[rune]string

var (
	// ICAO (NATO) phonetic alphabet
	icaoAlphabet = SpellingAlphabet{
		'A': "Alfa", 'B': "Bravo", 'C': "Charlie", 'D': "Delta",
		'E': "Echo", 'F': "Foxtrot", 'G': "Golf", 'H': "Hotel",
		'I': "India", 'J': "Juliett", 'K': "Kilo", 'L': "Lima",
		'M': "Mike", 'N': "November", 'O': "Oscar", 'P': "Papa",
		'Q': "Quebec", 'R': "Romeo", 'S': "Sierra", 'T': "Tango",
		'U': "Uniform", 'V': "Victor", 'W': "Whiskey", 'X': "X-ray",
		'Y': "Yankee", 'Z': "Zulu",
		'0': "Zero", '1': "One", '2': "Two", '3': "Three",
		'4': "Four", '5': "Five", '6': "Six", '7': "Seven",
		'8': "Eight", '9': "Nine",
	}

	// German Postal Spelling Alphabet
	postalAlphabet = SpellingAlphabet{
		'A': "Anton", 'Ä': "Ärger", 'B': "Berta", 'C': "Cäsar",
		'D': "Dora", 'E': "Emil", 'F': "Friedrich", 'G': "Gustav",
		'H': "Heinrich", 'I': "Ida", 'J': "Julius", 'K': "Kaufmann",
		'L': "Ludwig", 'M': "Martha", 'N': "Nordpol", 'O': "Otto",
		'Ö': "Ökonom", 'P': "Paula", 'Q': "Quelle", 'R': "Richard",
		'S': "Samuel", 'T': "Theodor", 'U': "Ulrich", 'Ü': "Übermut",
		'V': "Viktor", 'W': "Wilhelm", 'X': "Xanthippe", 'Y': "Ypsilon",
		'Z': "Zacharias",
		'0': "Null", '1': "Eins", '2': "Zwei", '3': "Drei",
		'4': "Vier", '5': "Fünf", '6': "Sechs", '7': "Sieben",
		'8': "Acht", '9': "Neun",
	}
)

// spellString converts a string to its phonetic spelling using the given alphabet
func spellString(input string, alphabet SpellingAlphabet) string {
	var result strings.Builder

	for _, char := range input {
		upperChar := unicode.ToUpper(char)

		if spelled, exists := alphabet[upperChar]; exists {
			if result.Len() > 0 {
				result.WriteString(" - ")
			}
			result.WriteString(spelled)
		} else if unicode.IsSpace(char) {
			// Handle spaces explicitly
			if result.Len() > 0 {
				result.WriteString(" - ")
			}
			result.WriteString("[Space]")
		} else if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			// For punctuation and special chars, output them directly
			if result.Len() > 0 {
				result.WriteString(" - ")
			}
			result.WriteString(string(char))
		}
		// Unknown Unicode characters are silently skipped
	}

	return result.String()
}

type spellCmd struct {
	Input    string `arg:"" optional:"" help:"String to spell (reads from STDIN if not provided)"`
	Alphabet string `default:"postal" short:"a" help:"Alphabet to use: 'postal' (default) or 'icao'"`
}

func (c *spellCmd) Run(globals *Globals) error {
	var in string
	var err error

	if c.Input != "" {
		in = c.Input
	} else {
		in, err = readFromSTDIN()
		if err != nil {
			return err
		}
	}

	var alphabet SpellingAlphabet
	switch c.Alphabet {
	case "icao":
		alphabet = icaoAlphabet
	case "postal":
		alphabet = postalAlphabet
	default:
		alphabet = postalAlphabet
	}

	printOutput(spellString(in, alphabet), globals.Trim)
	return nil
}
