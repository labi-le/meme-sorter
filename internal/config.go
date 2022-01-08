package internal

type Config struct {
	Dsn      string `toml:"dsn"`
	LogLevel string `toml:"log_level"`
	Addr     string `toml:"addr"`

	DB *DB
}
