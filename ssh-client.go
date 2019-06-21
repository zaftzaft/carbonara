package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"golang.org/x/crypto/ssh"
)

type CtrlBuffer struct {
	Ctrl     bool
	Positive io.Writer
	Negative io.Writer
}

func (c *CtrlBuffer) Write(p []byte) (n int, err error) {
	if c.Ctrl {
		return c.Positive.Write(p)
	}
	return c.Negative.Write(p)
}

func SSHClient(h *Host) (string, error) {
	var b string
	w := &bytes.Buffer{}
	cb := &CtrlBuffer{}

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

		//out, err := session.StdoutPipe()
		//if err != nil {
		//	return b, fmt.Errorf("Failed to get stdout: %s", err)
		//}

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

		wait := h.ShellWait
		if h.ShellWait <= 0 {
			wait = 1
		}

		if h.PreBeforeWait <= 0 {
			h.PreBeforeWait = 3
		}

		if h.PreAfterWait <= 0 {
			h.PreAfterWait = 1
		}

		if h.PostBeforeWait <= 0 {
			h.PostBeforeWait = 1
		}

		//cb.Positive = os.Stdout
		cb.Positive = w
		cb.Negative = ioutil.Discard
		session.Stdout = cb

		cb.Ctrl = false

		time.Sleep(time.Duration(h.PreBeforeWait) * time.Second)
		for _, cmd := range h.CommandsPre {
			fmt.Fprintln(in, cmd)
			time.Sleep(time.Duration(wait) * time.Second)
		}

		time.Sleep(time.Duration(h.PreAfterWait) * time.Second)
		cb.Ctrl = true

		for _, cmd := range h.Commands {
			fmt.Fprintln(in, cmd)
			time.Sleep(time.Duration(wait) * time.Second)
		}

		time.Sleep(time.Duration(h.PostBeforeWait) * time.Second)
		cb.Ctrl = false

		for _, cmd := range h.CommandsPost {
			fmt.Fprintln(in, cmd)
			time.Sleep(time.Duration(wait) * time.Second)
		}

		b = w.String()

		session.Wait()
	} else {
		var bu bytes.Buffer
		for _, cmd := range h.Commands {
			session, err := client.NewSession()
			if err != nil {
				return b, fmt.Errorf("Failed to create session: %s", err)
			}
			defer session.Close()

			session.Stdout = &bu

			if err := session.Run(cmd); err != nil {
				return b, fmt.Errorf("Failed to run: %s", err.Error())
			}
		}
		b = bu.String()
	}

	return b, nil
}
