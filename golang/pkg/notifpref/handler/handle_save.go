package handler

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/redisliu/dev-env/golang/models/notifpref"
	. "github.com/redisliu/dev-env/golang/util/awsutil"
)

type NotifPref struct {
	Owner        string `json:"owner,omitempty"`
	InstrumnetID string `json:"instrument_id,omitempty"`
	PrefType     string `json:"pref_type,omitempty"`

	Threshold float64 `json:"threshold,omitempty"`
	UpdatedAt *int64  `json:"updated_at,omitempty"`
}

func HandleRequest(ctx context.Context, fakeID uint64) (string, error) {
	pref := NewNotifPref(ctx, fakeID)
	if err := SaveNotifPref(ctx, pref); err != nil {
		log.Printf("failed to SaveNotifPref: %+v\n", err)
		return fmt.Sprintf("%v", err), err
	}
	return "", nil
}

func NewNotifPref(ctx context.Context, fakeID uint64) *NotifPref {
	ownerID := transformToOwnerID(ctx, int(fakeID))
	instrumentID := "cffcbbe0-7b7b-45a9-a2e9-02058a9fa16c"
	return &NotifPref{
		Owner:        ownerID,
		InstrumnetID: instrumentID,
		PrefType:     "custom_price_above",
		Threshold:    100,
	}
}

func SaveNotifPref(ctx context.Context, pref *NotifPref) error {
	client := GetDdbClient()

	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{},
	}

	if err := transformPriceAbovePrefToRequest(ctx, pref, input); err != nil {
		log.Printf("failed to transformPriceAbovePrefToRequest: %+v\n", err)
		return err
	}

	output, err := client.BatchWriteItemWithContext(ctx, input)
	if err != nil {
		log.Printf("failed to BatchWriteItem: %+v\n", err)
		return nil
	}
	// transformPriceBelowPrefToRequest(ctx, pref, req)
	log.Printf("BatchWriteItem: %+v\n", output)
	return nil
}

func transformPriceAbovePrefToRequest(ctx context.Context, pref *NotifPref, req *dynamodb.BatchWriteItemInput) error {
	nowTS := time.Now().Unix()

	req.RequestItems[notifpref.TableName] = []*dynamodb.WriteRequest{
		{
			PutRequest: &dynamodb.PutRequest{
				Item: map[string]*dynamodb.AttributeValue{
					// Partition Key
					notifpref.ColOwner: DdbVal().SetS(pref.Owner),
					// Sort key
					notifpref.ColPrefType: DdbVal().SetS(pref.PrefType),
					// Indexed
					notifpref.ColThld:      DdbVal().SetN(fmt.Sprintf("%f", pref.Threshold)),
					notifpref.ColUpdatedAt: DdbVal().SetN(fmt.Sprintf("%d", nowTS)),
					notifpref.ColInstrPref: DdbVal().SetS(toInstrPref(ctx, pref)),
				},
			},
		},
	}
	return nil
}

func toInstrPref(ctx context.Context, pref *NotifPref) string {
	return fmt.Sprintf("%s#%s", pref.InstrumnetID, pref.PrefType)
}

func transformToOwnerID(ctx context.Context, id int) string {
	readable := fmt.Sprintf("0000-%06d", id)
	h := sha1.New()
	h.Write([]byte(readable))
	bs := h.Sum(nil)
	return fmt.Sprintf("rh:%x-%x-%s", bs[0:3], bs[3:4], readable)
}
