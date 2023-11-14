package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/nolostra/goproject/api"
	"github.com/nolostra/goproject/internal/tools"
	log "github.com/sirupsen/logrus"
)

// get Request
func GetCoinBalance(w http.ResponseWriter, r *http.Request){
	var params  = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	err  = decoder.Decode(&params , r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()

	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.UserName)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var Response = api .CoinBalanceResponse{
		Balance: (*&tokenDetails).Coins,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(Response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

// Post Request
func PostCoinBalance(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	var Data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		log.Error("Error decoding request body:", err)
		api.InternalErrorHandler(w)
		return
	}

	// Close the request body after reading
	defer r.Body.Close()
	// post operation to be happen on database


	// Create the response
	var Response = struct {
		Message string `json:"message"`
	}{
		Message: "Data posted successfully",
	}
	fmt.Println("Posted Data Receieved", Data)

	// Set the response header
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	err = json.NewEncoder(w).Encode(Response)
	if err != nil {
		log.Error("Error encoding response:", err)
		api.InternalErrorHandler(w)
		return
	}
}

// Delete Request 
func DelCoinBalance(w http.ResponseWriter, r *http.Request){
	var Data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		log.Error("Error decoding request body:", err)
		api.InternalErrorHandler(w)
		return
	}

	// Close the request body after reading
	defer r.Body.Close()

	// should delete operation be happen on database

	// Create the response
	var Response = struct {
		Message string `json:"message"`
	}{
		Message: "Data Deleted successfully",
	}
	fmt.Println("Posted Deleted Receieved", Data)

	// Set the response header
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	err = json.NewEncoder(w).Encode(Response)
	if err != nil {
		log.Error("Error encoding response:", err)
		api.InternalErrorHandler(w)
		return
	}
}

func ChangeCoinBalance(w http.ResponseWriter, r *http.Request){
	var Data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		log.Error("Error decoding request body:", err)
		api.InternalErrorHandler(w)
		return
	}

	// Close the request body after reading
	defer r.Body.Close()

	// should delete operation be happen on database

	// Create the response
	var Response = struct {
		Message string `json:"message"`
	}{
		Message: "Data Altered successfully",
	}
	fmt.Println("Posted Altered Receieved", Data)

	// Set the response header
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	err = json.NewEncoder(w).Encode(Response)
	if err != nil {
		log.Error("Error encoding response:", err)
		api.InternalErrorHandler(w)
		return
	}
}