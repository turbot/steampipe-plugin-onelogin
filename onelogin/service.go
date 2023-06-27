package onelogin

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/onelogin/onelogin-go-sdk/pkg/onelogin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connectOnelogin(ctx context.Context, d *plugin.QueryData) (*onelogin.OneloginSDK, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "onelogin"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*onelogin.OneloginSDK), nil
	}

	oneloginConfig := GetConfig(d.Connection)

	// Default to using env vars (#2)
	clientId, ok := os.LookupEnv("ONELOGIN_CLIENT_ID")
	if !ok {
		clientId = *oneloginConfig.ClientId
		os.Setenv("ONELOGIN_CLIENT_ID", clientId)
	}
	clientSecret, ok := os.LookupEnv("ONELOGIN_CLIENT_SECRET")
	if !ok {
		clientSecret = *oneloginConfig.ClientSecret
		os.Setenv("ONELOGIN_CLIENT_SECRET", clientSecret)
	}
	subdomain, ok := os.LookupEnv("ONELOGIN_SUBDOMAIN")
	if !ok {
		subdomain = *oneloginConfig.Subdomain
		os.Setenv("ONELOGIN_SUBDOMAIN", subdomain)
	}

	// Error if the minimum config is not set
	if clientId == "" {
		return nil, errors.New("client_id must be configured")
	}
	if clientSecret == "" {
		return nil, errors.New("client_secret must be configured")
	}
	if subdomain == "" {
		return nil, errors.New("subdomain must be configured")
	}

	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println("Unable to initialize client:", err)
		return nil, err
	}

	if client != nil {
		d.ConnectionManager.Cache.Set(cacheKey, client)
	}

	return client, nil
}
