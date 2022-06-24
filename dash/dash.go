package dash

import (
	"adminmsyql/dash/adapter"
	"adminmsyql/dash/models"
	"adminmsyql/dash/mysql"
)

// Dash
type Dash interface {
	//
	GetCapabilities() []adapter.Capability
	//
	ListProfile() ([]models.Credential, error)
	//
	ListTables() ([]models.Tables, error)
	//
	DescribeTables(nameTable string) ([]models.TableDescribe, error)
	//
	InfoStatusBar() models.Info
	//
	LoadClients() error
}

func New(clientType *string) (Dash, error) {
	var client Dash

	switch *clientType {
	case "mysql":
		client = new(mysql.Mysql)
	}

	err := client.LoadClients()
	if err != nil {
		return nil, err
	}

	return client, nil
}
