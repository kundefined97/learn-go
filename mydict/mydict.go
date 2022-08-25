package mydict

import (
	"errors"
)

type Dictionary map[string]string

var (
	errNotFound   = errors.New("not found")
	errCantUpdate = errors.New("cant update")
	errWordExists = errors.New("aleary word")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
