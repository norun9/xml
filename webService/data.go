package webService

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init(){
	var err error
	Db, err = sql.Open("postgres", "user=tomoyaueno dbname=gwp password=gwp sslmode=unable")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	
}