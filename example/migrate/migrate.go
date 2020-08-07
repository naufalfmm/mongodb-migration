package main

import (
	"context"

	"github.com/naufalfmm/mongodb-migration/constants"

	"github.com/naufalfmm/mongodb-migration/config"
	"github.com/naufalfmm/mongodb-migration/constants/direction"
	"github.com/naufalfmm/mongodb-migration/constants/steps"
	"github.com/naufalfmm/mongodb-migration/driver"
	"github.com/naufalfmm/mongodb-migration/history"
	"github.com/naufalfmm/mongodb-migration/migration"
)

func main() {
	var (
		DBName     = "db_migration_trial"
		DBUser     = "user"
		DBPassword = "123456789"
		DBHost     = "localhost"
		DBPort     = "27017"

		cfg config.Config = config.Config{
			Name:     DBName,
			User:     DBUser,
			Password: DBPassword,
			Host:     DBHost,
			Port:     DBPort,
		}

		migr        = migration.MongoMigration{}
		mongoDriver = driver.MongoDriver{}
	)

	ctx := context.TODO()

	cfg.SetURI()
	mongoDriver.SetClientWithContext(ctx, &cfg)

	historyRecord := history.MigrationRecord{
		DB:             mongoDriver.GetDB(),
		CollectionName: constants.DEFAULT_MIGRATION_HISTORY_COLLECTION,
	}

	err := migr.StartMigrationWithDriver(ctx, "./example/migrate/migrations/", &mongoDriver, &historyRecord)
	if err != nil {
		panic(err)
	}

	err = migr.Run(ctx, direction.UP, steps.ALL)
	if err != nil {
		panic(err)
	}
}
