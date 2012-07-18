package storage

import (
	"io/ioutil"
	"os"
	"sync"
	"testing"
)

func init() {
	f, _ := ioutil.TempFile(os.TempDir(), "gocms_db_test")
	Startup(f.Name())
}

func TestDatabase(t *testing.T) {
	t.Parallel()

	if want, got := "Welcome to go.cms!", GetArticleByID(1).Title; want != got {
		t.Error("ID->Title is incorrect: Want(", want, ") but Got(", got, ")")
	}

	got, err := GetArticleIDByURL("welcome-to-go-cms")
	if err != nil {
		t.Error(err)
	}
	if 1 != got {
		t.Error("URL->Article ID is incorrect: Want(", 1, ") but Got(", got, ")")
	}
}

func BenchmarkExtremeConcurrency(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			GetLatestArticles()
			wg.Done()
		}()
	}
	wg.Wait()
}
