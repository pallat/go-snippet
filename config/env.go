package config

import "os"

// PORT is port to serve like 8080
var PORT = os.Getenv("PORT")
