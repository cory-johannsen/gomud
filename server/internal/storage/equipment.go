package storage

import (
	"context"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
	log "github.com/sirupsen/logrus"
)

type Equipment struct {
	database *Database
	loaders  *loader.Loaders
	items    map[int]domain.Item
}

func NewEquipment(database *Database, loaders *loader.Loaders) *Equipment {
	return &Equipment{
		database: database,
		loaders:  loaders,
		items:    make(map[int]domain.Item),
	}
}

func (e *Equipment) CreateItem(ctx context.Context, item domain.Item) (domain.Item, error) {
	log.Debugf("Creating item %s", item.Name())
	var id int
	tx, err := e.database.Conn.Begin(ctx)
	if err != nil {
		log.Errorf("failed to start transaction: %s", err)
		return nil, err
	}
	row := e.database.Conn.QueryRow(ctx, "INSERT INTO items (name, description, mass_in_grams, type, cost) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		item.Name(), item.Description(), item.MassInGrams(), item.Type(), item.Cost())
	err = row.Scan(&id)
	if err != nil {
		log.Errorf("failed to insert item: %s", err)
		return nil, err
	}
	log.Debugf("Created item %s with id %d", item.Name(), id)
	err = tx.Commit(ctx)
	if err != nil {
		log.Errorf("failed to commit transaction: %s", err)
		return nil, err
	}
	return e.FetchItemByID(ctx, id)
}

func (e *Equipment) FetchItemByID(ctx context.Context, id int) (domain.Item, error) {
	log.Debugf("Fetching item %d", id)
	if item, ok := e.items[id]; ok {
		return item, nil
	}

	var name string
	var description string
	var massInGrams int
	var cost int
	var itemType domain.ItemType
	err := e.database.Conn.QueryRow(ctx, "SELECT name, description, mass_in_grams, type, cost FROM items WHERE id = $1", id).
		Scan(&name, &description, &massInGrams, &itemType, &cost)
	if err != nil {
		log.Errorf("failed to get item: %s", err)
		return nil, err
	}

	templateItem, err := e.loaders.EquipmentLoader.ItemFromName(name)
	if err != nil {
		log.Errorf("failed to get item template: %s", err)
		return nil, err
	}

	item := templateItem.NewInstance(id)
	e.items[id] = item
	return item, nil
}
