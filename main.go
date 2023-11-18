package main

import (
	"fmt"
	"net/http"
	"receipt-processor-backend/handlers"
	"receipt-processor-backend/models"

	"github.com/julienschmidt/httprouter"
)

func main() {

	fmt.Print("Started service...")

	models.Receipts = make(map[string]models.Receipt)

	router := httprouter.New()

	router.POST("/receipts/process", handlers.ProcessReceipt)
	router.GET("/receipts/:id/points", handlers.GetPoints)

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
