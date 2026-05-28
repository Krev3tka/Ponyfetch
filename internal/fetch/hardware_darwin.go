//go:build darwin

package fetch

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func GetCPU() string {
	output, err := exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output()
	if err != nil {
		fmt.Println("Error: failed to parse CPU name: ", err)
	}

	return strings.TrimSpace(string(output))
}

func GetTotalRAM() string {
	output, err := exec.Command("sysctl", "-n", "hw.memsize").Output()
	if err != nil {
		fmt.Println("Error: failed to parse RAM info: ", err)
		return ""
	}

	ramBytes, err := strconv.Atoi(strings.Trim(string(output), "\n"))
	if err != nil {
		fmt.Println("Error: failed to parse RAM info to int: ", err)
		return ""
	}

	RAMGb := ramBytes / (1024 * 1024 * 1024)

	return strconv.Itoa(RAMGb) + " GB"
}

func getPageSize() int64 {
	out, err := exec.Command("sysctl", "-n", "hw.pagesize").Output()
	if err != nil {
		return 16384
	}
	val, _ := strconv.ParseInt(strings.TrimSpace(string(out)), 10, 64)
	return val
}

func GetUsedRAM() string {
	out, err := exec.Command("vm_stat").Output()
	if err != nil {
		fmt.Println("Error: failed to execute vm_stat: ", err)
		return "unknown"
	}

	pageSize := getPageSize()
	var activePages, wiredPages, compressedPages int64

	lines := strings.Split(string(out), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Pages active:") {
			activePages = parseVmStatValue(line)
		}

		if strings.HasPrefix(line, "Pages wired:") {
			wiredPages = parseVmStatValue(line)
		}

		if strings.HasPrefix(line, "Pages occupied by compressor:") {
			compressedPages = parseVmStatValue(line)
		}
	}

	usedBytes := (activePages + wiredPages + compressedPages) * pageSize

	usedGb := float64(usedBytes) / (1024 * 1024 * 1024)

	return fmt.Sprintf("%.2f GB", usedGb)
}

func parseVmStatValue(line string) int64 {
	parts := strings.Split(line, ":")
	if len(parts) < 2 {
		return 0
	}

	valueStr := strings.TrimSpace(parts[1])
	valueStr = strings.TrimSuffix(valueStr, ".")

	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		return 0
	}
	return value
}

func GetGPU() string {
	out, err := exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output()
	if err != nil {
		return "Unknown GPU"
	}

	gpuName := strings.TrimSpace(string(out))

	if gpuName != "" {
		return gpuName
	}

	return "Unknown GPU"
}
