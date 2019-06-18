package main

import (
	"strconv"
)

type Host struct {
	Hostname   string   `yaml:"hostname"`
	Address    string   `yaml:"addr"`
	Username   string   `yaml:"username"`
	Password   string   `yaml:"password"`
	Commands   []string `yaml:"cmds"`
	SSH        bool     `yaml:"ssh"`
	Telnet     bool     `yaml:"telnet"`
	SSHPortNum uint16   `yaml:"ssh_port"`
	WebhookUrl string   `yaml:"webhook"`
}

func (h *Host) SSHPort() string {
	if h.SSHPortNum > 0 {
		return strconv.Itoa(int(h.SSHPortNum))
	}

	return "22"
}
