package utils

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/cloudflare/cloudflare-go"
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

func ShellRun(name string, arg ...string) (string, error) {
	var outBuf, errBuf bytes.Buffer
	cmd := exec.Command(name, arg...)
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("exec: [%s] err: %v", cmd.String(), errBuf.String())
	}
	return strings.TrimSpace(outBuf.String()), nil
}

func InitCfApi(token string) (*cloudflare.API, context.Context) {
	ctx := context.Background()
	api, err := cloudflare.NewWithAPIToken(token)
	if err != nil {
		log.Fatal(err)
	}

	return api, ctx
}
