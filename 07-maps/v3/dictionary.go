package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add A map value is a pointer to a runtime.hmap structure.
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
