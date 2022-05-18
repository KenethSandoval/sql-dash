package mysql

import (
	"adminmsyql/dash/models"
	"database/sql"
)

func (client *Mysql) ListTables() ([]models.Tables, error) {
	var tables []models.Tables

	db, err := sql.Open("mysql", "root:root@/testdash")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SHOW TABLES")
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
			table := models.Tables{
				Name: value,
			}
			tables = append(tables, table)
		}
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return tables, nil
}
