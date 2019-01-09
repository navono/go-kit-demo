package cmd

import (
	"context"
	"flag"
	"fmt"

	"navono/go-kit-demo/pkg/transport/rest"
)

// Config for Server
type Config struct {
	HTTPPort string
}

// RunServer a service instance
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	// flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	// flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	// flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	// flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	// flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	// flag.IntVar(&cfg.LogLevel, "log-level", 0, "Global log level")
	// flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "",
	// 	"Print time format for logger e.g. 2006-01-02T15:04:05Z07:00")
	flag.Parse()

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	return rest.RunServer(ctx, cfg.HTTPPort)
}
