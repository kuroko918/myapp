package database

import "cloud.google.com/go/firestore"

type DbHandler interface {
	Collection(string) *firestore.CollectionRef
	Doc(string) *firestore.DocumentRef
}
