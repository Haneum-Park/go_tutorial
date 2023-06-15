package dict

import "errors"

type Dictionary map[string]string

var (
	errNotFound   = errors.New("not Found")
	errWordExists = errors.New("that word already exists")
	errCantUpdate = errors.New("can't update non-existing word")
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
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = def
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
