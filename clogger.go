// clogger package provides a small, simple, fast logging library for your Go projects.
// It can log to the syslog and/or your terminal's std output.
//
// There are two main types in this package that you need to understand: Decoration, and Clogger.
// Decoration is an ANSI escape sequence, hence a string, that can be used to format a message
// that is logged to the standard out (terminal).
//
// Clogger is the primary logger object, a profile. It holds information
// neccesary for both Syslog and Std. Out logging for that particular profile. Therefore, messages
// logged with the same Clogger will show same behavior and use the same decorations. This package
// comes with some default Cloggers, namely Debug, Info, Warning, Error, Critical, Fatal. These cloggers have
// preset configuration making it very easy to use it out of the box.
package clogger

import (
	"fmt"
	"log"
	"log/syslog"
	"regexp"
	"time"
)

const PACKAGE_NAME string = `clogger`

// LogToStdOut flag determines if messages should be logged to the standard terminal output
var LogToStdOut bool = true

// LogToSyslog flag determines if messages should be logged to the syslog
var LogToSyslog bool = false

// UseDecoration flag determines whether standard output logs should use any of the decorations associated with the logger
var UseDecoration bool

// UseTimestamp flag determines whether standard output logs should prepend timestamp
var UseTimestamp bool = true

// TimestampFormat is the format of the timestamp that is prepernded to std out logs. The default value
// is 2006/01/02 15:04:05
var TimestampFormat string = "2006/01/02 15:04:05"

/********************************************************************************
* D E C O R A T I O N 															*
*********************************************************************************/

// Decoration represents an ANSI escape sequence, that can be used to format a message
// logged to the standard out (terminal).
type Decoration string

const (
	RESET Decoration = "\x1b[0m"

	// decorations
	BRIGHT     Decoration = "\x1b[1m"
	DIM        Decoration = "\x1b[2m"
	UNDERSCORE Decoration = "\x1b[4m"
	BLINK      Decoration = "\x1b[5m"
	REVERSE    Decoration = "\x1b[7m"
	HIDDEN     Decoration = "\x1b[8m"

	// foreground colors represent the color of the logged text
	FG_BLACK   Decoration = "\x1b[30m"
	FG_RED     Decoration = "\x1b[31m"
	FG_GREEN   Decoration = "\x1b[32m"
	FG_YELLOW  Decoration = "\x1b[33m"
	FG_BLUE    Decoration = "\x1b[34m"
	FG_MAGENTA Decoration = "\x1b[35m"
	FG_CYAN    Decoration = "\x1b[36m"
	FG_WHITE   Decoration = "\x1b[37m"

	// background colors represent the background color of the logged text
	BG_BLACK   Decoration = "\x1b[40m"
	BG_RED     Decoration = "\x1b[41m"
	BG_GREEN   Decoration = "\x1b[42m"
	BG_YELLOW  Decoration = "\x1b[43m"
	BG_BLUE    Decoration = "\x1b[44m"
	BG_MAGENTA Decoration = "\x1b[45m"
	BG_CYAN    Decoration = "\x1b[46m"
	BG_WHITE   Decoration = "\x1b[47m"
)

// NewDecoration takes a string representation of sgr code (ANSI), casts it as a Decoration, and returns it. It panics if the sgrCode is not
// a valid ansi escape sequence code.
func NewDecoration(sgrCode string) Decoration {
	// verify that it's an ansi code
	// regex from: https://superuser.com/questions/380772/removing-ansi-color-codes-from-text-stream
	reg := regexp.MustCompile(`^\x1b\[[0-9;]*[mG]$`)
	if !reg.MatchString(sgrCode) {
		panic(fmt.Sprintf("%s: invalid sgr code '%s' provided", PACKAGE_NAME, sgrCode))
	}
	return Decoration(sgrCode)
}

/********************************************************************************
* S Y S L O G G E R   															*
*********************************************************************************/

// Clogger is the primary logger of this package. It represents a logger profile that has
// associated decorations Decorations, syslog priority level. This package come with some
// default Cloggers, but Clogger can also be created using the NewClogger() method.
type Clogger struct {
	syslog.Priority
	Decorations []Decoration
	*log.Logger
}

// NewClogger accepts priority in the form of syslog.Priority, and a Color, and returns
// a pointer to a new Clogger object with those properties. It panics if it encounters an error.
func NewClogger(priority syslog.Priority, decorations ...Decoration) *Clogger {
	clogger := new(Clogger)
	clogger.Priority = priority
	clogger.Decorations = decorations
	if LogToSyslog {
		logger, err := syslog.NewLogger(clogger.Priority, 0)
		if err != nil {
			log.Panic(err)
		}
		clogger.Logger = logger
	}

	return clogger
}

// AddDecoration adds the decoration to the Clogger.
func (l *Clogger) AddDecoration(d Decoration) {
	l.Decorations = append(l.Decorations, d)
}

// RemoveDecoration removes the decoration from the Clogger.
func (l *Clogger) RemoveDecoration(d Decoration) {
	for i, _d := range l.Decorations {
		if d == _d {
			// delete the decoration from the list
			l.Decorations = append(l.Decorations[:i], l.Decorations[i+1:]...)
		}
	}
}

