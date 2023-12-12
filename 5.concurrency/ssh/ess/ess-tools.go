package ess

import (
	"fmt"
	"os"
)

type environment map[string]string

var Env = environment{
	"host": os.Getenv("REMOTE_HOST"),
	"port": os.Getenv("SSH_PORT"),
}

func (e environment) Check() error {
	wrongEnv := false
	var errString string
	var err error = nil
	for k, v := range e {
		if len(v) == 0 {
			errString += fmt.Sprintf("ERROR: Environment variable for %#v is missing\n", k)
			wrongEnv = true
		}
	}
	if wrongEnv {
		err = fmt.Errorf(errString)
	}

	return err
}

func DoSSH(h, p string) (chan string, chan string, error) {
	in := make(chan string)
	out := make(chan string)
	err := fmt.Errorf("i don't want to do this")
	return in, out, err
}
