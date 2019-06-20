package main

import (
	"bytes"
	"fmt"
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
	//if err := session.Run("/usr/bin/whoami"); err != nil {

	return b, nil
}
