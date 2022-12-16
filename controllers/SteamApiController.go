package controllers

import (
	"io/ioutil"
	"log"

	"github.com/Philipp15b/go-steam/v3"
	"github.com/Philipp15b/go-steam/v3/protocol/steamlang"
	"github.com/gin-gonic/gin"
)

type SteamApiController interface {
	SteamAuth(ctx *gin.Context, login, password string)
}

func SteamAuth(context *gin.Context, login, password string) {
	myLoginInfo := new(steam.LogOnDetails)
	myLoginInfo.Username = login
	myLoginInfo.Password = password

	client := steam.NewClient()
	client.Connect()
	for event := range client.Events() {
		switch e := event.(type) {
		case *steam.ConnectedEvent:
			client.Auth.LogOn(myLoginInfo)
		case *steam.MachineAuthUpdateEvent:
			ioutil.WriteFile("sentry", e.Hash, 0666)
		case *steam.LoggedOnEvent:
			client.Social.SetPersonaState(steamlang.EPersonaState_Online)
		case steam.FatalErrorEvent:
			log.Print(e)
		case error:
			log.Print(e)
		}
	}
}
