package asciiart

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"strings"
)

var HeightChar = 9
var tabChar = map[int][]string{}
var errNum int = 200

func AsciiArt(textInput, bannerName string) (string, int) {
	var tabChars []string
	scanner, err := ioutil.ReadFile("./asciiart/banner/" + bannerName + ".txt")
	if err != nil {
		errNum = http.StatusInternalServerError
		return "Invalid Banner !!!", errNum
	}
	data := bufio.NewScanner(strings.NewReader(string(scanner)))
	for data.Scan() {
		lines := data.Text()
		tabChars = append(tabChars, lines)
	}
	characters := len(tabChars) / HeightChar
	for i := 0; i < characters; i++ {
		charLines := tabChars[i*HeightChar : (i+1)*HeightChar]
		tabChar[int(i)] = charLines

	}
	return generateAsciiArt(textInput)

}

func generateAsciiArt(input string) (string, int) {
	var result string
	inputLines := strings.Split(input, "\r\n")
	for _, inputLine := range inputLines {
		for i := 1; i < HeightChar; i++ {
			for _, char := range inputLine {
				if int(char) >= 32 && int(char) <= 126 {
					chars := int(char) - 32
					line := tabChar[chars][i]
					result += string(line)
				} else {
					errNum = 400
					return ("\"" + string(char) + "\"" + " is not printable in Ascii-Art"), errNum
				}
			}
			result += "\n"
		}
	}
	errNum = 200
	return result, errNum
}
