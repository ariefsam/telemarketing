package repository

import (
	"cloud.google.com/go/firestore"
	"github.com/ariefsam/telemarketing/entity"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type Telemarketer struct{}

func (t *Telemarketer) Save(telemarketer entity.Telemarketer) (err error) {
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	dataInsert := map[string]interface{}{}
	mapstructure.Decode(telemarketer, &dataInsert)
	_, err = docRef.Collection("telemarketer").Doc(telemarketer.Email).Set(ctx, dataInsert)
	if err != nil {
		err = errors.Wrap(err, "Failed to save callLog")
		return
	}
	return
}

func (t *Telemarketer) Get(filter entity.FilterTelemarketer, limit int) (telemarketers []entity.Telemarketer, err error) {
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}

	var dsnap *firestore.DocumentSnapshot
	if filter.Email != "" {
		dsnap, err = docRef.Collection("telemarketer").Doc(filter.Email).Get(ctx)
		if err != nil {
			return
		}
		var telemarketer entity.Telemarketer
		mapstructure.Decode(dsnap.Data(), &telemarketer)
		telemarketers = append(telemarketers, telemarketer)
		return
	}
	return
}
