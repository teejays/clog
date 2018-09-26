package main

import (
	"fmt" // Required for Example 3
	"github.com/teejays/clog"
)

func main() {

	clog.LogToSyslog = true // Windows does not have Syslog, so syslog wouldn't work on Windows

	clog.LogLevel = clog.LogLevelInfo

	// Example 1: Default Cloggers
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

	// Example 2: Creating your own Clogger
	myLogger := clog.NewClogger("myClogger", clog.LogLevelInfo, clog.BLINK, clog.FG_GREEN)
	myLogger.Print("myLogger: This is a myLogger statement")
	myLogger.Printf("myLogger: This is a %s statement.", "myLogger")

	// Example 3: clog.Print functions
	clog.PrintWithDecorations("This is my own message!", clog.BLINK, clog.BG_YELLOW, clog.BRIGHT, clog.FG_BLUE)
	clog.PrintWithDecorations(fmt.Sprintf("This is %s message!", "my own"), clog.BLINK, clog.BG_YELLOW, clog.BRIGHT, clog.FG_BLUE)

	// Example 4: Color functions
	clog.Red("This is a red statement")
	clog.Yellow("This is a yellow statement")
	clog.Green("This is a green statement")
	clog.Blue("This is a blue statement")
}
