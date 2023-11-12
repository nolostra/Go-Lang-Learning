package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/nolostra/goproject/api"
	"github.com/nolostra/goproject/internal/tools"
	log "github.com/sirupsen/logrus"
)

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

func PostCoinBalance(w http.ResponseWriter, r *http.Request){
	// Read the request body
	var Data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Close the request body after reading
	defer r.Body.Close()

	// Create the response
	var Response = struct {
		Data map[string]interface{} `json:"data"`
		Code int                    `json:"code"`
	}{
		Data: Data,
		Code: http.StatusOK,
	}

	// Set the response header
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	err = json.NewEncoder(w).Encode(Response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}