// clogger package provides a small, simple, fast logging library for your Go projects.
// It logs to the syslog, but allows logging to terminal's std output as well. It comes
// with several default logging profiles (called Cloggers) such as Debug, Info, Warning,
// Error, Critical, Fatal.
// You can also create your own Clogger using the exported methods, choosing a color of
// your choice.
package clogger

import (
	"fmt"
	"log"
	"log/syslog"
	"regexp"
	"time"
)

const PACKAGE_NAME string = `clogger`

// LogToStdOut flag determines if messages should be to the standard terminal output
var LogToStdOut bool = true

// UseColor flag determines whether standard output logs should have color
var UseColor bool

// UseTimestamp flag determines whether standard output logs should prepend timestamp
var UseTimestamp bool = true

// TimestampFormat is the format of the timestamp that is prepernded to std out logs. The default value
// is 2006/01/02 15:04:05
var TimestampFormat string = "2006/01/02 15:04:05"

/********************************************************************************
* C O L O R  																	*
*********************************************************************************/

// Color defines a SGR code
type Color string

const (
	COLOR_RESET  Color = "\033[0m"
	COLOR_RED    Color = "\x1b[31;1m"
	COLOR_GREEN  Color = "\x1b[32;1m"
	COLOR_YELLOW Color = "\x1b[33;1m"
	COLOR_BLUE   Color = "\x1b[34;1m"
	COLOR_MAROON Color = "\x1b[35;1m"
	COLOR_CYAN   Color = "\x1b[36;1m"
	COLOR_WHITE  Color = "\x1b[37;1m"
)

// NewColor takes a string representation of sgr code, casts it as Color, and returns it. It panics if sgrCode is not
// a valid ansi color code
func NewColor(sgrCode string) Color {
	// verify that it's an ansi code
	// regex from: https://superuser.com/questions/380772/removing-ansi-color-codes-from-text-stream
	reg := regexp.MustCompile(`^\x1b\[[0-9;]*[mG]$`)
	if !reg.MatchString(sgrCode) {
		panic(fmt.Sprintf("%s: invalid sgr code '%s' provided", PACKAGE_NAME, sgrCode))
	}
	return Color(sgrCode)
}

/********************************************************************************
* S Y S L O G G E R   															*
*********************************************************************************/

// Clogger is the primary logger of this package. It is basically a syslog.Logger but with extra attributes, such as color.
type Clogger struct {
	syslog.Priority
	Color
	*log.Logger
}

// Print logs the message in the Syslog and, if LogToStdOut flag is set to true, it logs
// the message to the standard out.
func (s *Clogger) Print(msg string) {
	s.Logger.Print(msg)
	s.StdPrint(msg)
}

// Printf formats the msg with the provided args and logs to Syslog. If LogToStdOut flag
//  is set to true, it also logs the message to the standard out.
func (s *Clogger) Printf(formatString string, args ...interface{}) {
	s.Logger.Printf(formatString, args...)
	s.StdPrintf(formatString, args...)
}

// StdPrint prints msg as a line in the std out. If configured, it also appends the timestamp and uses color.
func (s *Clogger) StdPrint(msg string) {
	if UseTimestamp {
		msg = appendTimestamp(msg)
	}
	if UseColor {
		msg = addColor(msg, s.Color)
	}
	fmt.Println(msg)
}

// StdPrintf formats msg with the provided args and prints it as a line in the standard output. If configured,
// it appends the timestamp and uses color.
func (s *Clogger) StdPrintf(formatString string, args ...interface{}) {
	msg := fmt.Sprintf(formatString, args...)
	s.StdPrint(msg)
}

// SetColor sets the color of the Clogger.
func (s *Clogger) SetColor(color Color) {
	s.Color = color
}

func appendTimestamp(msg string) string {
	return fmt.Sprintf("%s %s", timestamp(), msg)
}
func addColor(msg string, Color Color) string {
	return fmt.Sprintf("%s%s%s", Color, msg, COLOR_RESET)
}
func addBreak(msg string) string {
	return fmt.Sprintf("%s%s", msg, "\n")
}

func timestamp() string {
	return time.Now().Format(TimestampFormat)
}

/********************************************************************************
* S Y S L O G G E R   															*
*********************************************************************************/

// default cloggers
var cloggers map[string]*Clogger = map[string]*Clogger{
	"Debug":   NewClogger(syslog.LOG_DEBUG|syslog.LOG_LOCAL1, COLOR_WHITE),
	"Info":    NewClogger(syslog.LOG_INFO|syslog.LOG_LOCAL1, COLOR_GREEN),
	"Notice":  NewClogger(syslog.LOG_NOTICE|syslog.LOG_LOCAL1, COLOR_WHITE),
	"Warning": NewClogger(syslog.LOG_WARNING|syslog.LOG_LOCAL1, COLOR_YELLOW),
	"Error":   NewClogger(syslog.LOG_ERR|syslog.LOG_LOCAL1, COLOR_RED),
	"Crit":    NewClogger(syslog.LOG_CRIT|syslog.LOG_LOCAL1, COLOR_MAROON),
}

