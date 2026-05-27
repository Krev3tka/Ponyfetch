//go:build darwin

package fetch

import (
	"fmt"
	"os"
	"os/exec"
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

func GetOSName() string {
	out, err := exec.Command("sw_vers", "-productName").Output()
	if err != nil {
		return "macOS"
	}
	return strings.TrimSpace(string(out))
}

func GetOSVersion() string {
	out, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func GetOSMarketingName(version string) string {
	parts := strings.Split(version, ".")
	if len(parts) == 0 {
		return ""
	}

	majorVersion := parts[0]

	switch majorVersion {
	case "26":
		return "Tahoe"
	case "15":
		return "Sequoia"
	case "14":
		return "Sonoma"
	case "13":
		return "Ventura"
	case "12":
		return "Monterey"
	case "11":
		return "Big Sur"
	default:
		return ""
	}
}

func GetOSInfo() string {
	name := GetOSName()
	version := GetOSVersion()
	marketing := GetOSMarketingName(version)

	if marketing != "" {
		return name + " " + marketing + " " + version
	}

	return strings.TrimSpace(string(name + " " + version))
}

func PrintHeader() {
	hostname := GetHostname()

	fmt.Printf("%s @ %s\n", GetUser(), hostname)
	fmt.Printf(strings.Repeat("-", len(GetUser())+len(hostname)+3))
	fmt.Println()
}
