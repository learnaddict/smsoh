package smsoh

import (
	"github.com/caddyserver/caddy/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for d.NextBlock(0) {
			switch d.Val() {
			case "username":
				if m.MySQL.Username != "" {
					return d.Err("username path already specified")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				m.MySQL.Username = d.Val()
			case "password":
				if m.MySQL.Password != "" {
					return d.Err("password path already specified")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				m.MySQL.Password = d.Val()
			case "database":
				if m.MySQL.Database != "" {
					return d.Err("database path already specified")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				m.MySQL.Database = d.Val()
			default:
				return d.Errf("unrecognized subdirective: %s", d.Val())
			}
		}
	}
	return nil
}
