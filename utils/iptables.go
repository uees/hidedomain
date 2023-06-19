package utils

import (
	"fmt"
	"log"
	"runtime"
)

func DenyDomain(domain string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	// iptables -I INPUT -p TCP --dport 443 -m string --string "domain" --algo bm -j DROP
	cmdStr := fmt.Sprintf("iptables -I INPUT -p TCP --dport 443 -m string --string \"%s\" --algo bm -j DROP", domain)
	log.Println(cmdStr)
	return ShellRun(cmdStr)
}

func AllowDomain(domain string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	cmdStr := fmt.Sprintf("iptables -D INPUT -p TCP --dport 443 -m string --string \"%s\" --algo bm -j DROP", domain)
	log.Println(cmdStr)
	return ShellRun(cmdStr)
}

func AllowIP(domain string, ip string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	cmdStr := fmt.Sprintf("iptables -I INPUT -s %s -p TCP --dport 443 -m string --string \"%s\" --algo bm -j ACCEPT", ip, domain)
	log.Println(cmdStr)
	return ShellRun(cmdStr)
}

func RemoveIP(domain string, ip string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	cmdStr := fmt.Sprintf("iptables -D INPUT -s %s -p TCP --dport 443 -m string --string \"%s\" --algo bm -j ACCEPT", ip, domain)
	log.Println(cmdStr)
	return ShellRun(cmdStr)
}

func SaveRules() (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	//return ShellRun("iptables-save > /etc/iptables/rules.v4")
	return ShellRun("netfilter-persistent save")
}

func RestoreRules() (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("no support os")
	}
	//return ShellRun("iptables-restore < /etc/iptables/rules.v4")
	return ShellRun("iptables-persistent reload")
}
