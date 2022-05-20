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

	rows, err := db.Query(`SELECT user, host, insert_priv, select_priv, update_priv, delete_priv, create_priv, drop_priv, reload_priv,
                               shutdown_priv, process_priv, file_priv, grant_priv, references_priv, index_priv, alter_priv
                               FROM user where user not like '%mysql%'`)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var (
			user           string
			host           string
			insertPriv     string
			selectPriv     string
			updatePriv     string
			deletePriv     string
			createPriv     string
			dropPriv       string
			reloadPriv     string
			shutdownPriv   string
			processPriv    string
			filePriv       string
			grantPriv      string
			referencesPriv string
			indexPriv      string
			alterPriv      string
		)

		err = rows.Scan(&user, &host, &insertPriv, &selectPriv, &updatePriv, &deletePriv, &createPriv, &dropPriv,
			&reloadPriv, &shutdownPriv, &processPriv, &filePriv, &grantPriv, &referencesPriv, &indexPriv, &alterPriv)
		if err != nil {
			panic(err.Error())
		}

		userFind := models.Credential{
			Name:       user,
			Host:       host,
			InsertPriv: insertPriv,
			SelectPriv: selectPriv,
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
