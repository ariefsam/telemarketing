package entity

type Customer struct {
	ID                string
	Name              string
	PhoneNumber       string
	Status            string
	IsClosing         bool
	TelemarketerID    string
	TimestampCreated  int64
	TimestampUpdated  int64
	LastCallTimestamp int64
	ClosingTimestamp  int64
	DataSource        string
	BuyAmount         int
	CreatedBy         string
}
