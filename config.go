package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Service struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Version string `yaml:"version"`
}

type File struct {
	Owner   string            `yaml:"owner"`
	Group   string            `yaml:"group"`
	Mode    os.FileMode       `yaml:"mode"`
	Name    string            `yaml:"name"`
	WorkDir string            `yaml:"work_dir"`
	Backup  map[string]string `yaml:"backup"`
}

type Environment struct {
	Get []string `yaml:"get"`
	Set []string `yaml:"set"`
}

type Run struct {
	User        string              `yaml:"user"`
	Group       string              `yaml:"group"`
	Environment *Environment        `yaml:"environment"`
	Command     map[string][]string `yaml:"command"`
	Log         []string            `yaml:"log"`
	Pid         string              `yaml:"pid"`
}

type Net struct {
	Bind   string   `yaml:"bind"`
	Tcp    []int    `yaml:"tcp"`
	Udp    []int    `yaml:"udp"`
	Socket []string `yaml:"socket"`
}

type Config struct {
	Service     *Service `yaml:"service"`
	File        *File    `yaml:"file"`
	WriteEnable []string `yaml:"write_enable"`
	Run         *Run     `yaml:"run"`
	Net         *Net     `yaml:"net"`
}

func (c *Config) Read(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	err = yaml.Unmarshal(data, c)
	return err
}
