package middleware

import (
	"errors"
	"net/http"

	"github.com/nolostra/goproject/api"
	"github.com/nolostra/goproject/internal/tools"
	log "github.com/sirupsen/logrus"

)

var UnAuthorizedError = errors.New("Invalid Username or token! ")

func Authorisation(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
		var userName string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		var err error

		if userName == "" || token == ""{
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return 
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(userName)

		if(loginDetails == nil || (token != (*loginDetails).AuthToken)){
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		next.ServeHTTP(w,r)
	})
}