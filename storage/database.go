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

var db *sql.DB

func Startup(loc string) {
	var err error
	db, err = sql.Open("sqlite3", loc)
	if err != nil {
		panic(err)
	}

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

func GetArticleIDByURL(url string) int64 {
	stmt, _ := db.Prepare("SELECT id FROM articles WHERE url = ?;")
	result, _ := stmt.Query(url)

	var id int64
	for result.Next() {
		result.Scan(&id)
		// TODO: Cache this
	}
	return id
}

func InsertArticle(article *document.Article) int64 {
	url := SanitizeURL(article.Title)
	if GetArticleIDByURL(url) != 0 {
		i := 2 // Post 1 already exists and post 0 would only exist if everyone were nerds.
		for GetArticleIDByURL(fmt.Sprintf("%s-%d", url, i)) != 0 {
			i++
		}
		url = fmt.Sprintf("%s-%d", url, i)
	}

	// I'm dropping the errors here because I'm tired. If someone can show me a case where this would actually
	// error, I'll go back in and write error handling code.
	blob, _ := FromElement(make([]byte, 0), article)

	stmt, _ := db.Prepare("INSERT INTO articles (title, timestamp, url, content) VALUES(?, ?, ?, ?);")
	result, _ := stmt.Exec(article.Title, article.Timestamp.Unix(), url, blob)
	id, _ := result.LastInsertId()
	return id
}

func GetArticleByID(id int64) *document.Article {
	// TODO: Cache

	stmt, _ := db.Prepare("SELECT content FROM articles WHERE id = ?;")
	result, _ := stmt.Query(id)
	defer result.Close()

	if result.Next() {
		var blob []byte
		result.Scan(&blob)
		article, _, _ := ToElement(blob)
		return article.(*document.Article)
	}
	return nil
}

func GetLatestArticles() document.Content {
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
