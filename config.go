package main

type Config struct {
	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
}

func Load() Config {

	cfg := Config{}

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "postgres"
	cfg.PostgresDatabase = "market_system"
	cfg.PostgresPassword = "3066586"
	cfg.PostgresPort = "5432"

	return cfg
}
