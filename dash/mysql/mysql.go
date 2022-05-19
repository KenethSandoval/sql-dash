package mysql

import (
	"adminmsyql/dash/adapter"
	"adminmsyql/dash/models"
	"database/sql"
)

type Mysql struct {
}

func (client *Mysql) ListProfile() ([]models.Credential, error) {
	var users []models.Credential

	db, err := sql.Open("mysql", "root:root@/mysql")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT user, host, insert_priv FROM user where user not like '%mysql%'")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var (
			user       string
			host       string
			insertPriv string
		)

		err = rows.Scan(&user, &host, &insertPriv)
		if err != nil {
			panic(err.Error())
		}

		userFind := models.Credential{
			Name:       user,
			Host:       host,
			InsertPriv: insertPriv,
		}

		users = append(users, userFind)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return users, nil
}

func (client *Mysql) GetCapabilities() []adapter.Capability {
	var caps []adapter.Capability

	caps = append(caps, adapter.Capability{
		ID:   "users",
		Name: "Users DB",
	})

	caps = append(caps, adapter.Capability{
		ID:   "tables",
		Name: "Tables DB",
	})

	return caps
}
