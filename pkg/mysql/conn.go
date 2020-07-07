package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDBConnection() (*sql.DB, error){
    db, err := sql.Open("mysql", "root:@/data_ta")
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    if err = db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}
