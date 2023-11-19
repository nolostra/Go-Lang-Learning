package handlers

import (
	// io helps to read all the body which is sent
	"io"
	"encoding/json"
	"fmt"
	"net/http"
	"errors"
	"github.com/nolostra/goproject/api"
	"github.com/nolostra/goproject/internal/token"
	log "github.com/sirupsen/logrus"
)

var LoginError = errors.New("Invalid Username or password ")

func Login(w http.ResponseWriter, r *http.Request){
	var userName string = r.URL.Query().Get("username")
	var passWord string = r.URL.Query().Get("password")
	body, err := io.ReadAll(r.Body)
	
	if userName == "" || passWord == "" {
		api.LoginErrorHandler(w,LoginError)
		log.Error(LoginError)
		return 
	}

	fmt.Println("username=>",userName)
	fmt.Println("password=>",passWord)
	fmt.Println("URLQuery=>",r.URL.RawQuery)
	fmt.Println("BodyQuery=>",string(body))

	// jwt token creation will be done based on username and password
	
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	tokenString, err2 := token.GenerateJWTToken(userName)


	if err2 != nil {
		log.Error(err2)
		api.LoginErrorHandler(w,err2)
		return
	}
	fmt.Println("JST token sent")

	var Response = struct{
		TokenString string `json:"message"`
	}{
		TokenString: tokenString,
	}

	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(Response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}