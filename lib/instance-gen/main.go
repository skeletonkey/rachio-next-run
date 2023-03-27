package instance_gen

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

const templateDir = "lib/instance-gen/templates"
const warning = "// instance-gen: File auto generated -- DO NOT EDIT!!!\n"

type app struct {
	dir string
}

func NewApp(dir string) app {
	return app{dir: dir}
}
func (a app) WithClients(clientNames ...string) app {
	for _, name := range clientNames {
		a.generateTemplate("client", name)
	}
	return a
}

func (a app) generateTemplate(templateName string, clientName string) {
	inputFileName := path.Join(templateDir, templateName + ".go.tpl")
	outputDir := path.Join(a.dir, clientName)
	outputFileName := path.Join(outputDir, templateName + ".go")

	//fmt.Printf("Creating Dir: %s\n", outputDir)
	os.MkdirAll(outputDir, 0755)
	//fmt.Printf("Creating File: %s\n", outputFileName)
	f, err := os.Create(outputFileName)
	f.WriteString(warning)
	defer f.Close()
	if err != nil {
		panic(fmt.Errorf("unable to create file (%s): %s", outputFileName, err))
	}
	temp := template.Must(template.ParseFiles(inputFileName))
	temp.Execute(f, clientName)
}
