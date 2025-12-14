package main

import "testing"

func TestSpellString_ICAO(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple word uppercase",
			input:    "HELLO",
			expected: "Hotel - Echo - Lima - Lima - Oscar",
		},
		{
			name:     "Lowercase input",
			input:    "world",
			expected: "Whiskey - Oscar - Romeo - Lima - Delta",
		},
		{
			name:     "Mixed case",
			input:    "TeSt",
			expected: "Tango - Echo - Sierra - Tango",
		},
		{
			name:     "With numbers",
			input:    "abc123",
			expected: "Alfa - Bravo - Charlie - One - Two - Three",
		},
		{
			name:     "With space",
			input:    "A B",
			expected: "Alfa - [Space] - Bravo",
		},
		{
			name:     "Multiple spaces",
			input:    "S O S",
			expected: "Sierra - [Space] - Oscar - [Space] - Sierra",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "X",
			expected: "X-ray",
		},
		{
			name:     "With punctuation",
			input:    "S.O.S",
			expected: "Sierra - . - Oscar - . - Sierra",
		},
		{
			name:     "All numbers",
			input:    "0123456789",
			expected: "Zero - One - Two - Three - Four - Five - Six - Seven - Eight - Nine",
		},
		{
			name:     "Complete alphabet",
			input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expected: "Alfa - Bravo - Charlie - Delta - Echo - Foxtrot - Golf - Hotel - India - Juliett - Kilo - Lima - Mike - November - Oscar - Papa - Quebec - Romeo - Sierra - Tango - Uniform - Victor - Whiskey - X-ray - Yankee - Zulu",
		},
		{
			name:     "With exclamation",
			input:    "Help!",
			expected: "Hotel - Echo - Lima - Papa - !",
		},
		{
			name:     "With question mark",
			input:    "Why?",
			expected: "Whiskey - Hotel - Yankee - ?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := spellString(tt.input, icaoAlphabet)
			if result != tt.expected {
				t.Errorf("spellString(%q, icaoAlphabet) = %q; want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestSpellString_Postal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "German characters with umlauts",
			input:    "ÄÖÜ",
			expected: "Ärger - Ökonom - Übermut",
		},
		{
			name:     "German characters lowercase umlauts",
			input:    "äöü",
			expected: "Ärger - Ökonom - Übermut",
		},
		{
			name:     "Simple German word",
			input:    "OTTO",
			expected: "Otto - Theodor - Theodor - Otto",
		},
		{
			name:     "Mixed German and numbers",
			input:    "AB12",
			expected: "Anton - Berta - Eins - Zwei",
		},
		{
			name:     "Lowercase German word",
			input:    "berlin",
			expected: "Berta - Emil - Richard - Ludwig - Ida - Nordpol",
		},
		{
			name:     "German word with space",
			input:    "GUTEN TAG",
			expected: "Gustav - Ulrich - Theodor - Emil - Nordpol - [Space] - Theodor - Anton - Gustav",
		},
		{
			name:     "All German numbers",
			input:    "0123456789",
			expected: "Null - Eins - Zwei - Drei - Vier - Fünf - Sechs - Sieben - Acht - Neun",
		},
		{
			name:     "Mixed umlauts and regular",
			input:    "MÜNCHEN",
			expected: "Martha - Übermut - Nordpol - Cäsar - Heinrich - Emil - Nordpol",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := spellString(tt.input, postalAlphabet)
			if result != tt.expected {
				t.Errorf("spellString(%q, postalAlphabet) = %q; want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestSpellString_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		alphabet SpellingAlphabet
		expected string
	}{
		{
			name:     "Only spaces",
			input:    "   ",
			alphabet: icaoAlphabet,
			expected: "[Space] - [Space] - [Space]",
		},
		{
			name:     "Only punctuation",
			input:    "!@#",
			alphabet: icaoAlphabet,
			expected: "! - @ - #",
		},
		{
			name:     "Mixed punctuation and letters",
			input:    "A-B-C",
			alphabet: icaoAlphabet,
			expected: "Alfa - - - Bravo - - - Charlie",
		},
		{
			name:     "Tab character",
			input:    "A\tB",
			alphabet: icaoAlphabet,
			expected: "Alfa - [Space] - Bravo",
		},
		{
			name:     "Newline character",
			input:    "A\nB",
			alphabet: icaoAlphabet,
			expected: "Alfa - [Space] - Bravo",
		},
		{
			name:     "Only special characters",
			input:    ".,;:!?",
			alphabet: icaoAlphabet,
			expected: ". - , - ; - : - ! - ?",
		},
		{
			name:     "Empty with postal alphabet",
			input:    "",
			alphabet: postalAlphabet,
			expected: "",
		},
		{
			name:     "Single space",
			input:    " ",
			alphabet: postalAlphabet,
			expected: "[Space]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := spellString(tt.input, tt.alphabet)
			if result != tt.expected {
				t.Errorf("spellString(%q, alphabet) = %q; want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}
