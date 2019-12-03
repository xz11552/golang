package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func getTicket() ([]ticket, bool) {
	var ticketList []ticket

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8004)/Ticket_DB")
	if err != nil {
		fmt.Println("DB err: ", err)
		return ticketList, false
	}
	defer db.Close()

	ticketRow, err := db.Query("SELECT * FROM Ticket_amount")
	if err != nil {
		fmt.Println("sql err: ", err)
		return ticketList, false
	}
	defer ticketRow.Close()

	for ticketRow.Next() {
		var tic ticket
		if err := ticketRow.Scan(&tic.Type, &tic.Name, &tic.Amount, &tic.Price); err != nil {
			fmt.Println("search err: ", err)
			return ticketList, false
		}
		ticketList = append(ticketList, tic)
	}
	if err := ticketRow.Err(); err != nil {
		fmt.Println("err: ", err)
		return ticketList, false
	}

	return ticketList, true

}

func updateTicket(order order) bool {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8004)/Ticket_DB")
	if err != nil {
		fmt.Println("DB err: ", err)
		return false
	}
	defer db.Close()

	sqlSearch := "SELECT ticket_amounts FROM Ticket_amount WHERE ticket_type = " + strconv.Itoa(order.OrderType)
	ticketRow, err := db.Query(sqlSearch)
	if err != nil {
		fmt.Println("sql err: ", err)
		return false
	}
	defer ticketRow.Close()

	var totalAmount int
	for ticketRow.Next() {
		if err := ticketRow.Scan(&totalAmount); err != nil {
			fmt.Println("search err: ", err)
			return false
		}
	}
	if err := ticketRow.Err(); err != nil || totalAmount < order.OrderAmount {
		fmt.Println("err: ", err)
		return false
	}

	sqlUpdateAmounts := fmt.Sprintf(
		"BEGIN; "+
			"INSERT INTO Order_Info (order_type,order_name,order_phone,order_amount,order_total) "+
			"VALUES ("+
			"%d,"+
			"%s,"+
			"%s,"+
			"%d,"+
			"%d); "+
			"UPDATE Ticket_amount "+
			"SET ticket_amounts = ticket_amounts - "+"%d "+
			"WHERE ticket_type = "+"%d; "+
			"COMMIT; ",
		order.OrderType,
		order.OrderName,
		order.OrderPhone,
		order.OrderAmount,
		order.OrderTotal,
		order.OrderAmount,
		order.OrderType,
	)

	result, err := db.Exec(sqlUpdateAmounts)
	if err != nil {
		fmt.Println("sql err: ", err)
		return false
	}

	idAff, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected failed:", err)
		return false
	}

	if idAff != 0 {
		return true
	}

	return false
}

func getOrder(input string) ([]order, bool) {
	var orderList []order

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8004)/Ticket_DB")
	if err != nil {
		fmt.Println("DB err: ", err)
		return orderList, false
	}
	defer db.Close()

	sql := fmt.Sprintf(
		"SELECT * FROM Order_Info WHERE order_phone = '%s'",
		input,
	)

	orderRow, err := db.Query(sql)
	if err != nil {
		fmt.Println("sql err: ", err)
		return orderList, false
	}
	defer orderRow.Close()

	for orderRow.Next() {
		var ord order
		if err := orderRow.Scan(&ord.OrderID, &ord.OrderType, &ord.OrderName, &ord.OrderPhone,
			&ord.OrderAmount, &ord.OrderTotal); err != nil {
			fmt.Println("search err: ", err)
			return orderList, false
		}
		orderList = append(orderList, ord)
	}
	if err := orderRow.Err(); err != nil {
		fmt.Println("err: ", err)
		return orderList, false
	}

	return orderList, true
}
