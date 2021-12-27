package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
	_ "github.com/sijms/go-ora/v2"
)

func main() {
	service_url := "oracle://GA731852:uwygnnr2@192.168.1.114:1521/ORCLCDB.localdomain"
	db, err := sql.Open("oracle", service_url)
	if err != nil {
		fmt.Println("1err", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("select * from contacts")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			CONTACT_ID   sql.NullString
			FIRST_NAME   sql.NullString
			LAST_NAME    sql.NullString
			EMAIL        sql.NullString
			PHONE        sql.NullString
			CUSTORMER_ID sql.NullString
		)

		rows.Scan(
			&CONTACT_ID,
			&FIRST_NAME,
			&LAST_NAME,
			&EMAIL,
			&PHONE,
			&CUSTORMER_ID,
		)

		fmt.Println(CONTACT_ID.String, FIRST_NAME.String, LAST_NAME.String, EMAIL.String, PHONE.String, CUSTORMER_ID.String)
	}
}
