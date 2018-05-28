package main

import (
	"github.com/teejays/clogger"
	//"log/syslog"
)

func main() {

	clogger.LogToSyslog = false

	// print a bunch of statements
	clogger.Debug("Debug: This is a debug statement.")
	clogger.Debugf("Debug: This is a %s statement.", "Debug")

	clogger.Info("Info: This is a Info statement.")
	clogger.Infof("Info: This is a %s statement.", "Info")

	clogger.Warning("Warning: This is a Warning statement.")
	clogger.Warningf("Warning: This is a %s statement.", "Warning")

	clogger.Error("Error: This is a Error statement.")
	clogger.Errorf("Error: This is a %s statement.", "Error")

	clogger.Crit("Crit: This is a Crit statement.")
	clogger.Critf("Crit: This is a %s statement.", "Crit")

	// self logger
	//myLogger := clogger.NewClogger(syslog.LOG_INFO|syslog.LOG_LOCAL1, clogger.BLINK, clogger.BG_WHITE, clogger.FG_GREEN)
	//myLogger.Print("myLogger: This is a myLogger statement")
	//myLogger.Printf("myLogger: This is a %s statement.", "myLogger")
}
