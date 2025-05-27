package fabric

// This file is the heart of the package. Has everything needed to generate code based on the provided parameters.
// Author: Joel Dos Santos. Date: 26/05/2025
// Feel free to contribute to this package by adding more templates or improving the existing ones!
// Also, feel free to create issues or pull requests if you find any bugs or have suggestions for improvements!

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

type (
	Param struct {
		Name      string
		ParamType string
		Optional  bool
	}
	Fabricable struct {
		Name     string
		Params   []Param
		Language string

		self any
	}
	FabricableErr struct {
		Message string
		Stdout  string
	}
)

func (f *Fabricable) Initialize(parent any) {
	f.self = parent
}

func (e FabricableErr) Error() string {
	return fmt.Sprintf(e.Message, e.Stdout)
}

const (
	InvalidLanguageErr         = "invalid language provided: %s. \n(Stdout: %s)"
	InvalidParamFormatErr      = "invalid parameter format, expected 'name:type:optional'. \n(Stdout: %s)"
	InvalidOptionalValueErr    = "invalid optional value, expected 'true' or 'false': \n(Stdout: %s)"
	UnsupportedTemplateTypeErr = "unsupported template type, supported types are: ts, dart. \n(Stdout: %s)"
	ParsingErr                 = "error parsing template. \n(Stdout: %s)"
	CreatingErr                = "error creating file. \n(Stdout: %s)"
	ExecuteTemplateErr         = "error executing template. \n(Stdout: %s)"
	NotEnoughArgsErr           = "not enough arguments provided. \nExpected: [lang] [name] [param1:type1:true] [param2:type2:false]... [paramN:typeN:optional]. \n(Stdout: %s)"
	WorkingDirectoryErr        = "error getting current working directory. \n(Stdout: %s)"
)

func (f *Fabricable) Generate(rawTemplate string, path string, fileprefix string) error {
	funcMap := template.FuncMap{
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		// Don't care + didn't ask (Not important if it is deprecated or not for the usecase of this)
		"title":     strings.Title,
		"camelCase": toCamelCase,
		"snakeCase": toSnakeCase,
	}

	filetype, err := GetFileType(f.Language)
	if err != nil {
		return FabricableErr{
			Message: UnsupportedTemplateTypeErr,
			Stdout:  err.Error(),
		}
	}

	tmpl := template.New("fabricable").Funcs(funcMap)
	prsdTmpl, err := tmpl.Parse(rawTemplate)
	if err != nil {
		return FabricableErr{
			Message: ParsingErr,
			Stdout:  err.Error(),
		}
	}

	dir, err := os.Getwd()
	if err != nil {
		return FabricableErr{
			Message: WorkingDirectoryErr,
			Stdout:  err.Error(),
		}
	}

	if _, err := os.Stat(dir + "/" + path); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Join(dir, path), os.ModePerm)
		if err != nil {
			return FabricableErr{
				Message: CreatingErr,
				Stdout:  err.Error(),
			}
		}
	}

	rawFilename := f.Name + fileprefix
	outputDir := filepath.Join(dir+"/"+path, strings.ToLower(rawFilename)+"."+filetype)

	file, err := os.Create(outputDir)
	if err != nil {
		file.Close()
		return FabricableErr{
			Message: CreatingErr,
			Stdout:  err.Error(),
		}
	}
	defer file.Close()

	err = prsdTmpl.Execute(file, f.self)
	if err != nil {
		return FabricableErr{
			Message: ExecuteTemplateErr,
			Stdout:  err.Error(),
		}
	}

	fmt.Printf("Successfully created %s\n", outputDir)
	return nil
}

func (f *Fabricable) CreateFromArgs(cmd *cobra.Command, args []string) error {
	// TODO: Make this better :)
	if len(args) < 3 {
		return FabricableErr{
			Message: NotEnoughArgsErr,
			Stdout:  "",
		}
	}

	f.Language = args[1]
	f.Name = args[2]
	f.Params = []Param{}

	for _, arg := range args[3:] {
		param := strings.Split(arg, ":")
		if len(param) != 3 {
			fmt.Printf("Invalid parameter format: %s\n", arg)
			continue
		}

		name := strings.TrimSpace(param[0])
		paramType := strings.TrimSpace(param[1])
		optional, err := strconv.ParseBool(strings.TrimSpace(param[2]))
		if err != nil {
			fmt.Printf("Invalid optional value for parameter %s: %v\n", name, err)
			continue
		}

		f.Params = append(f.Params, Param{name, paramType, optional})
	}

	return nil
}

func toCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' '
	})

	if len(words) == 0 {
		return s
	}

	result := strings.ToLower(words[0])
	for _, word := range words[1:] {
		result += strings.Title(strings.ToLower(word))
	}
	return result
}

func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && (r >= 'A' && r <= 'Z') {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
