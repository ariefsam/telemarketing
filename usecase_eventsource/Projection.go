package usecase_eventsource

import (
	"encoding/json"
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

func (p *Projection) FirestoreToRedis() {
	conn, err := gore.Dial("localhost:6379") //Connect to redis server at localhost:6379
	if err != nil {
		return
	}
	defer conn.Close()
	log.Println("Starting copy to redis...")
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
				// access the change.Doc returns the Document,
				// which contains Data() and DataTo(&p) methods.
				switch change.Kind {
				case firestore.DocumentAdded:
					// on added it returns the existing ones.

					log.Printf("%+v", change.Doc.Data())
					data := change.Doc.Data()
					jsonData, _ := json.Marshal(data)

					var event Event
					mapstructure.Decode(data, &event)
					gore.NewCommand("ZADD", "event", float64(event.Timestamp), string(jsonData)).Run(conn)
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

func (p *Projection) ProjectRedis() {
	conn, err := gore.Dial("localhost:6379") //Connect to redis server at localhost:6379
	if err != nil {
		return
	}
	defer conn.Close()

	temp, _ := gore.NewCommand("GET", "lastEventCursor").Run(conn)
	lastEventCursor, _ := temp.Int()

	for {
		temp, _ = gore.NewCommand("ZRANGE", "event", lastEventCursor, -1).Run(conn)
		list, _ := temp.Array()
		if len(list) > 0 {
			for _, val := range list {
				log.Println("Getting ccursor ", lastEventCursor)
				eventByte, _ := val.Bytes()
				var event Event
				json.Unmarshal(eventByte, &event)
				p.ProcessEvent(event)
				lastEventCursor++
				gore.NewCommand("SET", "lastEventCursor", lastEventCursor).Run(conn)
			}
		}
		time.Sleep(1 * time.Second)
	}

	// temp, _ := redis.Get("lastEventCursor")
	// lastEventCursor, _ := strconv.Atoi(temp)
	// for {
	// 	events, _ := redis.ZRange("event", lastEventCursor, 1)
	// 	lastEventCursor++
	// 	redis.Set("lastEventCursor", lastEventCursor)
	// }

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
