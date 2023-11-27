package intro

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"golang.org/x/crypto/ssh"
)

type environment map[string]string

var Env = environment{
	"node":        os.Getenv("NODE"),
	"user":        os.Getenv("SSH_USER"),
	"known_hosts": os.Getenv("KNOWN_HOSTS"),
	"userKey":     os.Getenv("USER_KEY"),
	"ssh_port":    os.Getenv("SSH_PORT"),
}

func (e environment) Check() {
	wrongEnv := false
	for k, v := range e {
		if len(v) == 0 {
			fmt.Printf("Environment variable for %v is missing!\n", k)
			wrongEnv = true
		}
	}
	if wrongEnv {
		os.Exit(1)
	}
	//fmt.Printf("Environment is available as: %#v\n", Env)
}

func DoSSH(h string, p string) (chan<- string, <-chan string, error) {
	var client *ssh.Client

	key, err := os.ReadFile(Env["userKey"])
	if err != nil {
		return nil, nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, nil, err
	}
	hostKey, err := getHostKey(h)
	if err != nil {
		return nil, nil, err
	}
	config := &ssh.ClientConfig{
		User: Env["user"],
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}
	client, err = ssh.Dial("tcp", h+":"+p, config)
	if err != nil {
		fmt.Println("error dial")
		return nil, nil, err
	}
	s, err := client.NewSession()

	//////////////////
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := s.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, nil, err
	}

	w, err := s.StdinPipe()
	if err != nil {
		return nil, nil, err
	}
	r, err := s.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	in, out := shell(w, r)
	if err := s.Start("/bin/sh"); err != nil {
		return nil, nil, err
	}

	return in, out, nil
}

func getHostKey(host string) (ssh.PublicKey, error) {
	file, err := os.Open(Env["known_hosts"])
	if err != nil {
		return nil, fmt.Errorf("cannot open known hosts file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, fmt.Errorf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		return nil, fmt.Errorf("no hostkey for %s", host)
	}
	return hostKey, nil
}

func shell(w io.Writer, r io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 1)
	out := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1) //for the shell itself
	go func() {
		for cmd := range in {
			wg.Add(1)
			w.Write([]byte(cmd + "\n"))
			wg.Wait()
		}
	}()
	go func() {
		var (
			buf [65 * 1024]byte
			t   int
		)
		for {
			n, err := r.Read(buf[t:])
			if err != nil {
				close(in)
				close(out)
				return
			}
			t += n
			if buf[t-2] == '$' { //assuming the $PS1 == 'sh-4.3$ '
				out <- string(buf[:t])
				t = 0
				wg.Done()
			}
		}
	}()
	return in, out
}
