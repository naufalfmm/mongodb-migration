package main

import (
	"context"

	"github.com/naufalfmm/mongodb-migration/config"
	"github.com/naufalfmm/mongodb-migration/constants"
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

		migr = migration.MongoMigration{}
		dr   = driver.MongoDriver{}
	)

	ctx := context.TODO()

	cfg.SetURI()
	dr.SetClientWithContext(ctx, &cfg)

	hs := history.MigrationRecord{
		DB:             dr.GetDB(),
		CollectionName: "migrationHistory",
	}

	err := migr.StartMigrationWithDriver("./example/migrate/migrations/", &dr, &hs)
	if err != nil {
		panic(err)
	}

	err = migr.Run(ctx, constants.DOWN, 2)
	if err != nil {
		panic(err)
	}
}
