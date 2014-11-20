package perd

import (
	"fmt"
)

// Lang is a struct to store language settings (name, Docker image, command to exec code, etc.).
type Lang struct {
	Name     string
	FileName string
	Ext      string
	Image    string
	Command  string
}

func (l *Lang) uniqFileName() string {
	return uniqFileName() + l.Ext
}

// RunCommand forms command using Lang.Command and a given string.
// Example: `ruby /tmp/perdocker/run.rb`
func (l *Lang) RunCommand(filePath string) string {
	return fmt.Sprintf(l.Command, filePath)
}

// ExecutableFile returns filename, which will be used to store user's code
func (l *Lang) ExecutableFile() string {
	return l.FileName
}

// Ruby language settings
var Ruby = &Lang{"ruby", "run.rb", ".rb", "vagmi/perdruby:attach", "ruby %s"}

// Nodejs settings
var Nodejs = &Lang{"nodejs", "index.js", ".js", "vagmi/perdnodejs:attach", "node %s"}

// Golang settings
var Golang = &Lang{"golang", "main.go", ".go", "vagmi/perdgo:attach", "go run %s"}

// Python settings
var Python = &Lang{"python", "run.py", ".py", "vagmi/perdpython:attach", "python %s"}

// C settings
var C = &Lang{"c", "a.c", ".c", "vagmi/perdc:attach", "gcc -o /tmp/a %s && /tmp/a"}

// JAVA settings
var JAVA = &Lang{"java", "program.java", ".java", "vagmi/perdjava:attach", "javac -d /tmp %s && java -cp /tmp Program"}

// CPP settings
var CPP = &Lang{"cpp", "a.cpp", ".cpp", "vagmi/perdc:attach", "g++ -o /tmp/a %s && /tmp/a"}

// PHP settings
var PHP = &Lang{"php", "index.php", ".php", "vagmi/perdphp:attach", "php %s"}

// Universal languages container
var Universal = &Lang{"universal", "file", "", "vagmi/perduniversal:latest", "cat %s"}

var Languages = map[string]*Lang{
	"ruby":       Ruby,
	"nodejs":     Nodejs,
	"javascript": Nodejs,
	"golang":     Golang,
	"python":     Python,
	"c":          C,
	"cpp":        CPP,
	"java":       JAVA,
	"php":        PHP,
}
