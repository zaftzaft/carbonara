package main

import (
	"strconv"
)

type Host struct {
	Hostname       string   `yaml:"hostname"`
	Address        string   `yaml:"addr"`
	Username       string   `yaml:"username"`
	Password       string   `yaml:"password"`
	Shell          bool     `yaml:"shell"`
	ShellWait      int      `yaml:"shell_wait"`
	Commands       []string `yaml:"cmds"`
	CommandsPre    []string `yaml:"cmds_pre"`
	CommandsPost   []string `yaml:"cmds_post"`
	PreBeforeWait  int      `yaml:"pre_before_wait"`
	PreAfterWait   int      `yaml:"pre_after_wait"`
	PostBeforeWait int      `yaml:"post_before_wait"`
	SSH            bool     `yaml:"ssh"`
	Telnet         bool     `yaml:"telnet"`
	SSHPortNum     uint16   `yaml:"ssh_port"`
	WebhookUrl     string   `yaml:"webhook"`
}

func (h *Host) SSHPort() string {
	if h.SSHPortNum > 0 {
		return strconv.Itoa(int(h.SSHPortNum))
	}

	return "22"
}
