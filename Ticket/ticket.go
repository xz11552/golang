package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", httpHandle)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	// ticketsss := getTicket()
	// fmt.Println(ticketsss)
	// num := updateTicket(1, 99)
	// fmt.Println(num)
}

func httpHandle(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(writer, "404 not Found", http.StatusNotFound)
		return
	}
	var value url.Values
	switch request.Method {
	case "GET":
		value = request.URL.Query()
	case "POST":
		if err := request.ParseForm(); err != nil {
			fmt.Println("ParseForm() err: ", err)
			return
		}
		value = request.Form
	}

	var param inJSON
	err := json.Unmarshal([]byte(GetValue(value, "name")), &param)
	if err != nil {
		fmt.Println("inJSON error:", err)
	}

	outParam := selectAPI(param)

	json.NewEncoder(writer).Encode(param)
}

func selectAPI(param inJSON) string {
	outParam := ""
	switch param.FuncName {
	case "API_GetTicket":

		outParam = APIGetTicketInformation()
		return outParam
	case "API_OrderTicket":
		var orderParam order

		err := json.Unmarshal([]byte(param.Parameter), &orderParam)
		if err != nil {
			fmt.Println("order JSON error:", err)
		}

		outParam = APIOrderTicket(orderParam)
		return outParam
	case "APIGetOrderInformation":
		var searchParam searchOrderParam

		err := json.Unmarshal([]byte(param.Parameter), &searchParam)
		if err != nil {
			fmt.Println("order JSON error:", err)
		}

		outParam = APIGetOrderInformation(searchParam.Phone)
		return outParam
	}
	return outParam
}

//取得Get資料
func GetValue(v url.Values, key string) string {
	if v == nil {
		return ""
	}
	vs := v[key]
	if len(vs) != 0 {
		return vs[0]
	}
	return ""
}
