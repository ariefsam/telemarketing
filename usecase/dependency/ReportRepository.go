package dependency

type ReportRepository interface {
	IncrementClosing(telemarketerID, periodType, perioString string) error
}
