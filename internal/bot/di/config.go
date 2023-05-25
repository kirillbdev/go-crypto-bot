package di

type Config struct {
	dbHost string
	dbUser string
	dbPass string
	dbName string
}

func NewConfig(dbHost string, dbUser string, dbPass string, dbName string) Config {
	return Config{dbHost: dbHost, dbUser: dbUser, dbPass: dbPass, dbName: dbName}
}
