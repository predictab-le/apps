package main

import (
	"context"
	"log"

	"go.flipt.io/flipt/rpc/flipt"
	sdk "go.flipt.io/flipt/sdk/go"
)

func main() {
	var (
		ctx       context.Context
		features  = sdk.New(nil).Flipt()
		documents = make(chan Document)
	)

	go ingestFromActivityPub(ctx, documents)

	resp, err := features.GetFlag(ctx, &flipt.GetFlagRequest{
		Key: "ingestFromDiscord",
	})
	if err != nil {
		log.Fatal(err)
	}

	if resp.Enabled {
		go ingestFromDiscord(ctx, documents)
	}

	for doc := range documents {
		if err := storeDocument(doc); err != nil {
			log.Println("[ERROR]", err)
		}
	}
}
