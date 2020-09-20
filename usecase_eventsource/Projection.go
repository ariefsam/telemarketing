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
	gore.NewCommand("SELECT", 15).Run(conn)
	defer conn.Close()
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}

	for {
		lastTimeTemp, _ := gore.NewCommand("Get", "event-from-server-last-time").Run(conn)
		lastTime, _ := lastTimeTemp.Int()

		log.Println("Get event from lasttime:", lastTime)
		query := docRef.Collection("Event").Where("Timestamp", ">", lastTime).OrderBy("Timestamp", firestore.Asc)
		iter := query.Snapshots(ctx)

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
					// log.Printf("%+v", change.Doc.Data())
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

	// log.Printf("Processing %+v", event)
	if event.Name == "CreateCustomer" {
		var data entity.Customer
		mapstructure.Decode(event.Data, &data)
		err := usecase.CreateCustomer(data)
		if err != nil {
			log.Println(err)
		}
	}

	if event.Name == "ClosingCustomer" {
		var data entity.Customer
		mapstructure.Decode(event.Data, &data)
		err := usecase.ClosingCustomer(data)
		if err != nil {
			log.Println(err)
		}
	}

	if event.Name == "CreateTelemarketer" {
		var data entity.Telemarketer
		mapstructure.Decode(event.Data, &data)
		err := usecase.CreateTelemarketer(data)
		if err != nil {
			log.Println(err)
		}
	}
	if event.Name == "AssignCustomer" {
		var data []string
		mapstructure.Decode(event.Data, &data)
		if len(data) != 2 {
			log.Println("Bad assign customer")
		} else {
			err := usecase.AssignCustomer(data[0], data[1])
			if err != nil {
				log.Println(err)
			}
		}
	}

	if event.Name == "SaveTelemarketer" {
		var data entity.Telemarketer
		mapstructure.Decode(event.Data, &data)
		err := usecase.SaveTelemarketer(data)
		if err != nil {
			log.Println(err)
		}
	}

	if event.Name == "Call" {
		data := struct {
			Telemarketer entity.Telemarketer
			Customer     entity.Customer
			Status       string
			Timestamp    int64
		}{}
		mapstructure.Decode(event.Data, &data)
		err := usecase.Call(data.Telemarketer, data.Customer, data.Status, data.Timestamp)
		if err != nil {
			log.Println(err)
		}
	}

}
