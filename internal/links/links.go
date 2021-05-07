package links

import (
	"database/sql"
	"log"

	database "github.com/truongcongminh96/graphql-golang-example/internal/pkg/db/mysql"
	"github.com/truongcongminh96/graphql-golang-example/internal/users"
)

// Link #1
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

// Save #2
func (link Link) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	//#4
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}

	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	log.Print("Row inserted!")

	return id
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select id, title, address from Links")
	if err != nil {
		log.Fatal(err)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return links
}
