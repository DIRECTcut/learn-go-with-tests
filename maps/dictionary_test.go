package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known key", func(t *testing.T) {
		dictionary := Dictionary{"test": "testing"}
		key := "test"
		got, _ := dictionary.Search(key)
		want := "testing"

		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		dictionary := Dictionary{"test": "testing"}
		key := "aintgot"
		_, err := dictionary.Search(key)
		want := ErrorNotFound

		if err == nil {
			t.Fatal("expected an error")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		newKey := "test"
		newValue := "thisisatest"

		dictionary := Dictionary{}
		dictionary.Add(newKey, newValue)

		got, err := dictionary.Search(newKey)

		assertStrings(t, got, newValue)
		if err != nil {
			t.Errorf("expected no error but got one: %q", err)
		}
	})

	t.Run("existing key", func(t *testing.T) {
		existingKey := "key"
		existingValue := "exists"
		newValue := "newvalue"

		dictionary := Dictionary{existingKey: existingValue}
		_, addErr := dictionary.Add(existingKey, newValue)
		assertError(t, addErr, ErrorKeyExists)

		got, searchErr := dictionary.Search("key")
		assertError(t, searchErr, nil)
		assertStrings(t, got, existingValue)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing key", func(t *testing.T) {
		dictionary := Dictionary{"key": "old"}
		_, updateErr := dictionary.Update("key", "new")
		assertError(t, updateErr, nil)

		got, searchErr := dictionary.Search("key")
		assertStrings(t, got, "new")
		assertError(t, searchErr, nil)
	})

	t.Run("updating nonexistent key", func(t *testing.T) {
		dictionary := Dictionary{}
		_, updateErr := dictionary.Update("key", "notexists")

		assertError(t, updateErr, ErrorUpdateFailedKeyNotFound)

		_, searchErr := dictionary.Search("key")
		assertError(t, searchErr, ErrorNotFound)
	})
}

func TestDelete(t *testing.T) {
	key := "key"
	dictionary := Dictionary{key: "exists"}
	dictionary.Delete(key)

	_, searchErr := dictionary.Search(key)
	assertError(t, searchErr, ErrorNotFound)
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if want != got {
		t.Errorf("Got %q want %q", got, want)
	}
}

// TODO: copied from another place in code, extract to a module maybe?
// https://stackoverflow.com/a/45813698
func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if want == nil && got != nil {
		t.Errorf("wanted no error but got one: %q", got)
	}

	if got == nil && want != nil {
		t.Fatal("wanted an error but did not get one")
	}

	if got != want {
		t.Errorf("error message is incorrect, want %q, got %q", want, got)
	}
}
