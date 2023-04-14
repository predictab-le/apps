package main

import "context"

type Document struct{}

func ingestFromActivityPub(context.Context, <-chan Document) {}

func ingestFromDiscord(context.Context, <-chan Document) {}

func storeDocument(Document) error {
	return nil
}
