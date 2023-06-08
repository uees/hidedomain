package services

import (
	"fmt"
	"runtime"

	"github.com/uees/hidedomain/utils"
)

func DenyDomain(domain string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	// iptables -I INPUT --dport 443 -m string --string "domain" --algo bm -j DROP
	cmdStr := fmt.Sprintf("iptables -I INPUT --dport 443 -m string --string \"%s\" --algo bm -j DROP", domain)
	return utils.ShellRun(cmdStr)
}

func AllowDomain(domain string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	cmdStr := fmt.Sprintf("iptables -D INPUT --dport 443 -m string --string \"%s\" --algo bm -j DROP", domain)
	return utils.ShellRun(cmdStr)
}

func AllowIP(ip string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	cmdStr := fmt.Sprintf("iptables -I INPUT -s %s --dport 443 -j ACCEPT", ip)
	return utils.ShellRun(cmdStr)
}

func RemoveIP(ip string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	cmdStr := fmt.Sprintf("iptables -D INPUT -s %s --dport 443 -j ACCEPT", ip)
	return utils.ShellRun(cmdStr)
}

func SaveRules() (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	return utils.ShellRun("iptables-save > /etc/iptables/rules.v4")
}

func RestoreRules() (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	return utils.ShellRun("iptables-restore < /etc/iptables/rules.v4")
}
