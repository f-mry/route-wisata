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
	Day1 struct {
		HotelArrive   []string        `json:"hotelArrive"`
		HotelInfo     []interface{}   `json:"hotelInfo"`
		HotelLeaving  []string        `json:"hotelLeaving"`
		HotelLocation []float64       `json:"hotelLocation"`
		Location      [][]interface{}     `json:"location"`
		NodeIndex     [][]interface{} `json:"nodeIndex"`
		VisitTime     [][]string      `json:"visitTime"`
	} `json:"day1"`
	Day2 struct {
		HotelArrive   []string        `json:"hotelArrive"`
		HotelInfo     []interface{}   `json:"hotelInfo"`
		HotelLeaving  []string        `json:"hotelLeaving"`
		HotelLocation []float64       `json:"hotelLocation"`
		Location      [][]interface{}     `json:"location"`
		NodeIndex     [][]interface{} `json:"nodeIndex"`
		VisitTime     [][]string      `json:"visitTime"`
	} `json:"day2"`
	Day3 []interface{} `json:"day3"`
}
