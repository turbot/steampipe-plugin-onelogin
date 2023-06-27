package main

import (
	"github.com/turbot/steampipe-plugin-onelogin/onelogin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: onelogin.Plugin})
}
