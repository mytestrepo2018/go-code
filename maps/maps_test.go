package maps

import (
	"testing"
)

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("want %q but got %q", want, got)
	}

}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if want != got {
		t.Errorf("want %q but got %q", want, got)
	}

}

func TestDictionary(t *testing.T) {
	d := Dict{"red": "a colour"}

	t.Run("word exists", func(t *testing.T) {

		want := "a colour"
		got, _ := d.Definition("red")

		assertStrings(t, got, want)
	})

	t.Run("word does not exist", func(t *testing.T) {

		_, got := d.Definition("yellow")

		assertError(t, got, ErrNotFound)
	})

}
func TestAdd(t *testing.T) {
	dictionary := Dict{}
	word := "test"
	definition := "this is just a test"
	dictionary.Add(word, definition)

	assertDefinition(t, dictionary, word, definition)
}

func assertDefinition(t *testing.T, dictionary Dict, word, definition string) {
	t.Helper()

	want := definition
	got, err := dictionary.Definition(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAddAlreadyExists(t *testing.T) {
	dictionary := Dict{}
	word := "test"
	definition := "this is just a test"
	err := dictionary.Add(word, definition)
	if err != nil {
		t.Fatalf("word %q should be added ok: %q", word, err)
	}

	assertDefinition(t, dictionary, word, definition)

	err = dictionary.Add(word, definition)
	if err == nil {
		t.Fatalf("word %q should not be added a second time: %q", word, err)
	}
}

func TestUpdate(t *testing.T) {
	dictionary := Dict{"test": "starting definition"}
	word := "test"
	definition := "this is the updated defintion"

	err := dictionary.Update(word, definition)
	assertError(t, err, nil)
	assertDefinition(t, dictionary, word, definition)
}

func TestUpdateNotExist(t *testing.T) {
	dictionary := Dict{}
	word := "test"
	definition := "this is the updated defintion"

	err := dictionary.Update(word, definition)
	assertError(t, err, ErrWordDoesNotExist)
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dict{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Definition(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}
