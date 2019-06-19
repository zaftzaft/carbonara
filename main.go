package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/pmezard/go-difflib/difflib"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

const Version = "0.0.1"

var (
	configfile = kingpin.Arg("configfile", "config file path").Required().String()
)

type RootConfig struct {
	Hosts []Host `yaml:"hosts"`
}

func UpdateLog(filepath string, data string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()
	file.Write(([]byte)(data))

	return nil
}

func FetchHost(host Host) int {
	if host.SSH {
		b, err := SSHClient(&host)
		if err != nil {
			fmt.Println(host.Hostname, err)
			return 1
		}

		result := b.String()

		basedir := filepath.Join("carbonara-log", host.Hostname)
		nowf := time.Now().Format("20060102030405")

		os.MkdirAll(basedir, 0755)

		fmt.Println(host.Hostname, "done", nowf)

		if _, err := os.Stat(filepath.Join(basedir, "_")); err != nil {
			// file not exists
			UpdateLog(filepath.Join(basedir, host.Hostname+"-"+nowf+".txt"), result)
			UpdateLog(filepath.Join(basedir, "_"), result)
		} else {
			final, _ := ioutil.ReadFile(filepath.Join(basedir, "_"))
			diff := difflib.UnifiedDiff{
				A:        difflib.SplitLines(string(final)),
				B:        difflib.SplitLines(result),
				FromFile: host.Hostname,
				ToFile:   nowf,
				Context:  3,
			}

			text, _ := difflib.GetUnifiedDiffString(diff)

			if len(text) > 0 {
				if len(host.WebhookUrl) > 0 {
					SendWebhook(host.WebhookUrl, "```"+text+"```")
				}

				UpdateLog(filepath.Join(basedir, host.Hostname+"-"+nowf+".txt"), result)
				UpdateLog(filepath.Join(basedir, "_"), result)
			}

		}

	}
	return 0
}

func Run() int {
	cbuf, err := ioutil.ReadFile(*configfile)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	var rc RootConfig
	err = yaml.Unmarshal(cbuf, &rc)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	ch := make(chan bool, len(rc.Hosts))
	done := 0

	for _, host := range rc.Hosts {
		host := host
		go func() {
			FetchHost(host)
			ch <- true
		}()
	}

	for {
		select {
		case <-ch:
			done++
			if done == len(rc.Hosts) {
				return 0
			}
		}
	}

	return 0
}

func main() {
	kingpin.Version(Version)
	kingpin.Parse()
	os.Exit(Run())
}