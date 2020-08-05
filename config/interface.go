package config

// DatabaseConfig --> Basic configuration of MongoDB database
type DatabaseConfig interface {
	DBName() string
	DBUser() string
	DBPassword() string
	DBHost() string
	DBPort() string

	DBURI() *string
	DBTimeout() *int64

	SetURI() error
}
