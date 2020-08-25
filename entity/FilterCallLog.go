package entity

type FilterCallLog struct {
	PhoneNumber    *string
	Status         *string
	TelemarketerID *string
	TimestampStart int64
	TimestampEnd   int64
}
