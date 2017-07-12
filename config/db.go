package config

type Database struct {
	Host    string `env:"GEN_DBHOST"`
	Port    int    `env:"GEN_DBPORT"`
	SslMode string `env:"GEN_DBSSLMODE"`

	Name string `env:"GEN_DBNAME"`

	Username string `env:"GEN_DBUSER"`
	Password string `env:"GEN_DBPASS"`
}

func DatabaseFromEnv() Database {
	var config Database
	fromEnv(&config)
	return config
}
