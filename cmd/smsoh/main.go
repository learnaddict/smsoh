package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// plugin standard Caddy modules
	_ "github.com/caddyserver/caddy/v2/modules/standard"

	// plugin additional modules
	_ "github.com/learnaddict/smsoh"
)

func main() {

	caddycmd.Main()
}
