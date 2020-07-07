package main

import (
	"fmt"
	"krisna/pkg/mysql"
	"log"
)

type Wisata struct {
    Id int32
    Nama string
    Deskripsi string
}

func main() {
    db, err :=mysql.OpenDBConnection()
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query("SELECT `id`, `nama`, `deskripsi` FROM `dummylocation`")
    if err != nil {
        log.Fatal(err)
    }

    var listWisata []*Wisata


    // fmt.Println(rows.Columns())
        for rows.Next() {
            s := &Wisata{}
            err = rows.Scan(&s.Id, &s.Nama, &s.Deskripsi)
            if err != nil {
                fmt.Println(err)
                continue
            }
            fmt.Println(s)
            listWisata = append(listWisata, s)
        }

    fmt.Printf("%v", listWisata)


}
