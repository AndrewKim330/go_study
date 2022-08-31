package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// 220831
var errWordExists = errors.New("That word already exists")

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	//switch err {
	//case errNotFound:
	//	d[word] = def
	//case nil:
	//	return errWordExists
	//}
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

var errCantUpdate = errors.New("Can't update non-existing word")

// can define variable as the way below
//var (
//	errNotFound = errors.New("Not Found")
//	errWordExists = errors.New("That word already exists")
//	errCantUpdate = errors.New("Can't update non-existing word")
//)

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	// **code challange for error handling
	delete(d, word)
}
