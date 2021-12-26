package config

import "os"

var (
	MongodbAuth     = os.Getenv("MONGODB_AUTH")
	MongodbDatabase = os.Getenv("MONGODB")
	MongodbUser     = os.Getenv("DBAAS_MONGODB_USER")
	MongodbPassword = os.Getenv("DBAAS_MONGODB_PASSWORD")
	MongodbHost     = os.Getenv("DBAAS_MONGODB_HOSTS")
	MongodbPort     = os.Getenv("DBAAS_MONGODB_PORT")
)
