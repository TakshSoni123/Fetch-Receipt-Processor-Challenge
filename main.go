package main

import (
	"fmt"
	"net/http"
	"receipt-processor-backend/handlers"
)

func main() {

	fmt.Print("Started service...")

	http.HandleFunc("/receipts/process", handlers.ProcessReceipt)

	http.ListenAndServe(":8080", nil)
}
