package config

import "context"

// DatastoreDAO :
type DatabaseDao struct {
	DbContext context.Context
	Client    *randomDB.Client // Here we can hook up a random DB and connect it to this microservice
}

// Connect : to the specific DB that is required for this microservice
func (c *DatastoreDAO) Connect() {
	// Inside this we will have database specific code to connect to a particular db
}
