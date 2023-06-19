package utils

import (
	"fmt"
	"runtime"
)

func DenyDomain(domain string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	// iptables -I INPUT -p TCP --dport 443 -m string --string "domain" --algo bm -j DROP
	// cmdStr := fmt.Sprintf("iptables -I INPUT -p TCP --dport 443 -m string --string \"%s\" --algo bm -j DROP", domain)
	return ShellRun("iptables", "-I INPUT", "-p TCP", "--dport 443", "-m string",
		fmt.Sprintf("--string \"%s\"", domain), "--algo bm", "-j DROP")
}

func AllowDomain(domain string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	// cmdStr := fmt.Sprintf("iptables -D INPUT -p TCP --dport 443 -m string --string \"%s\" --algo bm -j DROP", domain)
	return ShellRun("iptables", "-D INPUT", "-p TCP", "--dport 443", "-m string",
		fmt.Sprintf("--string \"%s\"", domain), "--algo bm", "-j DROP")
}

func AllowIP(domain string, ip string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	// cmdStr := fmt.Sprintf("iptables -I INPUT -s %s -p TCP --dport 443 -m string --string \"%s\" --algo bm -j ACCEPT", ip, domain)
	return ShellRun("iptables", "-I INPUT", fmt.Sprintf("-s %s", ip), "-p TCP", "--dport 443", "-m string",
		fmt.Sprintf("--string \"%s\"", domain), "--algo bm", "-j ACCEPT")
}

func RemoveIP(domain string, ip string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	// cmdStr := fmt.Sprintf("iptables -D INPUT -s %s -p TCP --dport 443 -m string --string \"%s\" --algo bm -j ACCEPT", ip, domain)
	return ShellRun("iptables", "-D INPUT", fmt.Sprintf("-s %s", ip), "-p TCP", "--dport 443", "-m string",
		fmt.Sprintf("--string \"%s\"", domain), "--algo bm", "-j ACCEPT")
}

func SaveRules() (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	//return ShellRun("iptables-save > /etc/iptables/rules.v4")
	return ShellRun("netfilter-persistent", "save")
}

func RestoreRules() (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	//return ShellRun("iptables-restore < /etc/iptables/rules.v4")
	return ShellRun("iptables-persistent", "reload")
}
