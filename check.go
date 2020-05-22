package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
)

func notExist(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return true
	}

	return false
}

func getUid(userName string) int {
	u, err := user.Lookup(userName)

	if err != nil {
		panic(fmt.Sprintf("[ERROR] user %s not found", userName))
	}

	uid, _ := strconv.Atoi(u.Uid)
	return uid
}

func getGid(groupName string) int {
	g, err := user.LookupGroup(groupName)

	if err != nil {
		panic(fmt.Sprintf("[ERROR] group %s not found", groupName))
	}

	gid, _ := strconv.Atoi(g.Gid)
	return gid
}

func tcpCheck(ports []int) {
	for _, p := range ports {
		if p > 65534 {
			continue
		}
		cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -i tcp:%d | grep LISTEN", p))
		result, _ := cmd.Output()
		if string(result) != "" {
			panic(fmt.Sprintf("[ERROR] tcp %d has been used", p))
		}
	}
}

func udpCheck(ports []int) {
	for _, p := range ports {
		if p > 65534 {
			continue
		}
		cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -i udp:%d", p))
		result, _ := cmd.Output()
		if string(result) != "" {
			panic(fmt.Sprintf("[ERROR] ucp %d has been used", p))
		}
	}
}

func preStartCheck(config *Config) {
	f := config.File
	file := f.WorkDir + "/" + f.Name
	mode := f.Mode

	if notExist(file) {
		panic(fmt.Sprintf("[ERROR] %s does not exist", file))
	}

	uid := getUid(f.Owner)
	gid := getGid(f.Group)

	err := os.Chown(file, uid, gid)
	if err != nil {
		panic(err)
	}

	err = os.Chmod(file, mode)
	if err != nil {
		panic(err)
	}

	for _, v := range config.WriteEnable {
		if notExist(v) {
			err = os.MkdirAll(v, 0755)
			if err != nil {
				panic(err)
			}
		}

		err = os.Chown(v, uid, gid)
		if err != nil {
			panic(err)
		}

		err = os.Chmod(v, 0755)
		if err != nil {
			panic(err)
		}
	}

	tcp := config.Net.Tcp
	if len(tcp) > 0 {
		tcpCheck(tcp)
	}

	udp := config.Net.Udp
	if len(udp) > 0 {
		udpCheck(udp)
	}
}
