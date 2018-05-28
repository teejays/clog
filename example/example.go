package main

import (
	"github.com/teejays/clogger"
	//"log"
	"log/syslog"
)

func main() {

	clog.LogToSyslog = true

	// print a bunch of statements
	clog.Debug("Debug: This is a debug statement.")
	clog.Debugf("Debug: This is a %s statement.", "Debug")

	clog.Info("Info: This is a Info statement.")
	clog.Infof("Info: This is a %s statement.", "Info")

	clog.Warning("Warning: This is a Warning statement.")
	clog.Warningf("Warning: This is a %s statement.", "Warning")

	clog.Error("Error: This is a Error statement.")
	clog.Errorf("Error: This is a %s statement.", "Error")

	clog.Crit("Crit: This is a Crit statement.")
	clog.Critf("Crit: This is a %s statement.", "Crit")

	// self logger
	myLogger := clog.NewClogger("myClogger", syslog.LOG_INFO|syslog.LOG_LOCAL1, clog.BLINK, clog.FG_GREEN)
	myLogger.Print("myLogger: This is a myLogger statement")
	myLogger.Printf("myLogger: This is a %s statement.", "myLogger")

	// Print function

	clog.Print("This is my own message!", clog.BLINK, clog.BG_YELLOW, clog.BRIGHT, clog.FG_BLUE)

}
