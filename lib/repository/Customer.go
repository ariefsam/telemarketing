package repository

import (
	"cloud.google.com/go/firestore"
	"github.com/ariefsam/telemarketing/entity"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

type Customer struct{}

func (c *Customer) Save(customer entity.Customer) (err error) {
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	dataInsert := map[string]interface{}{}
	mapstructure.Decode(customer, &dataInsert)

	_, err = docRef.Collection("customer").Doc(customer.PhoneNumber).Set(ctx, dataInsert)
	if err != nil {
		err = errors.Wrap(err, "Failed to save customer")
		return
	}
	return
}

func (c *Customer) Get(filter entity.FilterCustomer, limit int) (customers []entity.Customer, err error) {
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}

	var dsnap *firestore.DocumentSnapshot
	if filter.PhoneNumber != nil {
		dsnap, err = docRef.Collection("customer").Doc(*filter.PhoneNumber).Get(ctx)
		if err != nil {
			return
		}
		var customer entity.Customer
		mapstructure.Decode(dsnap.Data(), &customer)
		customers = append(customers, customer)
		return
	}

	fWhere := []filterWhere{}

	if filter.TelemarketerEmail != nil {
		fWhere = append(fWhere, filterWhere{
			path:     "TelemarketerEmail",
			operator: "==",
			value:    *filter.TelemarketerEmail,
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
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	_, err = docRef.Collection("customer").Doc(id).Delete(ctx)
	return
}