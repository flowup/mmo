package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	services := os.Args[1:]

	for _, service := range services {
		log.Println("Generating service of " + service + " from interface")
		if err := Parse("/source/"+service+"/proto.pb.go", "/source/"+service+"/service.go"); err != nil {
			log.Fatal(err)
		}
	}
}

var regInterfaces = regexp.MustCompile(`type .[^\s]*Client interface {(\s*.[^\n]*\s*[^\}]*)}`)

func Parse(inputPath, outputPath string) error {
	// load content of proto.pb.go file
	protoContent, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return err
	}

	// load content of service.go file
	serviceContent, err := ioutil.ReadFile(outputPath)
	if err != nil {
		return err
	}

	// this regexp find all interfaces
	for _, match := range regInterfaces.FindAllStringSubmatch(string(protoContent), -1) {

		// parse line by line
		for _, lines := range strings.Split(match[1], "\n") {
			result := strings.FieldsFunc(lines, func(r rune) bool {
				return r == '(' || r == ')' || r == ',' || r == '\b' || r == ' ' || r == '\t'
			})

			if len(result) < 9 {
				continue
			}

			// check if is interface already at service file
			if !regexp.MustCompile(`\s*` + regexp.QuoteMeta(result[0]) +
				`\s*\(\s*` +
				regexp.QuoteMeta(result[1]) +
				`\s*` +
				regexp.QuoteMeta(result[2]) +
				`\s*\,\s*` +
				regexp.QuoteMeta(result[3]) +
				`\s*` +
				regexp.QuoteMeta(result[4]) +
				`\s*\)\s*\(\s*` +
				regexp.QuoteMeta(result[7]) +
				`\s*\,\s*` +
				regexp.QuoteMeta(result[8])).
				MatchString(string(serviceContent)) {

				log.Println("Adding " + result[0] + "interface to service")

				// open file to append interfaces
				var file, err = os.OpenFile(outputPath, os.O_APPEND|os.O_WRONLY, 0600)
				if err != nil {
					return err
				}
				defer file.Close()

				// added to file new interfaces
				if _, err := file.WriteString("\nfunc (s *Service) " +
					result[0] +
					"(" +
					result[1] +
					" " +
					result[2] +
					", " +
					result[3] +
					" " +
					result[4] +
					") (" +
					result[7] +
					", " +
					result[8] +
					") {\n" +
					"\n" +
					"\treturn &" + result[7][1:] + "{}, nil\n" +
					"}\n"); err != nil {
					return err
				}

				if err := file.Sync(); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
