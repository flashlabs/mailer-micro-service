package server

import "time"

const (
	ReadTimeout       = 1 * time.Second
	WriteTimeout      = 1 * time.Second
	IdleTimeout       = 30 * time.Second
	ReadHeaderTimeout = 2 * time.Second
	Port              = 8080
)
