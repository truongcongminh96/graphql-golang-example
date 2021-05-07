package links

import (
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
