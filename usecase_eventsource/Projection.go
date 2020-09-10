package usecase_eventsource

import (
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/ioc"
	"github.com/keimoon/gore"

	"github.com/mitchellh/mapstructure"

	"google.golang.org/api/iterator"
)

type Projection struct {
}

var NewEvent bool

func (p *Projection) FirestoreProjection() {
	conn, err := gore.Dial("localhost:6379") //Connect to redis server at localhost:6379
	if err != nil {
		return
	}
	defer conn.Close()
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	limit := 10000

	for {
		lastTimeTemp, _ := gore.NewCommand("Get", "event-from-server-last-time").Run(conn)
		lastTime, _ := lastTimeTemp.Int()

		log.Println("Get event from lasttime:", lastTime)
		query := docRef.Collection("Event").Where("Timestamp", ">", lastTime).OrderBy("Timestamp", firestore.Asc)
		iter := query.Limit(limit).Snapshots(ctx)

		for {
			dsnap, err := iter.Next()

			if err == iterator.Done {
				err = nil
				break
			}
			if err != nil {
				return
			}

			for _, change := range dsnap.Changes {
				switch change.Kind {
				case firestore.DocumentAdded:
					log.Printf("%+v", change.Doc.Data())
					data := change.Doc.Data()

					var event Event
					mapstructure.Decode(data, &event)
					var p Projection
					p.ProcessEvent(event)
					gore.NewCommand("SET", "event-from-server-last-time", fmt.Sprint(event.Timestamp)).Run(conn)

				case firestore.DocumentModified:
					// [...]
				case firestore.DocumentRemoved:
					// [...]
				}
			}

			time.Sleep(1 * time.Second)
		}
	}
}

func (p *Projection) ProcessEvent(event Event) {
	usecase := ioc.Usecase()
	if event.Name == "SaveCustomer" {
		log.Printf("Processing %+v", event)
		var data entity.Customer
		mapstructure.Decode(event.Data[0], &data)
		usecase.SaveCustomer(data)
	}

}
