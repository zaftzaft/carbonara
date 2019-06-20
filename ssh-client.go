package main

import (
	"bytes"
	"fmt"
	//"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func SSHClient(h *Host) (bytes.Buffer, error) {
	var b bytes.Buffer

	config := ssh.Config{}
	config.SetDefaults()

	clientConfig := &ssh.ClientConfig{
		Config: config,
		User:   h.Username,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
				var ans []string
				for _, q := range questions {
					if q == "Password: " {
						ans = append(ans, h.Password)
					}
				}
				return ans, nil
			}),
			ssh.Password(h.Password),
		},
		Timeout: time.Second * 10,

		// FIXME
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	client, err := ssh.Dial("tcp", h.Address+":"+h.SSHPort(), clientConfig)
	if err != nil {
		return b, fmt.Errorf("Failed to dial: %s", err)
	}

	if h.Shell {
		session, err := client.NewSession()
		if err != nil {
			return b, fmt.Errorf("Failed to create session: %s", err)
		}
		defer session.Close()

		in, err := session.StdinPipe()
		if err != nil {
			return b, fmt.Errorf("Failed to get stdin: %s", err)
		}

		modes := ssh.TerminalModes{
			ssh.ECHO:          0,     // disable echoing
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		}
		// Request pseudo terminal
		if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
			return b, fmt.Errorf("request for pseudo terminal failed: %s", err)
		}

		if err := session.Shell(); err != nil {
			return b, fmt.Errorf("Failed to start shell: %s", err)
		}

		//session.Stdout = os.Stdout
		//session.Stderr = os.Stderr
		session.Stdout = &b

		wait := h.ShellWait
		if h.ShellWait <= 0 {
			wait = 1
		}

		for _, cmd := range h.Commands {
			time.Sleep(time.Duration(wait) * time.Second)
			fmt.Fprintln(in, cmd+"\n")
		}

		//fmt.Fprint(in, "enable\n")
		//fmt.Fprint(in, "")

		session.Wait()

	} else {
		for _, cmd := range h.Commands {
			session, err := client.NewSession()
			if err != nil {
				return b, fmt.Errorf("Failed to create session: %s", err)
			}
			defer session.Close()

			session.Stdout = &b

			if err := session.Run(cmd); err != nil {
				return b, fmt.Errorf("Failed to run: %s", err.Error())
			}
		}
	}
	//if err := session.Run("/usr/bin/whoami"); err != nil {

	return b, nil
}
