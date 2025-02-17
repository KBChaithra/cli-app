package cmd

import (
	"cli-app/tasks"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	configFile string
	processDir string
)

type Process struct {
	Name   string `yaml:"name"`
	Params Params `yaml:"params"`
	Tasks  []Task `yaml:"tasks"`
}

type Params struct {
	User string `yaml:"user"`
	Host string `yaml:"host"`
}

type Task struct {
	Name   string            `yaml:"id"`
	Class  string            `yaml:"class"`
	Params map[string]string `yaml:"params"`
}

var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the process defined in YAML file",
	Run: func(cmd *cobra.Command, args []string) {
		if configFile == "" || processDir == "" {
			log.Println("config file and process directory must be provided")
			os.Exit(1)
		}

		err := runProcess()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&processDir, "taskdir", "t", "", "directory containing process YAML files")
}

func runProcess() error {
	log.Printf("reading config file: %s", configFile)
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	var config map[string]interface{}
	if err := json.Unmarshal(configData, &config); err != nil {
		return fmt.Errorf("error unmarshaling config JSON: %v", err)
	}

	processName, ok := config["process_name"].(string)
	if !ok {
		return fmt.Errorf("invalid process name format")
	}
	processParams, ok := config["process_params"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid process params format")
	}

	log.Printf("reading process file for process name: %s", processName)
	processYAML, err := readProcessFromDir(processDir, processName)
	if err != nil {
		return fmt.Errorf("error reading process file: %v", err)
	}

	var process Process
	if err := yaml.Unmarshal(processYAML, &process); err != nil {
		return fmt.Errorf("error unmarshaling process YAML: %v", err)
	}

	log.Printf("executing process: %s", process.Name)
	return executeProcess(process, processParams)
}

func readProcessFromDir(dir, name string) ([]byte, error) {
	log.Printf("reading directory: %s", dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %v", dir, err)
	}

	for _, file := range files {
		log.Printf("checking file: %s", file.Name())
		if file.Name() == name+".yaml" {
			log.Printf("found process file: %s", file.Name())
			return ioutil.ReadFile(filepath.Join(dir, file.Name()))
		}
	}
	return nil, fmt.Errorf("process file %s.yaml not found in directory %s", name, dir)
}

func executeProcess(process Process, params map[string]interface{}) error {
	for _, task := range process.Tasks {
		for key, value := range task.Params {
			tmplStr := (value)
			tmpl, err := template.New(key).Parse(tmplStr)
			if err != nil {
				log.Printf("error parsing template: %v", err)
				continue
			}

			var executedStr strings.Builder
			if err := tmpl.Execute(&executedStr, params); err != nil {
				log.Printf("error executing template: %v", err)
				continue
			}

			task.Params[key] = executedStr.String()
		}

		log.Printf("executing task: %s", task.Name)
		switch task.Class {
		case "localCmd":
			tasks.RunLocalCmd(task.Params["command"])
		case "sshCmd":
			tasks.RunSSHCmd(task.Params)
		case "writefile":
			tasks.WriteFile(task.Params)
		default:
			log.Printf("unknown task class: %v", task.Class)
		}
	}

	return nil
}

func Execute() error {
	return rootCmd.Execute()
}
