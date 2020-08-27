package usecase_eventsource

import "github.com/teris-io/shortid"

func generateID() string {
	id, _ := shortid.Generate()
	return id
}
