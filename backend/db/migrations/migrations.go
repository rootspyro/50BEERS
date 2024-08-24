package migrations

import (
	"github.com/rootspyro/50BEERS/config/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type MigrationManager struct {
	database *mongo.Database
}

func NewMigrationManager(db *mongo.Database) *MigrationManager {
	return &MigrationManager{
		database: db,
	}
}

func(m *MigrationManager) Migrate() error {

	log.Info("Running Migrations")
	return nil
}
