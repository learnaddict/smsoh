package smsoh

import (
	"fmt"
	"net/http"
	"os"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(Middleware{})
	httpcaddyfile.RegisterHandlerDirective("smsoh", parseCaddyfile)
}

// Middleware implements an HTTP handler that writes the
// visitor's IP address to a file or stream.
type Middleware struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.smsoh",
		New: func() caddy.Module { return new(Middleware) },
	}
}

// Provision implements caddy.Provisioner.
func (m *Middleware) Provision(ctx caddy.Context) error {

	return nil
}

// Validate implements caddy.Validator.
func (m *Middleware) Validate() error {

	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	r.ParseForm()

	ud := r.FormValue("ud")
	scts := r.FormValue("scts")
	oa := r.FormValue("oa")
	da := r.FormValue("da")

	if ud != "" && scts != "" && oa != "" && da != "" {
		err := m.mysqlInsert(ud, scts, oa, da)
		if err != nil {
			return err
		}
	} else {

		fmt.Fprintf(w, `<html><body><form action="/">
		<label for="scts">Date/Time:</label><br>
		<input type="text" id="scts" name="scts"><br>
		<label for="ud">Text:</label><br>
		<input type="text" id="ud" name="ud">
		<label for="oa">Sender:</label><br>
		<input type="text" id="oa" name="oa">
		<label for="da">Receipient:</label><br>
		<input type="text" id="da" name="da">
		<input type="submit" value="Submit">
	  </form></body></html>`)
	}
	return nil
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
				if m.Username != "" {
					return d.Err("username path already specified")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				m.Username = d.Val()
			case "password":
				if m.Password != "" {
					return d.Err("password path already specified")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				m.Password = d.Val()
			case "database":
				if m.Database != "" {
					return d.Err("database path already specified")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				m.Database = d.Val()
			default:
				return d.Errf("unrecognized subdirective: %s", d.Val())
			}
		}
	}
	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

func writeSMS(text string) error {
	f, err := os.Create("sms.txt")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	_, err = f.WriteString(text)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ caddyfile.Unmarshaler       = (*Middleware)(nil)
)
