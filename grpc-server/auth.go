package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type AuthResult struct {
	Name         string
	IconURL      string
	AuthProvider string
	AuthUID      string
}

type FirebaseVerifier struct{}

func (v *FirebaseVerifier) VerifyIDToken(idToken string) *AuthResult {
	opt := option.WithCredentialsFile("./service_account.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	claims := token.Claims

	name, ok := claims["name"].(string)
	if !ok {
		log.Fatalf("Type assertion error: %v", claims["name"])
	}

	iconURL, ok := claims["picture"].(string)
	if !ok {
		log.Fatalf("Type assertion error: %v", claims["picture"])
	}

	firebaseField, ok := claims["firebase"].(map[string]interface{})
	if !ok {
		log.Fatalf("Type assertion error: %v", claims["firebase"])
	}

	provider, ok := firebaseField["sign_in_provider"].(string)
	if !ok {
		log.Fatalf("Type assertion error: %v", firebaseField["sign_in_provider"])
	}

	identities, ok := firebaseField["identities"].(map[string]interface{})
	if !ok {
		log.Fatalf("Type assertion error: %v", firebaseField["identities"])
	}

	authProvider := extractAuthProvider(provider)
	authUID := extractAuthUID(identities, provider)

	return &AuthResult{
		Name:         name,
		IconURL:      iconURL,
		AuthProvider: authProvider,
		AuthUID:      authUID,
	}
}

func extractAuthUID(identities map[string]interface{}, provider string) string {
	arr, ok := identities[provider].([]interface{})
	if !ok {
		log.Fatalf("Type assertion error: %v", identities[provider])
	}

	uid, ok := arr[0].(string)
	if !ok {
		log.Fatalf("Type assertion error: %v", arr[0])
	}

	return uid
}

func extractAuthProvider(provider string) string {
	switch provider {
	case "google.com":
		return "google"
	case "facebook.com":
		return "facebook"
	case "github.com":
		return "github"
	case "twitter.com":
		return "twitter"
	default:
		log.Fatalf("Unknown provider: %s", provider)
	}
	return "" // XXX
}
