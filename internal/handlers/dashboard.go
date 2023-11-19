package handlers

import (
	// io helps to read all the body which is sent
	"encoding/json"
	"fmt"
	"net/http"
	"errors"
	"github.com/nolostra/goproject/api"
	"github.com/nolostra/goproject/internal/token"
	log "github.com/sirupsen/logrus"
)

var DashboardError = errors.New("Invalid token ")

func Dashboard(w http.ResponseWriter, r *http.Request){
	authHeader := r.Header.Get("Authorization")

	
	tokenString, err := token.VerifyJWTToken(authHeader)


	if err != nil {
		log.Error(err)
		api.LoginErrorHandler(w,err)
		return
	}
	fmt.Println("JWT token Verified")

	var Response = struct{
		Response string `json:"response"`
    	Body     string `json:"body"`
	}{
		Response: `Token Verified` + tokenString.Raw,
		Body: "some object",
	}

	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(Response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}