package tdata

import (
	"math/rand"
	"unicode"
	"unicode/utf8"
)

// RandomASCIIString returns random string consisting of ascii characters
func RandomASCIIString() string {
	return randomWord(0.1, 1, unicode.MaxASCII)
}

// RandomUtf8String returns random string unicode of ascii characters
func RandomUtf8String() string {
	return randomWord(0.1, 1, utf8.MaxRune)
}

func randomWord(p float64, startingCodepoint, endCodepoint int) string {
	result := ""
	for {
		codepoint := rand.Intn(endCodepoint-startingCodepoint+1) + startingCodepoint
		result = result + string(rune(codepoint))
		if rand.Float64() <= p {
			return result
		}
	}
}

// SampleStrings returns slice containing various strings, which can be used to test
// functions accepting strings as arguments
func SampleStrings() []string {
	return []string{
		"'",
		"{",
		"}",
		"`",
		`"`,
		"",
		",",
		";",
		"{}",
		" ",
		`/`,
		`//`,
		`\`,
		`\\`,
		"\n",
		`
`,
		"\\n",
		"mary",
		"had",
		"a",
		"little",
		"lambd",
		"mary had a litle lamb",
		"mary, had, a little, lamb",
		"marry {had} a little lamb",
		`mary "had" a little lamb`,
		`mary 'had' a little lamb`,
		`mary 'had' "a" little lambd`,
		`mary
		had
		a little
		lamb`,
		"the quick brown fox jumps over the lazy dog",
		"THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG",
		`
░░░░░░░░░░░░░░░░░░░░░░░▄▄
░░░░░░░░░░░░░░░░░░░░░░█░░█
░░░░░░░░░░░░░░░░░░░░░░█░░█
░░░░░░░░░░░░░░░░░░░░░█░░░█
░░░░░░░░░░░░░░░░░░░░█░░░░█
░░░░░░░░░░░███████▄▄█░░░░░██████▄
░░░░░░░░░░░▓▓▓▓▓▓█░░░░░░░░░░░░░░█
░░░░░░░░░░░▓▓▓▓▓▓█░░░░░░░░░░░░░░█
░░░░░░░░░░░▓▓▓▓▓▓█░░░░░░░░░░░░░░█
░░░░░░░░░░░▓▓▓▓▓▓█░░░░░░░░░░░░░░█
░░░░░░░░░░░▓▓▓▓▓▓█░░░░░░░░░░░░░░█
░░░░░░░░░░░▓▓▓▓▓▓█████░░░░░░░░░█
░░░░░░░░░░░██████▀░░░░▀▀██████▀
◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈
░█░░░█░█░▄▀░█▀▀░░░░▀█▀░█░█░█░▄▀▀░
░█░░░█░█▀░░░█▀░░▄▄░░█░░█▀█░█░░▀▄░
░█▄▄░█░█░▀▄░█▄▄░░░░░█░░█░█░█░▄▄▀░
◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈☻◈
		`,
		`⌘`,
	}
}

// Ngrams creates n-grams from the supplied array.
// It does not generate all n-grams (there is simply too much of them
// for bigger n). Number of generated ngrams is equal to length of supplied array
func Ngrams(ss []string, n int) [][]string {
	result := make([][]string, 0, len(ss))
	for i := range ss {
		ngram := make([]string, 0, n)
		for j := 0; j < n; j++ {
			wordIdx := (i + j) % len(ss)
			ngram = append(ngram, ss[wordIdx])
		}
		result = append(result, ngram)
	}
	return result
}

// NgramsFromSupplier generates ngrams using provided string supplier
func NgramsFromSupplier(supplier func() string, length, number int) [][]string {
	result := make([][]string, 0, number)
	for i := 0; i < number; i++ {
		ngram := make([]string, 0, number)
		for j := 0; j < length; j++ {
			ngram = append(ngram, supplier())
		}
		result = append(result, ngram)
	}
	return result
}
