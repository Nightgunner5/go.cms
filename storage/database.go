package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"llamaslayers.net/go.cms/document"
	"log"
	"time"
	"unicode"
)

const NUM_DB_CONNECTIONS = 10

var dbCache = make(chan *sql.DB, NUM_DB_CONNECTIONS)

func acquire() *sql.DB {
	return <-dbCache
}
func recycle(db *sql.DB) {
	dbCache <- db
}

func Startup(loc string) {
	for i := 0; i < NUM_DB_CONNECTIONS; i++ {
		db, err := sql.Open("sqlite3", loc)
		if err != nil {
			panic(err)
		}
		dbCache <- db
	}

	db := acquire()
	defer recycle(db)
	version, err := db.Query("SELECT revision FROM version WHERE module='database';")
	if err != nil { // No database? No problem!
		if err.Error() != "no such table: version" {
			panic(err)
		}

		_, err = db.Exec("CREATE TABLE version ( module VARCHAR(64) PRIMARY KEY, revision UNSIGNED BIGINT NOT NULL ); INSERT INTO version VALUES( 'database', 0 );")
		if err != nil {
			panic(err)
		}

		version, err = db.Query("SELECT revision FROM version WHERE module='database';")
		if err != nil {
			panic(err)
		}
	}
	var versionNumber uint64
	for version.Next() {
		version.Scan(&versionNumber)
	}
	version.Close()

	log.Print("Checking database version...")
	switch versionNumber {
	case 0:
		_, err = db.Exec(`
CREATE TABLE articles (
	id        INTEGER PRIMARY KEY ASC AUTOINCREMENT,
	title     TEXT,
	timestamp DATETIME,
	url       TEXT,
	content   BLOB
);
CREATE INDEX timestamp ON articles (timestamp DESC);
CREATE UNIQUE INDEX url ON articles (url);
`)
		InsertArticle(&document.Article{
			"Welcome to go.cms!",
			time.Now(),
			document.Content{
				&document.Paragraph{
					document.Content{
						&document.LeafElement{"This is an example post. You can edit or delete it."},
					},
				},
			},
		})
		if err != nil {
			panic(err)
		}
		updateDBVersion("database", 1)
		fallthrough
	default:
		log.Print("Database is up to date.")
	}
}

func updateDBVersion(module string, version uint64) {
	db := acquire()
	defer recycle(db)
	stmt, _ := db.Prepare("REPLACE INTO version VALUES(?, ?);")
	stmt.Exec(module, version)
	log.Printf("Updated %s to version %d.", module, version)
}

func SanitizeURL(text string) string {
	prevSpace := true
	url := make([]rune, 0, len(text))
	var char rune
	for _, char = range text {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			if !prevSpace {
				url = append(url, '-')
				prevSpace = true
			}
			continue
		}
		prevSpace = false
		url = append(url, unicode.ToLower(char))
	}
	if prevSpace && len(url) > 0 {
		url = url[:len(url)-1]
	}
	return string(url)
}

func GetArticleIDByURL(url string) (int64, error) {
	db := acquire()
	defer recycle(db)

	result, err := db.Query("SELECT id FROM articles WHERE url = ?;", url)
	if err != nil {
		return 0, err
	}
	defer result.Close()

	var id int64
	for result.Next() {
		result.Scan(&id)
		// TODO: Cache this
	}
	return id, nil
}

func InsertArticle(article *document.Article) (int64, error) {
	db := acquire()
	defer recycle(db)

	url := SanitizeURL(article.Title)
	if id, _ := GetArticleIDByURL(url); id != 0 {
		i := 1
		for id != 0 {
			i++
			id, _ = GetArticleIDByURL(fmt.Sprintf("%s-%d", url, i))
		}
		url = fmt.Sprintf("%s-%d", url, i)
	}

	blob, err := FromElement(make([]byte, 0), article)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec("INSERT INTO articles (title, timestamp, url, content) VALUES(?, ?, ?, ?);", article.Title, article.Timestamp.Unix(), url, blob)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetArticleByID(id int64) *document.Article {
	db := acquire()
	defer recycle(db)

	// TODO: Cache

	articles, _ := db.Query("SELECT content FROM articles WHERE id = ?;", id)

	if articles.Next() {
		var blob []byte
		articles.Scan(&blob)
		article, _, _ := ToElement(blob)
		return article.(*document.Article)
	}
	return nil
}

func GetLatestArticles() document.Content {
	db := acquire()
	defer recycle(db)

	// TODO: Cache

	content := make(document.Content, 0, 5)

	ids, _ := db.Query("SELECT id FROM articles ORDER BY timestamp DESC LIMIT 5;")
	defer ids.Close()

	for ids.Next() {
		var id int64
		ids.Scan(&id)
		content = append(content, GetArticleByID(id))
	}

	return content
}
