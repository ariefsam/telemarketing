package main

import (
	"github.com/ariefsam/telemarketing/ioc_eventsource"
	"github.com/ariefsam/telemarketing/restapi"
	"github.com/ariefsam/telemarketing/usecase_eventsource"
)

func main() {
	var p usecase_eventsource.Projection
	go p.FirestoreProjection()

	usecase := ioc_eventsource.Usecase()
	restapi.Serve(&usecase)
}
