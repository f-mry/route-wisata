package models

type InfoWisata struct {
    Id int
    Nama string
    Deskripsi string
}

type InfoHotel struct {
    Id int
    Nama string
    
}

type Route struct {
	Response []struct {
		HotelArrive   []string        `json:"hotelArrive"`
		HotelInfo     []interface{}   `json:"hotelInfo"`
		HotelLeaving  []string        `json:"hotelLeaving"`
		HotelLocation []float64       `json:"hotelLocation"`
		Location      [][]float64     `json:"location"`
		NodeIndex     [][]interface{} `json:"nodeIndex"`
		VisitTime     [][]string      `json:"visitTime"`
	} `json:"response"`
}
