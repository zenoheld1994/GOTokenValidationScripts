package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func createCustomToken(app *firebase.App, uid string) (string, error) {
	ctx := context.Background()
	client, err := app.Auth(ctx)
	if err != nil {
		return "", err
	}

	// Additional claims that will be included in the token
	claims := map[string]interface{}{
		"premiumAccount": true,
	}

	token, err := client.CustomTokenWithClaims(ctx, uid, claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func main() {
	//change your key
	opt := option.WithCredentialsFile("KEYJSON")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	// UID of the user for whom the token is created
	print("this is run")
	//UID OF YOUR USER
	uid := "UID"

	token, err := createCustomToken(app, uid)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Custom Token: %s\n", token)
}


