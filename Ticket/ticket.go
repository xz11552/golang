package main

import (
	"fmt"

	_ "github.com/xz11552/golang/Ticket/DB_Helper"
)

func main() {
	ticketsss := ticket
	num := updateTicket(1, 99)
	fmt.Println(num)
}
