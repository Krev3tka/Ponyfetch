//go:build linux

package fetch

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetShell() string {
	return os.Getenv("SHELL")
}

func GetUser() string {
	return os.Getenv("USER")
}

func GetHostname() string {
	res, err := os.Hostname()
	if err != nil {
		fmt.Println("Error: failed to get hostname: ", err)
		return "unknown"
	}
	return res
}

func GetOSInfo() string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "Linux"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			parts := strings.Split(line, "=")
			if len(parts) > 1 {
				return strings.Trim(parts[1], `"`)
			}
		}
	}
	return "Linux"
}

func GetDesktopEnvironment() string {
	currentDesktop := os.Getenv("XDG_CURRENT_DESKTOP")
	if currentDesktop != "" {
		return currentDesktop
	}

	currentDesktop = os.Getenv("DESKTOP_SESSION")
	if currentDesktop != "" {
		return currentDesktop
	}

	currentDesktop = os.Getenv("XDG_SESSION_DESKTOP")
	if currentDesktop != "" {
		return currentDesktop
	}

	return "Unknown DE/WM"
}
