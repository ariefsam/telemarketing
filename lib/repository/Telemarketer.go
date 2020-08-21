package repository

import (
	"cloud.google.com/go/firestore"
	"github.com/ariefsam/telemarketing/entity"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

type Telemarketer struct{}

func (t *Telemarketer) Save(telemarketer entity.Telemarketer) (err error) {
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	if telemarketer.Email == "" {
		err = errors.New("Cannot save telemarketer because email is empty, " + telemarketer.Name)
		return
	}
	dataInsert := map[string]interface{}{}
	mapstructure.Decode(telemarketer, &dataInsert)
	_, err = docRef.Collection("telemarketer").Doc(telemarketer.Email).Set(ctx, dataInsert)
	if err != nil {
		err = errors.Wrap(err, "Failed to save telemarketer email: "+telemarketer.Email)
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
	iter := docRef.Collection("telemarketer").Limit(limit).Documents(ctx)
	for {
		dsnap, err = iter.Next()

		if err == iterator.Done {
			err = nil
			break
		}
		if err != nil {
			return
		}
		var telemarketer entity.Telemarketer
		data := dsnap.Data()
		mapstructure.Decode(data, &telemarketer)
		telemarketers = append(telemarketers, telemarketer)
	}
	return
}

func (c *Telemarketer) Delete(email string) (err error) {
	id := email
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	_, err = docRef.Collection("telemarketer").Doc(id).Delete(ctx)
	return
}
