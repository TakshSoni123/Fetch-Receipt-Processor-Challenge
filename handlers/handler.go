package handlers

import (
	"fmt"
	"net/http"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API called")
}
