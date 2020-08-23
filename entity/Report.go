package entity

type Report struct {
	Timestamp      int64
	PeriodType     string
	Tertarik       int
	HubungiKembali int
	TidakTertarik  int
	TidakAktif     int
	TidakMenjawab  int
	TidakTerdaftar int
	Closing        int
}
