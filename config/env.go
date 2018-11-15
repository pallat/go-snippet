package config

import "os"

var ServiceURL = os.Getenv("SERVICE_URL")
