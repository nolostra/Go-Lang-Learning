package tools

import (
	"time"
)


type mockDB struct{} 

var mockLoginDetails = map[string]LoginDetails{
	"sai":{
		AuthToken: "1400",
		Username: "Saish",
	},
	"Tahil":{
		AuthToken: "140033",
		Username: "Tahil",
	},
	
}

var mockCoinDetails = map[string]CoinDetails{
	"sai":{
		Coins: 1400,
		Username: "Saish",
	},
	"Tahil":{
		Coins: 14040,
		Username: "Tahil",
	},
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData,ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData

}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData,ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData

}