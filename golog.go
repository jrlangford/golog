package golog

import (
	"io"
	"log"
)

type Loggers struct {
	info    *log.Logger
	debug   *log.Logger
	warning *log.Logger
	error   *log.Logger
	fatal   *log.Logger
}

func (l *Loggers) Info(v ...interface{}) {
	l.info.Println(v)
}

func (l *Loggers) Debug(v ...interface{}) {
	l.debug.Println(v)
}

func (l *Loggers) Warning(v ...interface{}) {
	l.warning.Println(v)
}

func (l *Loggers) Error(v ...interface{}) {
	l.error.Println(v)
}

func (l *Loggers) Fatal(v ...interface{}) {
	// TODO: send signal to sigchan in main instead of exiting fatally to
	// shutdown server as cleanly as possible
	l.fatal.Fatalln(v)
}

func (l *Loggers) Init(w io.Writer) {

	logFlags := log.Ldate | log.Ltime

	l.info = log.New(w, "INFO: ", logFlags)
	l.debug = log.New(w, "DEBUG: ", logFlags)
	l.warning = log.New(w, "WARNING: ", logFlags)
	l.error = log.New(w, "ERROR: ", logFlags)
	l.fatal = log.New(w, "FATAL: ", logFlags)
}
