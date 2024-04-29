package gooqu_test

import (
	"testing"

	gooqu "github.com/flat35hd99/gooqu"
)

func TestSelect(t *testing.T) {
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

	t.Run("SELECT * FROM `todos` WHERE `id` = 2 LIMIT 1;", func(t *testing.T) {
		query := gooqu.Where(gooqu.Record{"id": 2}).From("todos").Limit(1).ToSQL()

		expected := "SELECT * FROM `todos` WHERE `id` = 2 LIMIT 1;"
		if query != expected {
			t.Errorf("expected %s but got %s", expected, query)
		}
	})

	t.Run("SELECT `id` FROM `users` WHERE `id` = 1;", func(t *testing.T) {
		query := gooqu.Where(gooqu.Record{"id": 1}).Select(gooqu.Column{"id"}).From("users").ToSQL()

		expected := "SELECT `id` FROM `users` WHERE `id` = 1;"
		if query != expected {
			t.Errorf("expected %s but got %s", expected, query)
		}
	})

	t.Run("SELECT `id`, `name` FROM `users` WHERE `id` = 1;", func(t *testing.T) {
		query := gooqu.Where(gooqu.Record{"id": 1}).Select(gooqu.Column{"id"}, gooqu.Column{"name"}).From("users").ToSQL()

		expected := "SELECT `id`, `name` FROM `users` WHERE `id` = 1;"
		if query != expected {
			t.Errorf("expected %s but got %s", expected, query)
		}
	})

	t.Run("SELECT `id`, `name` FROM `users` WHERE `id` = 1;", func(t *testing.T) {
		query := gooqu.Select(gooqu.COUNT("id").As("count")).From("users").GroupBy("age").ToSQL()

		expected := "SELECT COUNT(`id`) AS `count` FROM `users` GROUP BY `age`;"
		if query != expected {
			t.Errorf("expected %s but got %s", expected, query)
		}
	})
}
