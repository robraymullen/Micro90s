package db

import (
	"fmt"

	"github.com/hashicorp/go-memdb"
)

type Guest struct {
	Id      int
	Name    string
	Email   string
	Comment string
	Date    int
}

func RunDB() []*Guest {
	fmt.Printf("running main function\n")

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"guest": &memdb.TableSchema{
				Name: "guest",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Id"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"email": &memdb.IndexSchema{
						Name:    "email",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"comment": &memdb.IndexSchema{
						Name:    "comment",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Comment"},
					},
					"date": &memdb.IndexSchema{
						Name:    "date",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Date"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// Create a write transaction
	txn := db.Txn(true)

	// Insert some people
	guests := []*Guest{
		&Guest{1, "Anna", "email@email.com", "a com1", 123456789},
		&Guest{2, "Bob", "email1@email.com", "a com2", 2456789798},
		&Guest{3, "Dorothy", "email2@email.com", "a com3", 789456123456},
		&Guest{4, "Daniel", "email3@email.com", "a com4", 123465468},
	}

	// insert to db
	for _, p := range guests {
		if err := txn.Insert("guest", p); err != nil {
			panic(err)
		}
	}

	// Commit the transaction
	txn.Commit()

	// Create read-only transaction
	txn = db.Txn(false)
	defer txn.Abort()

	// query all the people
	it, err := txn.Get("guest", "name")
	if err != nil {
		panic(err)
	}

	fmt.Println("All the guests:")
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Guest)
		fmt.Printf("  %s\n", p.Name)
	}

	return guests
}
