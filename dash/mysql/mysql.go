package mysql

import (
	"adminmsyql/dash/adapter"
	"database/sql"
)

type Mysql struct {
}

func (client *Mysql) ListProfile(profile *string) ([]string, error) {
	var usersFind []string

	db, err := sql.Open("mysql", "root:@/mysql")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT user FROM user")
	if err != nil {
		panic(err.Error())
	}

	colums, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(colums))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string

		for _, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			usersFind = append(usersFind, value)
		}
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return usersFind, nil
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
