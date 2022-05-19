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

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var nameTable string
		err = rows.Scan(&nameTable)
		if err != nil {
			panic(err.Error())
		}

		table := models.Tables{
			Name: nameTable,
		}

		tables = append(tables, table)

		if err = rows.Err(); err != nil {
			panic(err.Error())
		}
	}

	return tables, nil
}

func (client *Mysql) DescribeTables() ([]models.TableDescribe, error) {
	var tableDesc []models.TableDescribe

	// TODO: refactor
	db, err := sql.Open("mysql", "root:root@/testdash")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("DESCRIBE datatest")
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var (
			field     string
			typeTb    string
			null      string
			key       string
			defaultTb sql.NullString
			extra     string
		)
		err = rows.Scan(&field, &typeTb, &null, &key, &defaultTb, &extra)
		if err != nil {
			panic(err.Error())
		}

		table := models.TableDescribe{
			Field: field,
		}

		tableDesc = append(tableDesc, table)

		if err = rows.Err(); err != nil {
			panic(err.Error())
		}
	}
	return tableDesc, nil
}
