package maps

const (
	ErrNotFound         = DictionaryErr("could not find that word")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dict map[string]string

func (d Dict) Definition(lookup string) (string, error) {
	definition, ok := d[lookup]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dict) Add(word, definition string) error {
	_, err := d.Definition(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dict) Update(word, definition string) error {
	_, err := d.Definition(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dict) Delete(word string) {
	delete(d, word)
}
