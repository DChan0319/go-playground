package dictionary

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound  = DictionaryErr("no word found")
	ErrKeyExists = DictionaryErr("key value already exists")
	ErrKeyNotFound = DictionaryErr("word does not exist for an update to be made")
)

// Wrapper func to wrap Errors as an Dictionary Error 
func (e DictionaryErr) Error() string {
	return string(e)
}

// Search func to look for a key in the map & returns the value.
func (d Dictionary) Search(key string) (string, error) {
	var result string = ""
	result, found := d[key]

	if !found {
		return "", ErrNotFound
	}
	return result, nil
}

// Add func to add k,v pair to map(dictionary)
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

// Update func updates a value for a specific key
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrKeyNotFound
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
