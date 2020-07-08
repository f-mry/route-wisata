package mysql

import (
	"database/sql"
	"krisna/pkg/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type HotelModel struct {
    DB *sql.DB
}

func (m *HotelModel) Get(page int) ([]*models.InfoHotel, error) {
    offset := page * 10

    stmt := "SELECT `id`, `nama` FROM `dummyhotel`"

    rows, err := m.DB.Query(stmt, offset)
    if err != nil {
        // log.Println(err)
        return nil, err
    }

    defer rows.Close()

    var infoHotel []*models.InfoHotel

    for rows.Next() {
        hotel := &models.InfoHotel{}
        err = rows.Scan(&hotel.Id, &hotel.Nama)
        if err != nil {
            log.Println(err)
        }
        infoHotel = append(infoHotel, wisata)
    }

    return infoHotel, nil
}
