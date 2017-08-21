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
var regFunc = regexp.MustCompile(`\s*(.[^\(]*)\s*\(\s*((\S*)\s*(\S*))\s*,\s*((\S*)\s*(\S*))\s*,\s*((\S*)\s*(\S*))\s*\)\s*\(((\S*)\s*,\s*(\S*))\s*\)\s*`)

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
	for _, match := range regInterfaces.FindAllString(string(protoContent), -1) {

		// parse line by line
		for _, lines := range strings.Split(match, "\n") {
			result := regFunc.FindStringSubmatch(lines)

			if len(result) <= 0 {
				continue
			}

			// check if is interface already at service file
			if !regexp.MustCompile(`\s*` + regexp.QuoteMeta(result[1]) +
				`\s*\(\s*` +
				regexp.QuoteMeta(result[3]) +
				`\s*` +
				regexp.QuoteMeta(result[4]) +
				`\s*\,\s*` +
				regexp.QuoteMeta(result[6]) +
				`\s*` +
				regexp.QuoteMeta(result[7]) +
				`\s*\)\s*\(\s*` +
				regexp.QuoteMeta(result[12]) +
				`\s*\,\s*` +
				regexp.QuoteMeta(result[13])).
				MatchString(string(serviceContent)) {

				log.Println("Adding " + result[1] + "interface to service")

				// open file to append interfaces
				var file, err = os.OpenFile(outputPath, os.O_APPEND|os.O_WRONLY, 0600)
				if err != nil {
					return err
				}
				defer file.Close()

				// added to file new interfaces
				if _, err := file.WriteString("\nfunc (s *Service) " +
					result[1] +
					"(" +
					result[2] +
					", " +
					result[5] +
					") (" +
					result[11] +
					") {\n" +
					"\n" +
					"\treturn &" + result[12][1:] + "{}, nil\n" +
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
