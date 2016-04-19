package data

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
)

// GetClients retrieves all available clients from the database.
func (db *data) GetClients() (*model.Clients, error) {
	records := &model.Clients{}

	err := db.Order(
		"name ASC",
	).Find(
		records,
	).Error

	return records, err
}

// CreateClient creates a new client.
func (db *data) CreateClient(record *model.Client) error {
	return db.Create(
		record,
	).Error
}

// UpdateClient updates a client.
func (db *data) UpdateClient(record *model.Client) error {
	return db.Save(
		record,
	).Error
}

// DeleteClient deletes a client.
func (db *data) DeleteClient(record *model.Client) error {
	return db.Delete(
		record,
	).Error
}

// GetClient retrieves a specific client from the database.
func (db *data) GetClient(id string) (*model.Client, *gorm.DB) {
	record := &model.Client{}

	res := db.Where(
		"id = ?",
		id,
	).Or(
		"slug = ?",
		id,
	).First(
		record,
	)

	return record, res
}

// GetClientPacks retrieves packs for a client.
func (db *data) GetClientPacks(id int) (*model.Packs, error) {
	records := &model.Packs{}

	err := db.Model(
		&model.Client{
			ID: id,
		},
	).Association(
		"Packs",
	).Find(
		records,
	).Error

	return records, err
}

// GetClientHasPack checks if a specific pack is assigned to a client.
func (db *data) GetClientHasPack(parent, id int) bool {
	record := &model.Pack{
		ID: id,
	}

	count := db.Model(
		&model.Client{
			ID: parent,
		},
	).Association(
		"Packs",
	).Find(
		record,
	).Count()

	return count > 0
}

func (db *data) CreateClientPack(parent, id int) error {
	return db.Model(
		&model.Client{
			ID: parent,
		},
	).Association(
		"Packs",
	).Append(
		model.Pack{
			ID: id,
		},
	).Error
}

func (db *data) DeleteClientPack(parent, id int) error {
	return db.Model(
		&model.Client{
			ID: parent,
		},
	).Association(
		"Packs",
	).Delete(
		model.Pack{
			ID: id,
		},
	).Error
}