//package main
//
//import (
//	"context"
//	"fmt"
//	"log"
//
//	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
//	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/auth"
//)
//
//func main() {
//	config := dropbox.Config{
//		Token: "<your-app-access-token>",
//	}
//
//	dbx := dropbox.New(config)
//
//	authResponse, err := auth.NewTokenCreateV2API(dbx).TokenCreateV2(
//		context.Background(),
//		&auth.TokenCreateArg{},
//	)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("Dropbox Access Token: %s\n", authResponse.GetAccessToken())
//}

package main

import (
	"fmt"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/auth"
	"log"
)

func main() {
	config := dropbox.Config{
		Token: "sl.BZTbFX-33knXY4pmJR7xePoU8z-9bRzQBK_f6xPxRs37wFfTB52JtvlratYOdhw3lgquWjDxdxnxzPeAL39BgSaGx5yHGuFSfl_cmRqnPDWg6V6phfCojM96jEFRkyZ24kwdxpQ",
	}

	dbx := auth.New(config)

	err := dbx.TokenRevoke()
	if err != nil {
		log.Fatal(err)
	}
	token, err := dbx.TokenFromOauth1(&auth.TokenFromOAuth1Arg{Oauth1Token: config.Token, Oauth1TokenSecret: ""})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
}
