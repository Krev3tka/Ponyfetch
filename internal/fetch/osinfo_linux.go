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
	file, err := os.Open("/proc/uptime")
	if err != nil {
		return "Unknown Uptime"
	}
	defer file.Close()

	scanner, err := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		if len(parts) > 0 {
			secFloat, err := strconv.ParseFloat(parts[0], 64)
			if err != nil {
				return "Unknown Uptime"
			}
			return formatTime(int(secFloat))
		}
	}

	return "Unknown Uptime"
}

func GetPackages() string {
	if files, err := os.ReadDir("/var/lib/pacman/local"); err == nil {
		return strconv.Itoa(len(files))
	}

	if file, err := os.Open("/var/lib/dpkg/status"); err == nil {
		defer file.Close()
		count := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "Package:") {
				count++
			}
		}
		if count > 0 {
			return strconv.Itoa(count)
		}
	}

	if _, err := exec.LookPath("rpm"); err == nil {
		if out, err := exec.Command("rpm", "-qa").Output(); err == nil {
			lines := strings.Split(strings.TrimSpace(string(out)), "\n")
			return strconv.Itoa(len(lines))
		}
	}

	if file, err := os.Open("/var/lib/apk/installed"); err == nil {
		defer file.Close()
		count := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "P:") {
				count++
			}
		}
		if count > 0 {
			return strconv.Itoa(count)
		}
	}

	return "Unknown"
}
