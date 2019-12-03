package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//取得票卷資訊
func APIGetTicketInformation() string {
	output := outJSON{
		Status:    false,
		Parameter: "",
	}

	ticketList, status := getTicket()

	if status {
		ticketInfo, err := json.Marshal(ticketList)
		if err != nil {
			fmt.Println("JSON err:", err)
		} else {
			output.Status = true
			output.Parameter = string(ticketInfo)
		}
	}

	outparam, err := json.Marshal(output)
	if err != nil {
		fmt.Println("JSON err:", err)
	}

	return string(outparam)
}

//進行訂票
func APIOrderTicket(Order order) string {

	output := outJSON{
		Status:    false,
		Parameter: "",
	}

	output.Status = updateTicket(Order)

	outparam, err := json.Marshal(output)
	if err != nil {
		fmt.Println("JSON err:", err)
	}

	return string(outparam)
}

//取得訂票資訊
func APIGetOrderInformation(phone string) string {
	output := outJSON{
		Status:    false,
		Parameter: "",
	}

	orderList, status := getOrder(phone)

	if status {
		orderInfo, err := json.Marshal(orderList)
		if err != nil {
			fmt.Println("JSON err:", err)
		} else {
			output.Status = true
			output.Parameter = string(orderInfo)
		}
	}

	outparam, err := json.Marshal(output)
	if err != nil {
		fmt.Println("JSON err:", err)
	}

	return string(outparam)
}
