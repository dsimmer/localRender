package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type configuration struct {
	// space separated "cmd arg1 arg2". Join together like BeforeCommmand + filepath + Aftercommand
	BeforeCommand string
	AfterCommand  string
	// Desired input format ".mp4"
	InputFormat string
	// Desired output format ".mp4". If this doesnt match the output of the command, we will keep retrying the item
	OutputFormat string
	// Service job frequency
	Freq string
	// Folder path to watch
	Target string
}

// Queue for files to be processed
type Queue struct {
	Queue  map[string]bool
	Config configuration
}

/*
FileParser holds the file to parse and the appropriate methods. Require for init are the public properties:
Config
Path
Info
*/
type FileParser struct {
	Config configuration
	Path   string
	Info   os.FileInfo
	Queue  *Queue
}

// ProcessFile uses the info from the FileParser struct to parse the dir and add unparsed items to the queue
func (fp *FileParser) ProcessFile() error {
	if fp.Info.IsDir() || (!strings.Contains(fp.Path, fp.Config.InputFormat)) {
		return nil
	}

	OutputFile := strings.TrimSuffix(fp.Info.Name(), fp.Config.InputFormat) + fp.Config.OutputFormat
	_, err := os.Stat(OutputFile)
	isOutputExist := err == nil
	fmt.Println(fp.Path + " added to queue")
	if isOutputExist {
		fp.Queue.Queue[fp.Path] = true
		check(err)
	}
	return nil
}

// ProcessQueue iterates over the queue and runs the command
func (q *Queue) ProcessQueue() error {
	for key := range q.Queue {
		beforeCommands := strings.Split(q.Config.BeforeCommand+key, " ")
		afterCommands := strings.Split(q.Config.AfterCommand, " ")
		commands := append(beforeCommands, afterCommands...)
		cmd := exec.Command(commands[0], commands[1:]...)
		err := cmd.Run()
		check(err)
		fmt.Println(key + " done")
		delete(q.Queue, key)
	}
	return nil
}

// main simply walks the directory it is invoked in
func main() {
	configFile, err := ioutil.ReadFile("./config.json")
	check(err)
	config := configuration{}
	err = json.Unmarshal(configFile, &config)
	check(err)

	queue := Queue{Config: config}

	for {
		err = filepath.Walk(config.Target, func(path string, info os.FileInfo, err error) error {
			check(err)
			newFileParser := FileParser{
				Path:   path,
				Info:   info,
				Config: config,
				Queue:  &queue,
			}
			return newFileParser.ProcessFile()
		})
		check(err)

		err = queue.ProcessQueue()
		check(err)

		interval, err := time.ParseDuration(config.Freq)
		check(err)
		time.Sleep(interval)
	}
}
