package mysql

import (
	"adminmsyql/dash/adapter"
	"adminmsyql/dash/models"
	"database/sql"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Mysql struct {
}

const (
	width = 96
)

var (
	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)

	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)
)

func errorDialog(error string) {
	doc := strings.Builder{}

	errorMessage := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(error)
	ui := lipgloss.JoinVertical(lipgloss.Center, errorMessage)
	dialog := lipgloss.Place(width, 9,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(subtle),
	)

	doc.WriteString(dialog + "\n\n")

	fmt.Println(docStyle.Render(doc.String()))
}

func (client *Mysql) ListProfile() ([]models.Credential, error) {
	var users []models.Credential

	db, err := sql.Open("mysql", "root:root@/mysql")

	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT user, host, insert_priv, select_priv, update_priv, delete_priv, create_priv, drop_priv,
                               grant_priv, index_priv, alter_priv
                               FROM db where user not like '%mysql%'`)

	if err != nil {
		errorDialog("127.0.0.1:3306: connect: connection refused")
	}

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

		err = rows.Scan(&user, &host, &insertPriv, &selectPriv, &updatePriv, &deletePriv, &createPriv, &dropPriv,
			&grantPriv, &indexPriv, &alterPriv)
		if err != nil {
			panic(err.Error())
		}

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
