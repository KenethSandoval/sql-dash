package dash

import (
	"adminmsyql/dash/adapter"
	"adminmsyql/dash/mysql"

	_ "github.com/go-sql-driver/mysql"
)

// Dash
type Dash interface {
	//
	GetCapabilities() []adapter.Capability
	//
	ListProfile(profile *string) ([]string, error)

	//
	ListTables() ([]string, error)
}

func New(clientType *string) (Dash, error) {
	var client Dash

	switch *clientType {
	case "mysql":
		client = new(mysql.Mysql)
	}

	return client, nil
}
