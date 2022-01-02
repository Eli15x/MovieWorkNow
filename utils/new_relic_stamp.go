package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// ContextTransactionKey is the default key of the echo context which holds
// the handler new relic transaction
const ContextTransactionKey = "users-intervention-txn"

// StartSegmentWithDatastoreProduct get new relic transaction from echo context and starts datastore segment, then return it to be stopped
func StartSegmentWithDatastoreProduct(echoContext echo.Context, segmentName string, datastoreProduct newrelic.DatastoreProduct, operation string, collection string) newrelic.DatastoreSegment {
	txn := GetNewRelicTransaction(echoContext)
	if txn != nil {
		segment := newrelic.DatastoreSegment{
			Product:    datastoreProduct,
			Operation:  operation,
			Collection: collection,
		}
		segment.StartTime = txn.StartSegmentNow()
		txn.StartSegment(segmentName)
		return segment
	}
	return newrelic.DatastoreSegment{}
}

func GetNewRelicTransaction(echoContext echo.Context) *newrelic.Transaction {
	if echoContext == nil {
		log.Warnf("[NewRelic] FAILED on GetNewRelicTransaction echoContext: %v", echoContext)
		return nil
	}
	txn := echoContext.Get(ContextTransactionKey)
	transaction, ok := txn.(*newrelic.Transaction)
	if ok {
		return transaction
	}
	return nil
}
