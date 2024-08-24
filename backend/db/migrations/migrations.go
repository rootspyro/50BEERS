package migrations

import (
	"context"
	"fmt"

	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db/models"
	"go.mongodb.org/mongo-driver/bson"
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

	log.Info("Running Migrations...")

	for _, collectionName := range models.Collections {
		log.Info(fmt.Sprintf("Creating %s collection...", collectionName))	
		if err := m.createCollection(collectionName); err != nil {
			return nil
		}
		log.Info(fmt.Sprintf("The %s collection was successfully created!", collectionName))	
	}

	log.Info("All migrations where successfully executed!")

	return nil
}

func(m *MigrationManager) createCollection(collectionName string) error {
	command := bson.D{{"create", collectionName}}

	var result bson.M

	err := m.database.RunCommand(context.TODO(), command).Decode(&result)

	return err
}