// Print logs the message in the Syslog if LogToSyslog is set to true. It logs to the standard out
// (terminal) if LogToStdOut flag is set to true.
func (l *Clogger) Print(msg string) {
	if LogToSyslog {
		l.Logger.Print(msg)
	}
	if LogToStdOut {
		l.StdPrint(msg)
	}
}

// Printf formats the msg with the provided args and logs to Syslog. If LogToStdOut flag
// is set to true, it also logs the message to the standard out. Printf formats the message
// with the provided args. It logs the message in the Syslog if LogToSyslog is
// set to true. It logs to the standard out (terminal) if LogToStdOut flag is set to true.
func (l *Clogger) Printf(formatString string, args ...interface{}) {
	if LogToSyslog {
		l.Logger.Printf(formatString, args...)
	}
	if LogToStdOut {
		l.StdPrintf(formatString, args...)
	}
}

// StdPrint prints msg as a line in the standard output (terminal). If UseTimestamp is set to true,
// it prepends timestamp to the log messages. If UseDecoration is set to true, it adds all the decorations
// associated with the l Clogger.
func (l *Clogger) StdPrint(msg string) {
	if UseTimestamp {
		msg = appendTimestamp(msg)
	}
	if UseDecoration {
		msg = decorate(msg, l.Decorations)
	}
	fmt.Println(msg)
}

// StdPrintf formats msg with the provided args and prints it as a line in the standard output. If UseTimestamp is
// set to true, it prepends timestamp to the log messages. If UseDecoration is set to true, it adds all the decorations
// associated with the l Clogger.
func (l *Clogger) StdPrintf(formatString string, args ...interface{}) {
	msg := fmt.Sprintf(formatString, args...)
	l.StdPrint(msg)
}

func appendTimestamp(msg string) string {
	return fmt.Sprintf("%s %s", timestamp(), msg)
}

func decorate(msg string, Decorations []Decoration) string {
	var decorationsCode string
	for _, d := range Decorations {
		decorationsCode += string(d)
	}
	return fmt.Sprintf("%s%s%s", decorationsCode, msg, RESET)
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
	"Debug":   NewClogger(syslog.LOG_DEBUG|syslog.LOG_LOCAL1, FG_WHITE),
	"Info":    NewClogger(syslog.LOG_INFO|syslog.LOG_LOCAL1, FG_GREEN),
	"Notice":  NewClogger(syslog.LOG_NOTICE|syslog.LOG_LOCAL1, FG_WHITE),
	"Warning": NewClogger(syslog.LOG_WARNING|syslog.LOG_LOCAL1, FG_YELLOW),
	"Error":   NewClogger(syslog.LOG_ERR|syslog.LOG_LOCAL1, FG_RED),
	"Crit":    NewClogger(syslog.LOG_CRIT|syslog.LOG_LOCAL1, FG_MAGENTA),
}

func init() {
	createDefaultCloggers()
}

func createDefaultCloggers() {
	var err error
	// https://en.wikipedia.org/wiki/Syslog
	for _, cl := range cloggers {
		if LogToSyslog {
			cl.Logger, err = syslog.NewLogger(cl.Priority, 0)
			if err != nil {
				log.Fatalf("%s: there has been an error starting the logger: %q", PACKAGE_NAME, err)
			}
		}

	}
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
}

// Debug logs the msg using the "Debug" default clogger.
func Debug(msg string) {
	clogger := GetCloggerByName("Debug")
	clogger.Print(msg)
}

// Notice logs the msg using the "Notice" default clogger.
func Notice(msg string) {
	clogger := GetCloggerByName("Notice")
	clogger.Print(msg)
}

// Warning logs the msg using the "Warning" default clogger.
func Warning(msg string) {
	clogger := GetCloggerByName("Warning")
	clogger.Print(msg)
}

// Error logs the msg using the "Error" default clogger.
func Error(msg string) {
	clogger := GetCloggerByName("Error")
	clogger.Print(msg)
}

// Crit logs the msg using the "Crit" default clogger.
func Crit(msg string) {
	clogger := GetCloggerByName("Crit")
	clogger.Print(msg)
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
}

// Debugf formats the message using the provided args, and logs the message using the 'Debug' default clogger.
func Debugf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Debug")
	clogger.Printf(formatString, args...)
}

// Noticef formats the message using the provided args, and logs the message using the 'Notice' default clogger.
func Noticef(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Notice")
	clogger.Printf(formatString, args...)
}

// Warningf formats the message using the provided args, and logs the message using the 'Warning' default clogger.
func Warningf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Warning")
	clogger.Printf(formatString, args...)
}

// Errorf formats the message using the provided args, and logs the message using the 'Error' default clogger.
func Errorf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Error")
	clogger.Printf(formatString, args...)
}

// Critf formats the message using the provided args, and logs the message using the 'Crit' default clogger.
func Critf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Crit")
	clogger.Printf(formatString, args...)
}

// Fatalf formats the message using the provided args, and logs the message using the 'Fatal' default clogger.
// It also terminates the process by calling log.Fatalf.
func Fatalf(formatString string, args ...interface{}) {
	Critf(formatString, args...)
	log.Fatalf(formatString, args...)
}
