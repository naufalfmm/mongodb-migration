<h3 align="center">MongoDB Migration</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Release](https://img.shields.io/github/release/naufalfmm/mongodb-migration.svg)](https://github.com/naufalfmm/mongodb-migration/releases)
[![License](https://img.shields.io/badge/license-Apache-blue.svg)](/LICENSE)

</div>

---

<p align="center"> MongoDB migration library for Golang.
    <br> 
</p>

## 📝 Table of Contents

- [Features](#features)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Built Using](#built_using)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 Features <a name = "features"></a>

- Migrate up or down all migration files or specific file
- Create history collection as migration files tracker
- Implement interface as template of the library
- Atomic migrations
- Automatically migrate down if migrate up fails
- Use json migration file as `*.sql` for sql migration file
- Implement `db.RunCommand()`. You can check [here](https://docs.mongodb.com/manual/reference/method/db.runCommand/)

## 🏁 Getting Started <a name = "getting_started"></a>

### Prerequisites

There is no specific prerequisities, just need [Golang](http://www.golang.org) :)

### Installing

Install the library by running the following command

```
go get -u github.com/naufalfmm/mongodb-migration
```

## 🎈 Usage <a name="usage"></a>

### Example

Example of usage of the library in your project

```go
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

		migr = migration.MongoMigration{}
		mongoDriver   = driver.MongoDriver{}
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

```

or you can open the example [here](/example/migrate/migrate.go)

### Migration File

You can check the example [here](/example/migrate/migrations)

#### Naming

You can name your json migration file as your need. If you want to follow the standard, you name the file by

```
<timestamp>_name_of_your_migration_file.json
```

For example

```
20200716152127_create_full_name_of_individual_of_client.json
```

#### Content

The content of migration file is the command - formed as document or string - of the `db.runCommand()`. Please refer the list of commands [here](https://docs.mongodb.com/manual/reference/command/). 

Example of the [content](/example/migrate/migrations/20200716152127_create_full_name_of_individual_of_client.json)

```
{
    "up": {
        "update": "Customer",
        "updates": [
            {
                "q": {"individual": {"$ne": null}},
                "u": [
                    {
                        "$set": { "individual.full_name": { "$concat": ["$individual.first_name", " ", "$individual.last_name"] } }
                    }
                ],
                "multi": true
            }
        ]
    },
    "down": {
        "update": "Customer",
        "updates": [
            {
                "q": {},
                "u": [
                    {
                        "$unset": "individual.full_name"
                    }
                ],
                "multi": true
            }
        ]
    }
}
```
`up` is for migrate up and `down` is for migrate down.

## ⛏️ Built Using <a name = "built_using"></a>

- [Golang](https://www.golang.org/) - Language

## ✍️ Authors <a name = "authors"></a>

- [@naufalfmm](https://github.com/naufalfmm) - Idea & Initial work

If you interest to join as contributor, please contact me on [email](muhammadnaufalfm@gmail.com) or just create PR :)

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Hat tip to anyone whose code was used
