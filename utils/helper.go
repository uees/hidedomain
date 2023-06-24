package utils

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/cloudflare/cloudflare-go"
)

func GenIPV4Cidr(ip string, ones int) string {
	ipv4 := net.ParseIP(ip).To4()
	if ipv4 != nil {
		mask := net.CIDRMask(ones, 32)
		return ipv4.Mask(mask).String() + "/" + strconv.Itoa(ones)
	}

	log.Fatal("not support ipv6")
	return ""
}

func IsIPv4(ip string) bool {
	return net.ParseIP(ip).To4() != nil
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
