package function

import (
	"bufio"

	"os"
	"strings"
)


func TraitmentData(bnr string,  arg string) string {
	banner := bnr

	
	fileName := "./banners/"+banner+".txt"
	// Open the ASCII art file
	file, err := os.Open(fileName)
	if err != nil {
		return "Error opening the file"
	}
	defer file.Close()
	var asciiArt []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "Error reading the file"
	}
	
	var result string
	lines := strings.Split(arg, "\\n")
	for _, line := range lines {
		if line == "" {
			result += "\n"
			continue
		}
		// Iterate over each row of the ASCII art (0 to 7, for the 8 rows)
		for i := 1; i <= 8; i++ {
			for _, r := range line {
				// Ensure the character is within the valid ASCII range
				if r < 32 || r > 126 {
					return "Please enter a valid character between ascii code 32 and 126"
				}
				index := 9*(int(r)-32) + i
				result += asciiArt[index]
			}
			result += "\n" // Add newline after finishing the current row of the line
		}
	}
	return result
}


func BannerExists(banner string) bool {
	// Check if the provided banner exists (this is a placeholder function)
	// For example, compare the banner string to a list of supported banners
	supportedBanners := []string{"standard", "shadow", "thinkertoy"}
	for _, b := range supportedBanners {
		if banner == b {
			return true
		}
	}
	return false
}
