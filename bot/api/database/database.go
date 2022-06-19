package database

import (
	"context"
	"os"

	firestore "cloud.google.com/go/firestore"
)

type Database struct {
	client *firestore.Client
	ctx    context.Context
}

func (d *Database) Connect() (err error) {
	d.ctx = context.TODO()
	d.client, err = firestore.NewClient(d.ctx, os.Getenv("PROJECT_ID"))

	return
}
