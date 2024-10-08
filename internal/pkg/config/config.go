package config

type Environment struct {
	Mongo Mongo
	Redis Redis
}

type Mongo struct {
	Connection string
	Database   string
}

type Redis struct {
	Host     string
	Port     int
	Username string
	Password string
	Db       int
}
