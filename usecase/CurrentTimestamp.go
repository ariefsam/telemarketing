package usecase

func (u *Usecase) CurrentTimestamp() (timestamp int64) {
	timestamp = u.Timer.CurrentTimestamp()
	return
}
