package fetch

import (
	"fmt"
	"ponyfetch"
	"strings"
)

func GetArt(ponyName string, size string) string {
	filename := ponyName + ".txt"

	if size == "little" {
		filename = filename[:len(filename)-4] + "_little.txt"
	}

	art, err := ponyfetch.AssetsFS.ReadFile("assets/" + filename)
	if err != nil {
		fmt.Println("Error: failed to get art: ", err)
	}

	return string(art)
}

func PrintFetch(ponyName string, size string) {
	artRaw := GetArt(ponyName, size)
	artLines := strings.Split(artRaw, "\n")

	var infoLines []string

	hostInfo := fmt.Sprintf("%s @ %s", GetUser(), GetHostname())
	infoLines = append(infoLines, hostInfo)
	infoLines = append(infoLines, strings.Repeat("-", len(hostInfo)))

	infoLines = append(infoLines, "")
	infoLines = append(infoLines, "OS: "+GetOSInfo())
	infoLines = append(infoLines, "Shell: "+GetShell())
	infoLines = append(infoLines, "")
	infoLines = append(infoLines, "Hardware")
	infoLines = append(infoLines, "CPU: "+GetCPU())
	infoLines = append(infoLines, fmt.Sprintf("Memory: %s / %s", GetUsedRAM(), GetTotalRAM()))

	maxLines := len(artLines)
	if len(infoLines) > maxLines {
		maxLines = len(infoLines)
	}

	artWidth := 61
	if size == "little" {
		artWidth = 36
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

		currentArtLen := len([]rune(artPart))

		var padding string
		if currentArtLen < artWidth {
			padding = strings.Repeat(" ", artWidth-currentArtLen)
		} else {
			padding = " "
		}

		fmt.Printf("%s%s   %s\n", artPart, padding, infoPart)
	}
}
