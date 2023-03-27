package util

import (
	"context"
	"fmt"
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

type FirebaseVerifier struct{}

func (v *FirebaseVerifier) VerifyIDToken(idToken string) (*AuthResult, error) {
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

	var name string
	name, ok := claims["name"].(string)
	if !ok {
		name = "No Name"
	}

	var iconURL string
	iconURL, ok = claims["picture"].(string)
	if !ok {
		iconURL = ""
	}

	firebaseField, ok := claims["firebase"].(map[string]interface{})
	if !ok {
		return nil, xerrors.Errorf("Failed to assert type [firebase]: %v", claims)
	}

	provider, ok := firebaseField["sign_in_provider"].(string)
	if !ok {
		return nil, xerrors.Errorf("Failed to assert type: %v", firebaseField["sign_in_provider"])
	}

	identities, ok := firebaseField["identities"].(map[string]interface{})
	if !ok {
		return nil, xerrors.Errorf("Failed to assert type: %v", firebaseField["identities"])
	}

	authProvider, err := extractAuthProvider(provider)
	if err != nil {
		return nil, err
	}

	authUID, err := extractAuthUID(identities, provider)
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		Name:         name,
		IconURL:      iconURL,
		AuthProvider: authProvider,
		AuthUID:      authUID,
	}, nil
}

func extractAuthUID(identities map[string]interface{}, provider string) (string, error) {
	arr, ok := identities[provider].([]interface{})
	if !ok {
		return "", xerrors.Errorf("Type assertion error: %v", identities[provider])
	}

	uid, ok := arr[0].(string)
	if !ok {
		return "", xerrors.Errorf("Type assertion error: %v", arr[0])
	}

	return uid, nil
}

func extractAuthProvider(provider string) (string, error) {
	switch provider {
	case "google.com":
		return "google", nil
	case "facebook.com":
		return "facebook", nil
	case "github.com":
		return "github", nil
	case "twitter.com":
		return "twitter", nil
	default:
		return "", xerrors.Errorf("Unknown provider: %s", provider)
	}
}
