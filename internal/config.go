package internal

type Config struct {
	Dsn      string `toml:"dsn"`
	LogLevel string `toml:"log_level"`
	Addr     string `toml:"addr"`
	ItemsDir string `toml:"items_dir"`

	DB *DB
}
