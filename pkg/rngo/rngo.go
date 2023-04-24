package rngo

import (
	"bytes"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
)

func Boop() string {
	return "Boop"
}

var numberSyllablesPik = [...]int{2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5}

// Rngo name generating thing.
type Rngo struct {
	preSyllables []Syllable
	midSyllables []Syllable
	surSyllables []Syllable
}

func (rngo Rngo) preSyllableSample() Syllable {
	return sampleSyllable(rngo.preSyllables)
}

func (rngo Rngo) midSyllableSample() Syllable {
	return sampleSyllable(rngo.midSyllables)
}

func (rngo Rngo) surSyllableSample() Syllable {
	return sampleSyllable(rngo.surSyllables)
}

// NameSyllables returns a slice of Syllables based upon the number passed in.
func (rngo Rngo) NameSyllables(number int) []Syllable {
	var nameSyllables []Syllable
	preSyllable := rngo.preSyllableSample()
	nameSyllables = append(nameSyllables, preSyllable)
	for i := 1; i < number-1; i++ {
		nameSyllables = append(nameSyllables, rngo.midSyllableSample())
	}
	if number > 1 {
		nameSyllables = append(nameSyllables, rngo.surSyllableSample())
	}
	return nameSyllables
}

// NameFromSyllables returns a Name from syllables.
func (rngo Rngo) NameFromSyllables(syllables []Syllable) string {
	var buffer bytes.Buffer
	for _, syllable := range syllables {
		buffer.WriteString(syllable.text)
	}
	return cases.Title(language.English, cases.Compact).String(buffer.String())
}

// Name random name.
func (rngo Rngo) Name(number int) string {
	return rngo.NameFromSyllables(rngo.NameSyllables(number))
}

// NameRnd generates random syllable length name
func (rngo Rngo) NameRnd() string {
	return rngo.Name(pickSyllablesCount())
}

// FirstLast generates a string with a first and last name from random syllables.
func (rngo Rngo) FirstLast() string {
	return rngo.NameRnd() + " " + rngo.NameRnd()
}

// New is a static factory method generating Rngos
func New(dialectMap map[string]Syllable) Rngo {
	var first []Syllable
	var mid []Syllable
	var last []Syllable

	for _, v := range dialectMap {
		if v.isPrefix {
			first = append(first, v)
		} else if v.isSuffix {
			last = append(last, v)
		} else {
			mid = append(mid, v)
		}
	}
	return Rngo{
		preSyllables: first,
		midSyllables: mid,
		surSyllables: last,
	}
}

func sampleSyllable(syllables []Syllable) Syllable {
	return syllables[rand.Intn(len(syllables))]
}

func pickSyllablesCount() int {
	return numberSyllablesPik[rand.Intn(len(numberSyllablesPik))]
}
