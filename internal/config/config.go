package config

type Config struct {
	Host, Port string
	SecretKey  string
	IPTrusted  []string

	PGZamIO *PostgresConf
	PGZamMe *PostgresConf
}

type PostgresConf struct {
	Host, Port, Password, Username, DBName, SSLMode string
}
