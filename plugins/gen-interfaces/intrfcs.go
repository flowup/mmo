package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"github.com/flowup/gogen"
	"path/filepath"
	"os/exec"
)

func main() {
	services := os.Args[1:]

	for _, service := range services {
		outputPath := "/source/" + service + "/service.go"

		log.Println("Generating service of " + service + " from interface")
		newInterfaces, err := Parse("/source/"+service+"/proto.pb.go", outputPath, true)
		if err != nil {
			log.Fatal(err)
		}
		err = writeToFile(newInterfaces, outputPath)
		if err != nil {
			log.Fatal(err)
		}

		if err := gofmt(outputPath); err != nil {
			log.Fatal(err)
		}
	}
}

var regInterfaces = regexp.MustCompile(`type .[^\s]*Client interface {(\s*.[^\n]*\s*[^\}]*)}`)
var types = map[string]string{
	"string": "\"mock\"",
	"uint32": "1",
	"uint64": "1",
	"int64":  "1",
	"int32":  "1",
	"bool":   "true",
}

func Parse(inputPath, outputPath string, mockGeneration bool) ([]string, error) {
	newInterfaces := []string{}

	// gogen parse all structs
	build, _ := gogen.ParseFile(inputPath)
	file := build.File(filepath.Base(inputPath))

	// load content of proto.pb.go file
	protoContent, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return newInterfaces, err
	}

	// load content of service.go file
	serviceContent, err := ioutil.ReadFile(outputPath)
	if err != nil {
		return newInterfaces, err
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
				`\s*\,\s*\S*\s*` +
				regexp.QuoteMeta(result[4]) +
				`\s*\)\s*\(\s*` +
				regexp.QuoteMeta(result[7]) +
				`\s*\,\s*` +
				regexp.QuoteMeta(result[8])).
				MatchString(string(serviceContent)) {

				// get struct for creating mock
				mock := ""
				if mockGeneration {
					mockStruct := file.Struct(result[7][1:])
					if mockStruct != nil {
						mock = "\n"
						for _, fids := range mockStruct.Fields() {
							key, _ := fids.Type()
							mock += "\t" + fids.Name() + ":" + types[key] + "\n"
						}
					}
				}

				log.Println("Adding " + result[0] + "interface to service")
				// open file to append interfaces
				newInterfaces = append(newInterfaces, "\nfunc (s *Service) "+
					result[0]+
					"("+
					result[1]+
					" "+
					result[2]+
					", "+
					result[3]+
					" "+
					result[4]+
					") ("+
					result[7]+
					", "+
					result[8]+
					") {\n"+
					"\n"+
					"\treturn &"+ result[7][1:]+ "{"+ mock+ "}, nil\n"+
					"}\n")
			}
		}
	}

	return newInterfaces, nil
}

func writeToFile(newInterfaces []string, outputPath string) error {
	var file, err = os.OpenFile(outputPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	// added to file new interfaces
	for _, newInterface := range newInterfaces {
		if _, err := file.WriteString(newInterface); err != nil {
			return err
		}
	}

	if err := file.Sync(); err != nil {
		return err
	}
	return nil
}

func gofmt(output string) error {
	var file, err = os.OpenFile(output, os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	out, err := exec.Command("gofmt", output).Output()
	if err != nil {
		return err
	}

	if _, err := file.Write(out); err != nil {
		return err
	}

	return nil
}
