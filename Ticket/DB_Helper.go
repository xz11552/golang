package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type ticket struct {
	ID     int
	name   string
	amount int
}

func getTicket() []ticket {
	var ticketList []ticket
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8004)/Ticket_DB")
	if err != nil {
		fmt.Println("DB err: ", err)
		panic(1)
	}
	defer db.Close()

	ticketRow, err := db.Query("SELECT * FROM Ticket_amount")
	if err != nil {
		fmt.Println("sql err: ", err)
	}
	defer ticketRow.Close()

	for ticketRow.Next() {
		var tic ticket
		if err := ticketRow.Scan(&tic.ID, &tic.name, &tic.amount); err != nil {
			fmt.Println("search err: ", err)
		}
		ticketList = append(ticketList, tic)
	}
	if err := ticketRow.Err(); err != nil {
		fmt.Println("err: ", err)
	}

	return ticketList
}

func updateTicket(ID, updateAmount int) int {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8004)/Ticket_DB")
	if err != nil {
		fmt.Println("DB err: ", err)
		panic(1)
	}
	defer db.Close()

	sql := fmt.Sprintf(
		"UPDATE Ticket_amount SET ticket_amounts = '%d'  WHERE ID = '%d'",
		updateAmount,
		ID,
	)

	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println("sql err: ", err)
	}

	idAff, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected failed:", err)
	}

	if idAff != 0 {
		return 1
	}

	return 0
}
