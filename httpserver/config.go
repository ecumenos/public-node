package httpserver

import "time"

// Config is a configuration of HTTP server.
type Config struct {
	Address        string        `default:"0.0.0.0:3000"`
	HandlerTimeout time.Duration `default:"30s"`
	ReadTimeout    time.Duration `default:"15s"`
	WriteTimeout   time.Duration `default:"15s"`
	IdleTimeout    time.Duration `default:"15s"`
}
