package services

import (
	"fmt"
	"os"
	"saw/models"

	"github.com/Philipp15b/go-steam/v3"
	"github.com/freman/go-steamauth"
)

const (
	BaseURL     = "https://api.steampowered.com"
	SteamApiKey = "BEFD135B9470130381E44DD67D9DA3FD"
)

var (
	AppIds = map[string]uint32{
		"CSGO": 730,
	}
)

type SteamService interface {
	InitAccount(account models.Account) int
}

func InitAccount(account models.Account) int {
	myLoginInfo := new(steam.LogOnDetails)
	myLoginInfo.Username = account.Login
	myLoginInfo.Password = account.Password
	myLoginInfo.AuthCode = GetGuardCode(account.SharedSecret)

	var r int = 0
	client := steam.NewClient()
	client.ConnectionTimeout = 60000

	// server := steam.GetRandomCM()
	// client.ConnectTo(server)
	client.Connect()

	for event := range client.Events() {
		if connected := client.Connected(); !connected {
			return -1
		}

		switch e := event.(type) {
		case *steam.ConnectedEvent:
			client.Auth.LogOn(myLoginInfo)
			fmt.Println("ConnectedEvent")
			fmt.Println("PersonaState:", client.Social.GetPersonaState())
			fmt.Println(client.SteamId(), client.SessionId())
		case *steam.MachineAuthUpdateEvent:
			os.WriteFile("sentry", e.Hash, 0666)
			fmt.Println("MachineAuthUpdateEvent")
		case *steam.LoggedOnEvent:
			fmt.Println("LoggedOnEvent")
			fmt.Println("Avatar:", client.Social.GetAvatar())
			r = client.Social.Friends.Count()
			fmt.Println(r, client.Social.GetPersonaName(), client.Social.Friends)
			return r
		case steam.WebLogOnErrorEvent:
			fmt.Println("WebLogOnErrorEvent:", e.Error())
			return -1
		case steam.FatalErrorEvent:
			fmt.Println("FatalErrorEvent:", e.Error())
			return -1
			// default:
			// 	fmt.Println("default")
			// 	log.Print("default")
		}
	}

	return 0
}

func Check(a models.Account) bool {
	account := steamauth.SteamGuardAccount{}
	if file, err := os.Open("steam_data.json"); err == nil {
		defer file.Close()
		if err = account.Load(file); err != nil {
			fmt.Println("Problem parsing steam_data.json,", err)
		}
	} else {
		fmt.Println(err)
	}

	if account.FullyEnrolled {
		fmt.Println("Already enrolled")
	} else {
		userLogin := steamauth.NewUserLogin(a.Login, a.Password)

		done := false
		for !done {
			res, err := userLogin.DoLogin()

			if err != nil {
				return false
			}

			switch res {
			case steamauth.NeedCaptcha:
				fmt.Printf("Requires captcha, go to %s to get it\n", userLogin.CaptchaURL())
				// userLogin.CaptchaText = getCode()
			case steamauth.Need2FA:
				fmt.Println("Need two factor code")
				userLogin.TwoFactorCode = GetGuardCode(a.SharedSecret)
			case steamauth.NeedEmail:
				fmt.Println("Please check your email for the code")
				// userLogin.EmailCode = getCode()
			case steamauth.LoginOkay:
				done = true
				fmt.Println("Logged in!")
				return true
			default:
				return false
			}
		}
	}
	return false
}
