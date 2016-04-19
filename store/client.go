package store

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
	"golang.org/x/net/context"
)

// GetClients retrieves all available clients from the database.
func GetClients(c context.Context) (*model.Clients, error) {
	return FromContext(c).GetClients()
}

// CreateClient creates a new client.
func CreateClient(c context.Context, record *model.Client) error {
	return FromContext(c).CreateClient(record)
}

// UpdateClient updates a client.
func UpdateClient(c context.Context, record *model.Client) error {
	return FromContext(c).UpdateClient(record)
}

// DeleteClient deletes a client.
func DeleteClient(c context.Context, record *model.Client) error {
	return FromContext(c).DeleteClient(record)
}

// GetClient retrieves a specific client from the database.
func GetClient(c context.Context, id string) (*model.Client, *gorm.DB) {
	return FromContext(c).GetClient(id)
}

// GetClientPacks retrieves packs for a client.
func GetClientPacks(c context.Context, id int) (*model.Packs, error) {
	return FromContext(c).GetClientPacks(id)
}

// GetClientHasPack checks if a specific pack is assigned to a client.
func GetClientHasPack(c context.Context, parent, id int) bool {
	return FromContext(c).GetClientHasPack(parent, id)
}

func CreateClientPack(c context.Context, parent, id int) error {
	return FromContext(c).CreateClientPack(parent, id)
}

func DeleteClientPack(c context.Context, parent, id int) error {
	return FromContext(c).DeleteClientPack(parent, id)
}