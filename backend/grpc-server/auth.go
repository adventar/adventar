package main

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"golang.org/x/xerrors"
	"google.golang.org/api/option"
)

// AuthResult represents result of authentication.
type AuthResult struct {
	Name         string
	IconURL      string
	AuthProvider string
	AuthUID      string
}

type firebaseVerifier struct{}

func (v *firebaseVerifier) VerifyIDToken(idToken string) (*AuthResult, error) {
	json := os.Getenv("FIREBASE_CREDENTIAL_JSON")
	if json == "" {
		return nil, fmt.Errorf("FIREBASE_CREDENTIAL_JSON is empty")
	}
	opt := option.WithCredentialsJSON([]byte(json))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, xerrors.Errorf("Failed to initialize firebase app: %w", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("Failed to get auth client: %w", err)
	}

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, xerrors.Errorf("Failed to verify token: %w", err)
	}

	claims := token.Claims

	name, ok := claims["name"].(string)
	if !ok {
		return nil, xerrors.Errorf("Failed to assert type: %v", claims["name"])
	}

	iconURL, ok := claims["picture"].(string)
	if !ok {
		return nil, xerrors.Errorf("Failed to assert type: %v", claims["picture"])
	}

	firebaseField, ok := claims["firebase"].(map[string]interface{})
	if !ok {
		return nil, xerrors.Errorf("Failed to assert type: %v", claims["firebase"])
	}

	provider, ok := firebaseField["sign_in_provider"].(string)
	if !ok {
		return nil, xerrors.Errorf("Failed to assert type: %v", firebaseField["sign_in_provider"])
	}

	identities, ok := firebaseField["identities"].(map[string]interface{})
	if !ok {
		return nil, xerrors.Errorf("Failed to assert type: %v", firebaseField["identities"])
	}

	authProvider := extractAuthProvider(provider)
	authUID := extractAuthUID(identities, provider)

	return &AuthResult{
		Name:         name,
		IconURL:      iconURL,
		AuthProvider: authProvider,
		AuthUID:      authUID,
	}, nil
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
