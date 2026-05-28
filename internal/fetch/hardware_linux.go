//go:build linux

package fetch

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetCPU() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "Unknown CPU (open err)"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(strings.ToLower(line), "model name") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	return "Unknown CPU (not found in cpuinfo)"
}

func parseMeminfo() (totalGB, usedGB float64) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return 0, 0
	}
	defer file.Close()

	var totalMem, availableMem int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal:") {
			fmt.Sscanf(line, "MemTotal: %d kB", &totalMem)
		}
		if strings.HasPrefix(line, "MemAvailable:") {
			fmt.Sscanf(line, "MemAvailable: %d kB", &availableMem)
		}
	}

	totalGB = float64(totalMem) / 1024 / 1024
	usedGB = float64(totalMem-availableMem) / 1024 / 1024
	return totalGB, usedGB
}

func GetTotalRAM() string {
	total, _ := parseMeminfo()
	return fmt.Sprintf("%.2f GB", total)
}

func GetUsedRAM() string {
	_, used := parseMeminfo()
	return fmt.Sprintf("%.2f GB", used)
}
func GetGPU() string {
	output, err := exec.Command("lspci").Output()
	if err != nil {
		return "Uknown GPU"
	}
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "VGA compatible controller") || strings.Contains(line, "3D controller") {
			parts := strings.Split(line, ".0")
			if len(parts) > 1 {
				return strings.Trim(parts[1], `"`)
			}

			if len(line) > 7 {
				return strings.TrimSpace(line[7:])
			}
		}
	}

	return "Unknown GPU"
}
