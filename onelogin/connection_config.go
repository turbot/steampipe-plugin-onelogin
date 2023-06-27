package onelogin

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type oneloginConfig struct {
	ClientId     *string `cty:"client_id"`
	ClientSecret *string `cty:"client_secret"`
	Subdomain    *string `cty:"subdomain"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"client_id": {
		Type: schema.TypeString,
	},
	"client_secret": {
		Type: schema.TypeString,
	},
	"subdomain": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &oneloginConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) oneloginConfig {
	if connection == nil || connection.Config == nil {
		return oneloginConfig{}
	}
	config, _ := connection.Config.(oneloginConfig)
	return config
}
