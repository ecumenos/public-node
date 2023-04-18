package database

import "github.com/boltdb/bolt"

//go:generate mockery --name=Connection
var _ Connection = (*connection)(nil)

// Connection is interface for database connection. It is wrapper for bolt DB driver.
type Connection interface{}

// NewConnection is constructor for Connection interface.
func NewConnection() (Connection, error) {
	driver, err := bolt.Open("clientsledger/db/clients.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	return &connection{driver: driver}, nil
}

type connection struct {
	driver *bolt.DB
}
