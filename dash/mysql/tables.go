package mysql

import (
	"adminmsyql/config"
	"adminmsyql/dash/models"
	"database/sql"
	"fmt"
	"log"
)

func (client *Mysql) ListTables() ([]models.Tables, error) {
	var tables []models.Tables
	cfg, err := config.ParserConfig()
	if err != nil {
		log.Fatal(err)
	}

	credentials := fmt.Sprintf("%s:%s@/mysql", cfg.Credentials.Username, cfg.Credentials.Password)

	db, err := sql.Open("mysql", credentials)

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

func (client *Mysql) DescribeTables(nameTable string) ([]models.TableDescribe, error) {
	var tableDesc []models.TableDescribe

	// TODO: refactor
	cfg, err := config.ParserConfig()
	if err != nil {
		log.Fatal(err)
	}

	credentials := fmt.Sprintf("%s:%s@/mysql", cfg.Credentials.Username, cfg.Credentials.Password)

	db, err := sql.Open("mysql", credentials)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	tableSeleted := fmt.Sprintf("DESCRIBE %s", nameTable)
	rows, err := db.Query(tableSeleted)
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
			Type:  typeTb,
			Null:  null,
		}

		tableDesc = append(tableDesc, table)

		if err = rows.Err(); err != nil {
			panic(err.Error())
		}
	}
	return tableDesc, nil
}
