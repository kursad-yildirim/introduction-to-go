package ess

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"golang.org/x/crypto/ssh"
)

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

func DoSSH(h, p string) (chan<- string, <-chan string, error) {

	signer, err := prepareSigner()
	if err != nil {
		return nil, nil, err
	}
	hostKey, err := prepareHostKey(h)
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
	session, err := createClientSession(config, h, p)
	if err != nil {
		return nil, nil, err
	}
	w, err := session.StdinPipe()
	if err != nil {
		return nil, nil, err
	}
	r, err := session.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	in, out := shell(w, r)

	return in, out, nil
}

func prepareSigner() (ssh.Signer, error) {
	key, err := os.ReadFile(Env["userKeyFile"])
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	return signer, nil
}

func prepareHostKey(h string) (ssh.PublicKey, error) {
	file, err := os.Open(Env["knownHostsFile"])
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %#v", Env["knownHostsFile"])
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], h) {
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, fmt.Errorf("parse error with %#v for %#v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		return nil, fmt.Errorf("host key not found for %#v in %#v", h, Env["knownHostsFile"])
	}
	return hostKey, nil
}

func createClientSession(config *ssh.ClientConfig, h, p string) (*ssh.Session, error) {
	client, err := ssh.Dial("tcp", h+":"+p, config)
	if err != nil {
		return nil, err
	}
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, err
	}

	return session, nil
}

func shell(w io.Writer, r io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 1)
	out := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1) //for the shell itself
	go func() {
		fmt.Println("hello 1")
		for cmd := range in {
			fmt.Println(">>>>>>>>>>>>>>", cmd)
			wg.Add(1)
			w.Write([]byte(cmd + "\n"))
			wg.Wait()
		}
		fmt.Println("bye 1")
	}()
	go func() {
		fmt.Println("hello 2")
		var (
			buf [65 * 1024]byte
			t   int
		)
		for {
			n, err := r.Read(buf[t:])
			if err != nil {
				fmt.Println("hello error")
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
		fmt.Println("bye 2")
	}()
	return in, out

}
