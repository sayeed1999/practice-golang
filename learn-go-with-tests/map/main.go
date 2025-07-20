package dictionary

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrAlreadyPresent = errors.New("the exact meaning of the word is already present in dictionary")

// First make a custom type Dictionary from the map of string:string
type Dictionary map[string]string

// Caution: For map, don't use d *Dictionary! Why?

// Then make a custom Search method on it
func (d Dictionary) Search(key string) (string, error) {
	value, found := d[key]

	if !found {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(word, definition string) error {

	value, found := d[word]

	if found && value == definition {
		return ErrAlreadyPresent
	}

	d[word] = definition
	return nil
}
