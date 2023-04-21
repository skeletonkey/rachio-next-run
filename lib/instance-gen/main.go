// Package instance_gen is a library for creating a bare-bones application with boilerplate files taken
// care of. The long term goal is to be the basis of all Go applications allowing for quick propagation of updates,
// bug fixes, and new features.
package instance_gen

import (
	"os"
	"path"
	"text/template"

	"github.com/rs/zerolog/log"
)

const templateBaseDir = "lib/instance-gen/templates"
const warning = "instance-gen: File auto generated -- DO NOT EDIT!!!\n"

var templateExts = map[string]string{
	"go":  ".go.tpl",
	"yml": ".yml.tpl",
}
var warnings = map[string]string{
	"go":  "// " + warning,
	"yml": "# " + warning,
}

// App struct containing necessary information for a new application
type App struct {
	dir string
}

// NewApp returns the struct for a new applications which allows for generating boiler plate files.
func NewApp(dir string) App {
	return App{dir: dir}
}

// WithPackages takes a list of strings which results in creating a skeleton subdirectory for each.
// This sets up a config.go file based off of config.go.tpl.
// Each generated file will be prepended with a 'warning' comment to not edit the file.
func (a App) WithPackages(packageNames ...string) App {
	for _, name := range packageNames {
		templateArgs := templateArgs{
			PackageName: name,
		}
		generateTemplate(generateTemplateArgs{
			fileType:       "go",
			outputName:     "config.go",
			outputSubDir:   path.Join(a.dir, name),
			templateName:   "config" + templateExts["go"],
			templateSubDir: "package",
			templateArgs:   templateArgs,
		})
	}
	return a
}

// WithGithubWorkflows sets up the specified workflows
func (a App) WithGithubWorkflows(flows ...string) App {
	for _, name := range flows {
		generateTemplate(generateTemplateArgs{
			fileType:       "yml",
			outputName:     name + ".yml",
			outputSubDir:   path.Join(".github", "workflows"),
			templateName:   name + templateExts["yml"],
			templateSubDir: "github_workflows",
		})
	}
	return a
}

type generateTemplateArgs struct {
	fileType       string       // type of file that the template is for the correct warning message
	outputName     string       // name of the template in its final form
	outputSubDir   string       // sub dir added to root dir for the final file
	templateName   string       // name of the template file
	templateSubDir string       // sub dir added to the template base dir to find template
	templateArgs   templateArgs // args that are fed to text/template
}
type templateArgs struct {
	PackageName string // name of the package
}

func generateTemplate(args generateTemplateArgs) {
	inputFileName := path.Join(templateBaseDir, args.templateSubDir, args.templateName)
	outputFileName := path.Join(args.outputSubDir, args.outputName)

	log.Debug().Str("outputDir", args.outputSubDir)
	err := os.MkdirAll(args.outputSubDir, 0755)
	if err != nil {
		log.Panic().Err(err).
			Str("outputDir", args.outputSubDir).
			Msg("unable to create directory structure")
	}
	log.Debug().Str("creating file", outputFileName)
	f, err := os.Create(outputFileName)
	if err != nil {
		log.Panic().Err(err).Str("file name", outputFileName).Msg("unable to create file")
	}
	_, err = f.WriteString(warnings[args.fileType])
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
	err = temp.Execute(f, args.templateArgs)
	if err != nil {
		log.Panic().Err(err).Str("clientName", args.outputName).Msg("unable to execute template")
	}
}
