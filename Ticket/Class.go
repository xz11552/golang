package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type outJSON struct {
	Status    bool
	Parameter string
}

type inJSON struct {
	FuncName  string
	Parameter string
}

type ticket struct {
	Type   int
	Name   string
	Amount int
	Price  int
}

type order struct {
	OrderID     int
	OrderType   int
	OrderName   string
	OrderPhone  string
	OrderAmount int
	OrderTotal  int
}

type searchOrderParam struct {
	Phone string
}
