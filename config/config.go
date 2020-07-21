package config

import (
	"fmt"
)

type MongoConfig struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBURI      *string

	DBTimeout *int
}

func (mc *MongoConfig) SetURL() error {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", mc.DBUser, mc.DBPassword, mc.DBHost, mc.DBPort, mc.DBName)

	mc.DBURI = &uri

	return nil
}
