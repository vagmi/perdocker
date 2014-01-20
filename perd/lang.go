package perd

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

func (l *Lang) RunCommand(filePath string) string {
	return l.Command + " " + filePath
}

func (l *Lang) ExecutableFile() string {
	return l.FileName
}

var Ruby *Lang = &Lang{"ruby", "run.rb", ".rb", "perdocker/ruby:user", "ruby"}
var Nodejs *Lang = &Lang{"nodejs", "index.js", ".js", "perdocker/nodejs", "node"}
