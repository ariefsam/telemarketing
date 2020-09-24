package repository

import (
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/ariefsam/telemarketing/entity"
	"github.com/keimoon/gore"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

type Customer struct{}

func (c *Customer) Save(customer entity.Customer) (err error) {
	err = saveToRedis("CUSTOMER", customer.ID, customer)
	err = saveToRedis("CUSTOMER_PHONE", customer.PhoneNumber, customer)
	// go c.saveToFirestore(customer)
	return
}

func saveToRedis(key string, id string, obj interface{}) (err error) {
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
		return
	}
	gore.NewCommand("HSET", key, id, jsonObj).Run(conn)
	return
}

func (c *Customer) saveToFirestore(customer entity.Customer) (err error) {
	ctx, client, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	defer client.Close()
	dataInsert := map[string]interface{}{}
	mapstructure.Decode(customer, &dataInsert)

	_, err = docRef.Collection("customer").Doc(customer.ID).Set(ctx, dataInsert)
	if err != nil {
		err = errors.Wrap(err, "Failed to save customer")
		return
	}
	return
}

func (c *Customer) Get(filter entity.FilterCustomer, limit int) (customers []entity.Customer, err error) {
	customers, err = c.getFromRedis(filter, limit)
	return
}

func getIndexIDFromRedis(key string) (data []string, err error) {
	tempIndex, err := gore.NewCommand("HKEYS", key).Run(conn)
	if err != nil {
		log.Println(err)
		return
	}
	tempArray, err := tempIndex.Array()
	if err != nil {
		log.Println(err)
		return
	}
	for _, val := range tempArray {
		tempID, err := val.String()
		if err != nil {
			log.Println(err)
			continue
		}

		data = append(data, tempID)
	}
	return
}

func getFromRedis(id string, key string) (data []byte, err error) {
	temp, err := gore.NewCommand("HGET", key, id).Run(conn)
	if err != nil {
		log.Println(err)
		return data, err
	}
	if temp == nil {
		return
	}
	data, err = temp.Bytes()
	if err != nil {
		if err.Error() == "nil value" {
			err = nil
			return
		}
		log.Println(err)
		return data, err
	}
	return
}

func (c *Customer) getFromRedisWithField(key string, id string) (customer *entity.Customer, err error) {
	if id != "" {
		temp, err := getFromRedis(id, key)
		if err != nil {
			return customer, err
		}
		err = json.Unmarshal(temp, &customer)
	}

	return
}

func (c *Customer) getFromRedis(filter entity.FilterCustomer, limit int) (customers []entity.Customer, err error) {
	key := "CUSTOMER"

	if filter.PhoneNumber != nil {
		key = "CUSTOMER_PHONE"
		customer, err := c.getFromRedisWithField(key, *filter.PhoneNumber)

		if err != nil {
			return customers, err
		} else if customer != nil {
			customers = append(customers, *customer)
		}

		return customers, nil
	}
	if filter.ID != nil {
		key = "CUSTOMER"
		customer, err := c.getFromRedisWithField(key, *filter.ID)

		if err != nil {
			return customers, err
		} else if customer != nil {
			customers = append(customers, *customer)
		}

		return customers, nil
	}
	ids, err := getIndexIDFromRedis(key)

	if err != nil {
		return
	}
	for _, id := range ids {

		customer, err := c.getFromRedisWithField(key, id)
		if err != nil {
			log.Println(err)
			continue
		}
		if filter.PhoneNumber != nil {
			if customer.PhoneNumber != *filter.PhoneNumber {
				continue
			}
		}
		if filter.TelemarketerID != nil {
			if customer.TelemarketerID != *filter.TelemarketerID {
				continue
			}
		}
		if filter.Status != nil {
			if customer.Status != *filter.Status {
				continue
			}
		}
		if filter.IsClosing != nil {
			if customer.IsClosing != *filter.IsClosing {
				continue
			}
		}

		if customer != nil {
			customers = append(customers, *customer)
		}

	}

	return
}

func (c *Customer) getFromFirestore(filter entity.FilterCustomer, limit int) (customers []entity.Customer, err error) {
	ctx, client, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	defer client.Close()

	var dsnap *firestore.DocumentSnapshot

	fWhere := []filterWhere{}
	if filter.ID != nil {
		dsnap, errs := docRef.Collection("customer").Doc(*filter.ID).Get(ctx)
		if errs != nil {
			err = errs
			return
		}
		m := dsnap.Data()
		var customer entity.Customer
		mapstructure.Decode(m, &customer)
		customers = []entity.Customer{customer}
		return
	}

	if filter.TelemarketerID != nil {
		fWhere = append(fWhere, filterWhere{
			path:     "TelemarketerID",
			operator: "==",
			value:    *filter.TelemarketerID,
		})
	}

	if filter.PhoneNumber != nil {
		fWhere = append(fWhere, filterWhere{
			path:     "PhoneNumber",
			operator: "==",
			value:    *filter.PhoneNumber,
		})
	}

	if filter.Status != nil {
		fWhere = append(fWhere, filterWhere{
			path:     "Status",
			operator: "==",
			value:    *filter.Status,
		})
	}

	var iter *firestore.DocumentIterator
	if len(fWhere) == 0 {
		iter = docRef.Collection("customer").Limit(limit).Documents(ctx)
	} else {
		var query firestore.Query
		for idx, val := range fWhere {
			if idx == 0 {
				query = docRef.Collection("customer").Where(val.path, val.operator, val.value)
			} else {
				query = query.Where(val.path, val.operator, val.value)
			}
		}
		iter = query.Limit(limit).Documents(ctx)
	}

	for {
		dsnap, err = iter.Next()

		if err == iterator.Done {
			err = nil
			break
		}
		if err != nil {
			return
		}
		var customer entity.Customer
		data := dsnap.Data()
		mapstructure.Decode(data, &customer)
		customers = append(customers, customer)
	}

	return
}

func (c *Customer) Delete(id string) (err error) {
	ctx, client, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	defer client.Close()
	_, err = docRef.Collection("customer").Doc(id).Delete(ctx)
	return
}
