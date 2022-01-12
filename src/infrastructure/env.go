package config

import "os"

var (
	MongodbAuth     = os.Getenv("MONGODB_AUTH")
	MongodbDatabase = os.Getenv("Cluster0MovieWorkNow")
	MongodbUser     = os.Getenv("elisacds")
	MongodbPassword = os.Getenv("elisacds")
	MongodbHost     = os.Getenv("DBAAS_MONGODB_HOSTS")
	MongodbPort     = os.Getenv("DBAAS_MONGODB_PORT")
)
