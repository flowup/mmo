package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("INTERFACES")

}

func Parse(path string) (error) {
	interfaceResult := ""

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	for i, match := range regexp.MustCompile(`type .[^\s]*Client interface {(\s*.[^\n]*\s*[^\}]*)}`).FindStringSubmatch(string(content)) {
		if i%2 == 0 {
			continue
		}

		for _, lines := range strings.Split(match, "\n") {
			result := regexp.MustCompile(`\s*(.[^\(]*)\s*\(\s*((\S*)\s*(\S*))\s*,\s*((\S*)\s*(\S*))\s*,\s*((\S*)\s*(\S*))\s*\)\s*\(((\S*)\s*,\s*(\S*))\s*\)\s*`).FindStringSubmatch(lines)

			if len(result) <= 0 {
				continue
			}

			interfaceResult += "func " + result[0][1:] + " {\n" +
				"\n" +
				"\treturn &"+result[7][1:]+"{}, nil\n" +
				"}\n\n"

		}
	}

	fmt.Println(interfaceResult)

	return nil
}
