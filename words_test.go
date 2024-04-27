package gooqu

import "testing"

func TestWords(t *testing.T) {
	t.Run("hoge", func(t *testing.T) {
		selectKeyword := newWord("SELECT", false)
		columns := newWord("*", false)
		whereKeyword := newWord("WHERE", false)
		selectKeyword.n(columns).n(whereKeyword)

		expected := "SELECT * WHERE"
		if selectKeyword.String() != expected {
			t.Errorf("expected %s but got %s", expected, selectKeyword.String())
		}
	})

	t.Run("fuga", func(t *testing.T) {
		root := newWord("", false)
		root.n(newWord("SELECT", false)).n(newWord("*", false)).n(newWord("WHERE", false))

		expected := "SELECT * WHERE"
		if root.String() != expected {
			t.Errorf("expected %s but got %s", expected, root.String())
		}
	})
}
