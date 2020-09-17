package entity

type Telemarketer struct {
	ID          string
	Name        string
	Email       string
	IsAdmin     bool
	CreatedBy   string
	Target      TargetPeriod
	Performance TargetPeriod
}

type TargetPeriod struct {
	Daily   TargetUnit
	Weekly  TargetUnit
	Monthly TargetUnit
}

type TargetUnit struct {
	Call      int
	Closing   int
	BuyAmount int
}
