// Package instance_gen instance_gen is a library for creating a barebones appliaction with boiler plate files taken
// care of. The long term goal is to be the basis of all Go applications allowing for quick propagation of updates,
// bug fixes, and new features.
package instance_gen

import (
	"os"
	"path"
	"text/template"

	"github.com/rs/zerolog/log"
)

const templateDir = "lib/instance-gen/templates"
const warning = "// instance-gen: File auto generated -- DO NOT EDIT!!!\n"

// App struct containing necessary information for a new application
type App struct {
	dir string
}

// NewApp returns the struct for a new applications which allows for generating boiler plate files.
func NewApp(dir string) App {
	return App{dir: dir}
}
func (a App) WithClients(clientNames ...string) App {
	for _, name := range clientNames {
		a.generateTemplate("client", name)
	}
	return a
}

func (a App) generateTemplate(templateName string, clientName string) {
	inputFileName := path.Join(templateDir, templateName+".go.tpl")
	outputDir := path.Join(a.dir, clientName)
	outputFileName := path.Join(outputDir, templateName+".go")

	log.Debug().Str("outputDir", outputDir)
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Panic().Err(err).
			Str("outputDir", outputDir).
			Msg("unable to create directory structure")
	}
	log.Debug().Str("creating file", outputFileName)
	f, err := os.Create(outputFileName)
	if err != nil {
		log.Panic().Err(err).Str("file name", outputFileName).Msg("unable to create file")
	}
	_, err = f.WriteString(warning)
	if err != nil {
		log.Panic().Err(err).Str("file name", outputFileName).Msg("unable to write warning to file")
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Error().Err(err).Str("file name", outputFileName).Msg("error closing file")
		}
	}()
	if err != nil {
		log.Panic().Err(err).
			Str("outputFilename", outputFileName).
			Msg("unable to create new file")
	}
	temp := template.Must(template.ParseFiles(inputFileName))
	err = temp.Execute(f, clientName)
	if err != nil {
		log.Panic().Err(err).Str("clientName", clientName).Msg("unable to execute template")
	}
}
