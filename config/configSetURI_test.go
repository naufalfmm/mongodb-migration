package config

import (
	"testing"
)

func ItShouldSetValueOfDBURI(t *testing.T) {
	config := Config{
		Name:     "db-coba",
		User:     "user",
		Password: "password",
		Host:     "localhost",
		Port:     "1111",
	}

	config.SetURI()

	// Validate the value
	urlShould := "mongodb://user:password@localhost:1111/db-coba"

	if *config.DBURI() != urlShould {
		t.Errorf("The value of DBURI should be %s", urlShould)
	}
}

func ItShouldReplaceValueOfDBURI(t *testing.T) {
	uriStart := "mongodb://user123:password123@localhost:2222/db-mana"

	config := Config{
		Name:     "db-coba",
		User:     "user",
		Password: "password",
		Host:     "localhost",
		Port:     "1111",

		URI: &uriStart,
	}

	config.SetURI()

	// Validate the value
	urlShould := "mongodb://user:password@localhost:1111/db-coba"

	if *config.DBURI() != urlShould {
		t.Errorf("The value of DBURI should be %s", urlShould)
	}
}

func TestSetURI(t *testing.T) {
	t.Run("It should set value of DBURI", ItShouldSetValueOfDBURI)
	t.Run("It should replace value of DBURI", ItShouldReplaceValueOfDBURI)
}
