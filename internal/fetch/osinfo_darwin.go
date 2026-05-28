//go:build darwin

package fetch

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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

	return strings.TrimSpace(name + " " + version)
}

func GetDesktopEnvironment() string {
	return "Aqua"
}

func formatTime(totalSeconds int) string {
	days := totalSeconds / 86400
	hours := (totalSeconds % 86400) / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}
	return fmt.Sprintf("%ds", seconds)
}

func GetUptime() string {
	out, err := exec.Command("sysctl", "-n", "kern.boottime").Output()
	if err != nil {
		return "Unknown Uptime"
	}

	strOut := string(out)
	idx := strings.Index(strOut, "sec = ")
	if idx == -1 {
		return "Unknown Uptime"
	}

	cut := strOut[idx+6:]
	endIdx := strings.Index(cut, ",")
	if endIdx == -1 {
		return "Unknown Uptime"
	}

	bootSecondsStr := strings.TrimSpace(cut[:endIdx])
	bootSeconds, err := strconv.ParseInt(bootSecondsStr, 10, 64)
	if err != nil {
		return "Unknown Uptime"
	}

	currentUnix := time.Now().Unix()
	uptimeSeconds := currentUnix - bootSeconds

	return formatTime(int(uptimeSeconds))
}

func GetPackages() string {
	count := 0

	brewPaths := []string{
		"/opt/homebrew/Cellar",
		"/opt/homebrew/Caskroom",
		"/usr/local/Cellar",
		"/usr/local/Caskroom",
	}

	for _, path := range brewPaths {
		if files, err := os.ReadDir(path); err == nil {
			for _, f := range files {
				if f.IsDir() && !strings.HasPrefix(f.Name(), ".") {
					count++
				}
			}
		}
	}

	if count > 0 {
		return strconv.Itoa(count) + " (brew)"
	}

	return "0"
}