func init() {
	createDefaultCloggers()
}

func createDefaultCloggers() {
	var err error
	// https://en.wikipedia.org/wiki/Syslog
	for _, cl := range cloggers {
		cl.Logger, err = syslog.NewLogger(cl.Priority, 0)
		if err != nil {
			log.Fatalf("%s: there has been an error starting the logger: %q", PACKAGE_NAME, err)
		}
	}
}

// NewClogger accepts priority in the form of syslog.Priority, and a Color, and returns
// a pointer to a new Clogger object with those properties. It panics if it encounters an error.
func NewClogger(priority syslog.Priority, color Color) *Clogger {
	clogger := new(Clogger)
	clogger.Priority = priority
	clogger.Color = color
	logger, err := syslog.NewLogger(clogger.Priority, 0)
	if err != nil {
		log.Panic(err)
	}
	clogger.Logger = logger

	return clogger
}

// RegisterLogger adds a new custom Clogger to the system, which can then be fetched by calling
// the GetCloggerByName method.
func RegisterClogger(name string, cl *Clogger) error {
	if _, exists := cloggers[name]; exists {
		return fmt.Errorf("%s: a logger with the name %s already exists", PACKAGE_NAME, name)
	}
	cloggers[name] = cl
	return nil
}

// GetCloggerByName provides the pointer to the Clogger that is stored by the given name.
// It panics if a clogger by that name doesn't exist.
func GetCloggerByName(name string) *Clogger {
	cl, exist := cloggers[name]
	// panics if loggers[name] doesn't exist
	if !exist {
		log.Panicf("%s: no logger with name %s", PACKAGE_NAME, name)
	}
	return cl
}

// Info logs the msg using the "Info" default clogger.
func Info(msg string) {
	clogger := GetCloggerByName("Info")
	clogger.Print(msg)
	if LogToStdOut {
		clogger.StdPrint(msg)
	}
}

// Debug logs the msg using the "Debug" default clogger.
func Debug(msg string) {
	clogger := GetCloggerByName("Debug")
	clogger.Print(msg)
	if LogToStdOut {
		clogger.StdPrint(msg)
	}
}

// Notice logs the msg using the "Notice" default clogger.
func Notice(msg string) {
	clogger := GetCloggerByName("Notice")
	clogger.Print(msg)
	if LogToStdOut {
		clogger.StdPrint(msg)
	}
}

// Warning logs the msg using the "Warning" default clogger.
func Warning(msg string) {
	clogger := GetCloggerByName("Warning")
	clogger.Print(msg)
	if LogToStdOut {
		clogger.StdPrint(msg)
	}
}

// Error logs the msg using the "Error" default clogger.
func Error(msg string) {
	clogger := GetCloggerByName("Error")
	clogger.Print(msg)
	if LogToStdOut {
		clogger.StdPrint(msg)
	}
}

// Crit logs the msg using the "Crit" default clogger.
func Crit(msg string) {
	clogger := GetCloggerByName("Crit")
	clogger.Print(msg)
	if LogToStdOut {
		clogger.StdPrint(msg)
	}
}

// Fatal logs the msg using the "Fatal" default clogger. It also terminates the process by calling log.Fatal.
func Fatal(msg string) {
	Crit(msg)
	log.Fatal(msg)
}

// Infof formats the message using the provided args, and logs the message using the 'Info' default clogger.
func Infof(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Info")
	clogger.Printf(formatString, args...)
	if LogToStdOut {
		clogger.StdPrintf(formatString, args...)
	}
}

// Debugf formats the message using the provided args, and logs the message using the 'Debug' default clogger.
func Debugf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Debug")
	clogger.Printf(formatString, args...)
	if LogToStdOut {
		clogger.StdPrintf(formatString, args...)
	}
}

// Noticef formats the message using the provided args, and logs the message using the 'Notice' default clogger.
func Noticef(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Notice")
	clogger.Printf(formatString, args...)
	if LogToStdOut {
		clogger.StdPrintf(formatString, args...)
	}
}

// Warningf formats the message using the provided args, and logs the message using the 'Warning' default clogger.
func Warningf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Warning")
	clogger.Printf(formatString, args...)
	if LogToStdOut {
		clogger.StdPrintf(formatString, args...)
	}
}

// Errorf formats the message using the provided args, and logs the message using the 'Error' default clogger.
func Errorf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Error")
	clogger.Printf(formatString, args...)
	if LogToStdOut {
		clogger.StdPrintf(formatString, args...)
	}
}

// Critf formats the message using the provided args, and logs the message using the 'Crit' default clogger.
func Critf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Crit")
	clogger.Printf(formatString, args...)
	if LogToStdOut {
		clogger.StdPrintf(formatString, args...)
	}
}

// Fatalf formats the message using the provided args, and logs the message using the 'Fatal' default clogger.
// It also terminates the process by calling log.Fatalf.
func Fatalf(formatString string, args ...interface{}) {
	Critf(formatString, args...)
	log.Fatalf(formatString, args...)
}
