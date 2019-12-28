package dictionary

import "testing" 

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		actual, _ := dictionary.Search("test")
		expected := "this is just a test"
	
		assertString(t, actual, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		expected := "no word found"

		assertError(t, err, ErrNotFound)
		assertString(t, err.Error(), expected)
	})
}

func TestAdd(t *testing.T) {
	
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
			word := "test"
			definition := "this is just a test"
			dictionary := Dictionary{word: definition}
			err := dictionary.Add(word, "new test")

			assertError(t, err, ErrKeyExists)
			assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		var word string = "test"
		var definition string = "this is a test"
		var newDefinition string = "new definition"
	
		dictionary := Dictionary{word: definition}
		err := dictionary.Update(word, newDefinition)
	
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		var word string = "test"
		var definition string = "this is a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrKeyNotFound)
	})
}

func TestDelete(t *testing.T) {
	var word string = "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertString(t *testing.T, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("actual: %q, expected: %q", actual, expected)
	}
}

func assertError(t *testing.T, actual, expected error) {
	t.Helper()

	if actual != expected {
		t.Errorf("actual error: %q, expected error: %q", actual, expected)
	}
}

func assertDefinition(t *testing.T, d Dictionary, key, value string) {
	t.Helper()

	actual, err := d.Search(key)
	if err != nil {
		t.Fatal("Should have searched for added word:", err)
	}

	assertString(t, actual, value)
}
