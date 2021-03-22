package jadwal_shalat

type JadwalShalat struct {
	Status string `json:"status"`
	Query  Query  `json:"query"`
	Jadwal Jadwal `json:"jadwal"`
}

type Query struct {
	Format  string `json:"format"`
	Kota    string `json:"kota"`
	Tanggal string `json:"tanggal"`
}

type Jadwal struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Ashar   string `json:"ashar"`
	Dhuha   string `json:"dhuha"`
	Dzuhur  string `json:"dzuhur"`
	Imsak   string `json:"imsak"`
	Isya    string `json:"isya"`
	Maghrib string `json:"maghrib"`
	Subuh   string `json:"subuh"`
	Tanggal string `json:"tanggal"`
	Terbit  string `json:"terbit"`
}
