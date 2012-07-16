package storage

import "testing"

func TestDatabase(t *testing.T) {
	t.Parallel()
	Startup(":memory:")

	if want, got := "Welcome to go.cms!", GetArticleByID(1).Title; want != got {
		t.Error("ID->Title is incorrect: Want(", want, ") but Got(", got, ")")
	}

	if want, got := int64(1), GetArticleIDByURL("welcome-to-go-cms"); want != got {
		t.Error("URL->Article ID is incorrect: Want(", want, ") but Got(", got, ")")
	}

	// For testing, free up some resources as we won't be using the database apart from this one test.
	db.Close()
}
