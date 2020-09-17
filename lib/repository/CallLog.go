package repository

import (
	"cloud.google.com/go/firestore"
	"github.com/ariefsam/telemarketing/entity"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
	"gopkg.in/webnice/b64.v1"
)

type CallLog struct{}

func (c *CallLog) Save(callLog entity.CallLog) (err error) {
	ctx, client, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	defer client.Close()
	dataInsert := map[string]interface{}{}
	mapstructure.Decode(callLog, &dataInsert)
	id := b64.NewInt64(callLog.Timestamp).String()
	_, err = docRef.Collection("call_log").Doc(id).Set(ctx, dataInsert)
	if err != nil {
		err = errors.Wrap(err, "Failed to save callLog")
		return
	}
	return
}

func (c *CallLog) Get(filter entity.FilterCallLog, limit int) (callLogs []entity.CallLog, err error) {
	ctx, client, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	defer client.Close()

	var dsnap *firestore.DocumentSnapshot
	fWhere := []filterWhere{}

	if filter.TelemarketerID != nil {
		fWhere = append(fWhere, filterWhere{
			path:     "TelemarketerID",
			operator: "==",
			value:    &filter.TelemarketerID,
		})
	}

	if filter.PhoneNumber != nil {
		fWhere = append(fWhere, filterWhere{
			path:     "PhoneNumber",
			operator: "==",
			value:    &filter.PhoneNumber,
		})
	}

	if filter.Status != nil {
		fWhere = append(fWhere, filterWhere{
			path:     "Status",
			operator: "==",
			value:    &filter.Status,
		})
	}

	if filter.TimestampStart > 0 {
		fWhere = append(fWhere, filterWhere{
			path:     "Timestamp",
			operator: ">=",
			value:    filter.TimestampStart,
		})
	}

	if filter.TimestampEnd > 0 {
		fWhere = append(fWhere, filterWhere{
			path:     "Timestamp",
			operator: "<",
			value:    filter.TimestampEnd,
		})
	}

	var iter *firestore.DocumentIterator
	if len(fWhere) == 0 {
		iter = docRef.Collection("call_log").Limit(limit).Documents(ctx)
	} else {
		var query firestore.Query
		for idx, val := range fWhere {
			if idx == 0 {
				query = docRef.Collection("call_log").Where(val.path, val.operator, val.value)
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
		var callLog entity.CallLog
		data := dsnap.Data()
		mapstructure.Decode(data, &callLog)
		callLogs = append(callLogs, callLog)
	}

	return
}

func (c *CallLog) Delete(timestamp int64) (err error) {
	id := b64.NewInt64(timestamp).String()
	ctx, client, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	defer client.Close()
	_, err = docRef.Collection("callLog").Doc(id).Delete(ctx)
	return
}
