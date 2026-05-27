//go:build linux

package fetch

import (
	"bufio"
	"fmt"
	"os"
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
