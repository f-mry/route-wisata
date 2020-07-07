package mysql

import (
	"database/sql"
	"krisna/pkg/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type WisataModel struct {
    DB *sql.DB
}

func (m *WisataModel) Get(page int) ([]*models.InfoWisata, error) {
    offset := page * 10

    stmt := "SELECT `id`, `nama`, `deskripsi` from `dummylocation` LIMIT ?,10"

    rows, err := m.DB.Query(stmt, offset)
    if err != nil {
        // log.Println(err)
        return nil, err
    }

    defer rows.Close()

    var infoWisata []*models.InfoWisata

    for rows.Next() {
        wisata := &models.InfoWisata{}
        err = rows.Scan(&wisata.Id, &wisata.Nama, &wisata.Deskripsi)
        if err != nil {
            log.Println(err)
        }
        infoWisata = append(infoWisata, wisata)
    }

    return infoWisata, nil
}
