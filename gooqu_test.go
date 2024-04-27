package gooqu_test

import (
	"testing"

	gooqu "github.com/flat35hd99/gooqu"
)

func TestQueryBuilder(t *testing.T) {
	t.Run("SELECT * FROM `users` WHERE `id` = 1;", func(t *testing.T) {
		query := gooqu.Where(gooqu.Record{"id": 1}).From("users").ToSQL()

		expected := "SELECT * FROM `users` WHERE `id` = 1;"
		if query != expected {
			t.Errorf("expected %s but got %s", expected, query)
		}
	})

	t.Run("SELECT * FROM `users` WHERE `id` = 2;", func(t *testing.T) {
		query := gooqu.Where(gooqu.Record{"id": 2}).From("users").ToSQL()

		expected := "SELECT * FROM `users` WHERE `id` = 2;"
		if query != expected {
			t.Errorf("expected %s but got %s", expected, query)
		}
	})

	t.Run("SELECT * FROM `todos` WHERE `id` = 2;", func(t *testing.T) {
		query := gooqu.Where(gooqu.Record{"id": 2}).From("todos").ToSQL()

		expected := "SELECT * FROM `todos` WHERE `id` = 2;"
		if query != expected {
			t.Errorf("expected %s but got %s", expected, query)
		}
	})
}
