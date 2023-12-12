package ess

import (
	"os"
)

type environment map[string]string

var Env = environment{
	"host":           os.Getenv("REMOTE_HOST"),
	"port":           os.Getenv("SSH_PORT"),
	"user":           os.Getenv("SSH_USER"),
	"userKeyFile":    os.Getenv("USER_PRIVATE_KEY"),
	"knownHostsFile": os.Getenv("KNOWN_HOSTS"),
}
