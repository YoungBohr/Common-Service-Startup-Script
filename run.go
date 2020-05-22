package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func startup(config *Config, instance string) {
	run := config.Run
	command := run.Command[instance]

	cmd := exec.Command(command[0], command[1:len(command)]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = os.Environ()
	getEnv := run.Environment.Get
	if len(getEnv) > 0 {
		for _, v := range getEnv {
			cmd.Env = append(cmd.Env, os.Getenv(v))
		}
	}

	setEnv := run.Environment.Set
	if len(setEnv) > 0 {
		for _, v := range setEnv {
			cmd.Env = append(cmd.Env, v)
		}
	}

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	pid := strconv.Itoa(cmd.Process.Pid)
	pidFile := run.Pid
	err = ioutil.WriteFile(pidFile, []byte(pid), 0644)
	if err != nil {
		panic(err)
	}
}
