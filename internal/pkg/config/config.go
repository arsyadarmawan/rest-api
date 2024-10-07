package config

type Environment struct {
	Mongo Mongo
}

type Mongo struct {
	Connection string
	Database   string
}
