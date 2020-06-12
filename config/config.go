package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type Config struct {
	Portal string   `yaml:"portal"`
	Portals []string `yaml:"portals"`
}

func Parse() Config {
	filename, configDir := getPaths()

	_ = os.MkdirAll(configDir, 0700)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("🐣 First run and no config, creating one in: %s\n", filename)
		return writeDummyConfig(filename)
	}

	var result Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&result); err != nil {
		panic("Invalid config format!")
	}

	return result
}

func (c Config) GetSelectedPortal() string {
	return c.Portal
}

func (c Config) SetDefaultPortal(portal string) {
	c.Portal = portal

	filename, _ := getPaths()
	writeConfig(filename, c)
}

func (c Config) GetPortals() []string {
	return c.Portals
}

func writeConfig(filename string, data Config) {
	content, _ := yaml.Marshal(&data)
	if err := ioutil.WriteFile(filename, content, 0600); err != nil {
		fmt.Printf("Could not write config to to file %s\n", filename)
	}
}

func writeDummyConfig(filename string) Config {
	portals := []string{
		"https://siasky.net",
		"https://skyportal.xyz",
		"https://skynet.luxor.tech",
		"https://www.siacdn.com",
	}

	dummy := Config{
		Portal: "https://siasky.net",
		Portals: portals,
	}

	writeConfig(filename, dummy)

	return dummy
}

func getPaths() (string, string) {
	currUser, _ := user.Current()
	confDir := filepath.Join(currUser.HomeDir, "/.config/skyput/")
	return filepath.Join(confDir, "config.yaml"), confDir
}