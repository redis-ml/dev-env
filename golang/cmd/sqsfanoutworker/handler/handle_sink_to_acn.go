package handler

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type AcnPref struct {
	Owner    string `json:"owner,omitempty"`
	PrefType string `json:"pref_type,omitempty"`

	PriceAbove float64 `json:"price_above,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

func WriteToACN(ctx context.Context, pref *AcnPref) error {
	client := getDdbClient()

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

func transformPriceAbovePrefToRequest(ctx context.Context, pref *AcnPref, req *dynamodb.BatchWriteItemInput) error {
	req.RequestItems[tableNameForNotificationPreference] = []*dynamodb.WriteRequest{
		{
			PutRequest: &dynamodb.PutRequest{
				Item: map[string]*dynamodb.AttributeValue{
					// Partition Key
					colOwnerInNotifPref: val().SetS(pref.Owner),
					// Sort key
					colPrefTypeInNotifPref: val().SetS(pref.PrefType),
					// Indexed
					"#PriceAbove": val().SetN(fmt.Sprintf("%f", pref.PriceAbove)),
				},
			},
		},
	}
	return nil
}

func transformToOwnerID(ctx context.Context, id int) string {
	readable := fmt.Sprintf("0000-%06d", id)
	h := sha1.New()
	h.Write([]byte(readable))
	bs := h.Sum(nil)
	return fmt.Sprintf("rh:%x-%x-%s", bs[0:3], bs[3:4], readable)
}
