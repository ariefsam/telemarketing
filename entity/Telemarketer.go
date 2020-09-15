package entity

type Telemarketer struct {
	ID           string
	Name         string
	Email        string
	IsAdmin      bool
	CreatedBy    string
	WeeklyTarget struct {
		Call    int
		Closing int
	}
	Performance struct {
		Call    int
		Closing int
	}
}
