package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"tuidb/config"
	"tuidb/dash/adapter"
	"tuidb/dash/models"
	"tuidb/ui/common"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Username string
	Password string
	Database string
}

func (client *Mysql) ListProfile() ([]models.Credential, error) {
	var users []models.Credential

	credentials := fmt.Sprintf("%s:%s@/mysql", client.Username, client.Password)

	db, err := sql.Open("mysql", credentials)

	// TODO: add flag debug
	defer func() {
		if r := recover(); r != nil {
			common.ErrorDialog(err.Error())
		}
	}()

	rows, err := db.Query(`SELECT user, host,
                                      insert_priv, select_priv,
                                      update_priv, delete_priv,
                                      create_priv, drop_priv,
                                      grant_priv, index_priv, alter_priv
                               FROM db where user not like '%mysql%'`)

	defer db.Close()

	for rows.Next() {
		var (
			user       string
			host       string
			insertPriv string
			selectPriv string
			updatePriv string
			deletePriv string
			createPriv string
			dropPriv   string
			grantPriv  string
			indexPriv  string
			alterPriv  string
		)

		err = rows.Scan(&user, &host, &insertPriv,
			&selectPriv, &updatePriv, &deletePriv,
			&createPriv, &dropPriv, &grantPriv,
			&indexPriv, &alterPriv)

		userFind := models.Credential{
			Name:       user,
			Host:       host,
			InsertPriv: insertPriv,
			SelectPriv: selectPriv,
			UpdatePriv: updatePriv,
			DeletePriv: deletePriv,
			CreatePriv: createPriv,
			DropPriv:   dropPriv,
			GrantPriv:  grantPriv,
			IndexPriv:  indexPriv,
			AlterPriv:  alterPriv,
		}

		users = append(users, userFind)
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

func (client *Mysql) InfoStatusBar() models.Info {
	var info models.Info

	credentials := fmt.Sprintf("%s:%s@/mysql", client.Username, client.Password)

	db, err := sql.Open("mysql", credentials)

	defer func() {
		if r := recover(); r != nil {
			common.ErrorDialog(err.Error())
		}
	}()

	rows, err := db.Query(`select version() as "version", current_user() as "user" from dual`)

	defer db.Close()

	for rows.Next() {
		var (
			version string
			user    string
		)

		err = rows.Scan(&version, &user)

		info = models.Info{
			Version:  version,
			UserConn: user,
		}

	}

	return info
}

func (client *Mysql) LoadClients() error {
	cfg, err := config.ParserConfig()
	if err != nil {
		log.Fatal(err)
	}

	client.Username = cfg.Settings.Username
	client.Password = cfg.Settings.Password
	client.Database = cfg.Settings.Database

	return nil
}
