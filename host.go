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
	SSHPortNum     uint16   `yaml:"ssh_port"`
	SSHKeyPath     string   `yaml:"ssh_key_path"`
	Telnet         bool     `yaml:"telnet"`
	TelnetPortNum  uint16   `yaml:"telnet_port"`
	WebhookUrl     string   `yaml:"webhook"`
}

func (h *Host) SSHPort() string {
	if h.SSHPortNum > 0 {
		return strconv.Itoa(int(h.SSHPortNum))
	}

	return "22"
}

func (h *Host) TelnetPort() string {
	if h.TelnetPortNum > 0 {
		return strconv.Itoa(int(h.TelnetPortNum))
	}

	return "telnet"
}
