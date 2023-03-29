package instance_gen

import (
	"github.com/rs/zerolog/log"
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
	inputFileName := path.Join(templateDir, templateName+".go.tpl")
	outputDir := path.Join(a.dir, clientName)
	outputFileName := path.Join(outputDir, templateName+".go")

	log.Debug().Str("outputDir", outputDir)
	os.MkdirAll(outputDir, 0755)
	log.Debug().Str("creating file", outputFileName)
	f, err := os.Create(outputFileName)
	f.WriteString(warning)
	defer f.Close()
	if err != nil {
		log.Panic().Err(err).
			Str("outputFilename", outputFileName).
			Msg("unable to create new file")
	}
	temp := template.Must(template.ParseFiles(inputFileName))
	temp.Execute(f, clientName)
}
