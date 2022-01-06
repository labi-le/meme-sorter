package structures

type Config struct {
	Dsn      string `toml:"dsn"`
	LogLevel string `toml:"log_level"`
}
