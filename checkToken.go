package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func verifyToken(app *firebase.App, idToken string) (*auth.Token, error) {
	ctx := context.Background()
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func main() {
	//change to yourpersonalkey.json it must be on the same level as these scripts
	opt := option.WithCredentialsFile("KEY")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	// This is the token you received from the client it is a very long one
	receivedToken := "TOKEN"
	print("we got the token")
	token, err := verifyToken(app, receivedToken)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Verified Token: %v\n", token)
}