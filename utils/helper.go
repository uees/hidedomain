package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func IPV4Belong(ip, cidr string) bool {
	ipAddr := strings.Split(ip, `.`)
	if len(ipAddr) < 4 {
		return false
	}
	cidrArr := strings.Split(cidr, `/`)
	if len(cidrArr) < 2 {
		return false
	}
	var tmp = make([]string, 0)
	for key, value := range strings.Split(`255.255.255.0`, `.`) {
		iint, _ := strconv.Atoi(value)

		iint2, _ := strconv.Atoi(ipAddr[key])

		tmp = append(tmp, strconv.Itoa(iint&iint2))
	}
	return strings.Join(tmp, `.`) == cidrArr[0]
}

var BaseDir = func() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(path.Dir(filename))
	}
	return abPath
}

func ShellRun(shell string) (string, error) {
	var outBuf, errBuf bytes.Buffer
	cmd := exec.Command(shell)
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("exec: [%s] err: %v", shell, errBuf.String())
	}
	return strings.TrimSpace(outBuf.String()), nil
}
