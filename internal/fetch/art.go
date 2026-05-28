package fetch

import (
	"fmt"
	"os"
	"path"
	"ponyfetch"
	"ponyfetch/pkg/colors"
	"regexp"
	"strings"
)

func GetArt(ponyName string, size string) string {
	filename := ponyName + ".txt"

	if size == "little" {
		base := strings.TrimSuffix(filename, path.Ext(filename))
		filename = base + "_little.txt"
	}

	art, err := ponyfetch.AssetsFS.ReadFile("assets/" + filename)
	if err != nil {
		fmt.Printf("Warning: Pony '%s' not found, falling back to twilight\n", ponyName)

		defaultFilename := "twilight.txt"
		if size == "little" {
			defaultFilename = "twilight_little.txt"
		}

		art, _ = ponyfetch.AssetsFS.ReadFile("assets/" + defaultFilename)
	}

	return string(art)
}

func GetArtFromFile(filename string, size string) string {
	if !strings.HasSuffix(filename, ".txt") {
		defaultFilename := "twilight.txt"
		if size == "little" {
			defaultFilename = "twilight_little.txt"
		}

		art, _ := ponyfetch.AssetsFS.ReadFile("assets/" + defaultFilename)

		return string(art)
	}

	art, _ := os.ReadFile(filename)

	return string(art)
}

func PrintFetch(ponyName string, size string, color string, filename string) {
	if size != "little" && size != "normal" {
		size = "normal"
	}

	var artRaw string

	if filename != "" {
		artRaw = GetArtFromFile(filename, size)
	} else {
		artRaw = GetArt(ponyName, size)
	}

	artLines := strings.Split(artRaw, "\n")
	colorCode := colors.GetColorCode(color)
	resetCode := colors.GetColorCode("reset")
	re := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

	var infoLines []string

	hostInfo := fmt.Sprintf("%s @ %s", GetUser(), GetHostname())
	infoLines = append(infoLines, hostInfo)
	infoLines = append(infoLines, strings.Repeat("-", len(hostInfo)))

	infoLines = append(infoLines, "")
	infoLines = append(infoLines, "OS: "+GetOSInfo())
	infoLines = append(infoLines, "Shell: "+GetShell())
	infoLines = append(infoLines, "DE: "+GetDesktopEnvironment())
	infoLines = append(infoLines, "")
	infoLines = append(infoLines, "Hardware")
	infoLines = append(infoLines, "CPU: "+GetCPU())
	infoLines = append(infoLines, "GPU: "+GetGPU())
	infoLines = append(infoLines, fmt.Sprintf("Memory: %s / %s", GetUsedRAM(), GetTotalRAM()))

	maxLines := len(artLines)
	if len(infoLines) > maxLines {
		maxLines = len(infoLines)
	}

	artWidth := 0
	for _, line := range artLines {
		cleanLine := re.ReplaceAllString(line, "")
		cleanLine = strings.ReplaceAll(cleanLine, "\r", "")
		lineLen := len([]rune(cleanLine))
		if lineLen > artWidth {
			artWidth = lineLen
		}
	}

	for i := 0; i < maxLines; i++ {
		artPart := ""
		infoPart := ""

		if i < len(artLines) {
			artPart = strings.ReplaceAll(artLines[i], "\r", "")
		}

		if i < len(infoLines) {
			infoPart = infoLines[i]
		}

		if artWidth == 0 {
			fmt.Println(infoPart)
			continue
		}

		cleanArtLine := re.ReplaceAllString(artPart, "")
		currentArtLen := len([]rune(cleanArtLine))

		var padding string
		if currentArtLen < artWidth {
			padding = strings.Repeat(" ", artWidth-currentArtLen)
		} else {
			padding = " "
		}

		var finalArtOutput string

		if color == "reset" {
			finalArtOutput = artPart
		} else {
			finalArtOutput = colorCode + cleanArtLine + resetCode
		}

		fmt.Printf("%s%s   %s\n", finalArtOutput, padding, infoPart)
	}
}
