package log

import (
	lg "log"

	aw "github.com/deanishe/awgo"
)

// isDebug is pseudo-private and indicates whether the workflow is in debug mode.
// It is set to true if the workflow is running in debug mode.
var isDebug bool

// Init initializes the logger with the given version and workflow.
// It sets the log flags and prefix for the logger.
// If the workflow is in debug mode, it sets the isDebug variable to true.
// This function should be called at the beginning of the program to set up logging.
func Init(version string, wf *aw.Workflow) {
	lg.SetFlags(lg.LstdFlags | lg.Lshortfile)
	lg.SetPrefix("[version: " + version + "] ")

	if wf.Debug() {
		isDebug = true
	}
}

func Debug(msg string, args ...interface{}) {
	if isDebug {
		lg.Println("[DEBUG]", msg, args)
	}
}

func Debugf(msg string, args ...interface{}) {
	if isDebug {
		lg.Printf("[DEBUG] "+msg, args...)
	}
}

func Info(msg string, args ...interface{}) {
	lg.Println("[INFO]", msg, args)
}

func Infof(msg string, args ...interface{}) {
	lg.Printf("[INFO] "+msg, args...)
}

func Error(msg string, args ...interface{}) {
	lg.Println("[ERROR] ", msg, args)
}

func Errorf(msg string, args ...interface{}) {
	lg.Printf("[ERROR] "+msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	lg.Fatalln("[FATAL]", msg, args)
}

func Fatalf(msg string, args ...interface{}) {
	lg.Fatalf("[FATAL] "+msg, args...)
}

func Panic(msg string, args ...interface{}) {
	lg.Panicln("[PANIC]", msg, args)
}

func Panicf(msg string, args ...interface{}) {
	lg.Panicf("[PANIC] "+msg, args...)
}
