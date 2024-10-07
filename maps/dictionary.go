package dictionary

type Dictionary map[string]string

const (
	ErrorNotFound                = DictionaryError("key not found")
	ErrorKeyExists               = DictionaryError("key already exists")
	ErrorUpdateFailedKeyNotFound = DictionaryError("failed to update key because it does not exist")
	ErrorDeleteFailedKeyNotFound = DictionaryError("failed to delete key because it does not exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]

	if !ok {
		return "", ErrorNotFound
	}

	return val, nil
}

func (d Dictionary) Add(key, value string) (string, error) {
	_, ok := d[key]

	if ok {
		return "", ErrorKeyExists
	}

	d[key] = value

	return value, nil
}

func (d Dictionary) Update(key, value string) (string, error) {
	_, ok := d[key]
	if !ok {
		return "", ErrorUpdateFailedKeyNotFound
	}

	d[key] = value
	return "", nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
