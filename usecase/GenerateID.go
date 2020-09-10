package usecase

func (u *Usecase) GenerateID() (id string) {
	id = u.IDGenerator.Generate()
	return
}
