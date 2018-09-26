package clog

import (
	"fmt"
	"log"
	"log/syslog"
)

const DEFAULT_LOG_FACILITY = syslog.LOG_LOCAL1

var cloggers map[string]*Clogger = make(map[string]*Clogger)

// default cloggers
var defaultCloggers []*Clogger = []*Clogger{
	NewClogger("Debug", LogLevelDebug, FG_WHITE),
	NewClogger("Info", LogLevelInfo, FG_GREEN),
	NewClogger("Notice", LogLevelNotice, FG_CYAN),
	NewClogger("Warning", LogLevelWarning, FG_YELLOW),
	NewClogger("Error", LogLevelError, FG_RED),
	NewClogger("Crit", LogLevelCrit, FG_MAGENTA),
}

// registerLogger adds a new Clogger to the cloggers map, which can then be fetched
// by calling the GetCloggerByName method.
func registerClogger(cl *Clogger) error {
	if _, exists := cloggers[cl.Name]; exists {
		return fmt.Errorf("%s: a logger with the name %s already exists", PACKAGE_NAME, cl.Name)
	}
	cloggers[cl.Name] = cl
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

var LogLevelSysLogPriorityMap map[int]syslog.Priority = map[int]syslog.Priority{
	LogLevelDebug:   syslog.LOG_DEBUG,
	LogLevelInfo:    syslog.LOG_INFO,
	LogLevelNotice:  syslog.LOG_NOTICE,
	LogLevelWarning: syslog.LOG_WARNING,
	LogLevelError:   syslog.LOG_ERR,
	LogLevelCrit:    syslog.LOG_CRIT,
}

/********************************************************************************
* C L O G G E R
*********************************************************************************/

// Clogger is the primary logger of this package. It represents a logger profile that has
// associated decorations, syslog priority level and the go's builtin log.logger struct that
// helps print to syslog. This package come with some default Cloggers, but Clogger can also
// be created using the NewClogger() method.
type Clogger struct {
	Name string
	syslog.Priority
	Decorations []Decoration
	*log.Logger
	LogLevel int
}

// NewClogger creates a new Clogger object. It accepts the name of the new Clogger, priority level
// in the form of syslog.Priority and one or more Decorations. It returns a pointer to a new Clogger
// object with those properties. It panics if it encounters an error.
func NewClogger(name string, logLevel int, decorations ...Decoration) *Clogger {
	clogger := new(Clogger)
	clogger.Name = name
	clogger.LogLevel = logLevel
	// Get the syslog.Level from the map
	priority, hasKey := LogLevelSysLogPriorityMap[logLevel]
	if !hasKey {
		log.Panicf("Invalid LogLevel parameter provided as no syslog.Priority associated with LogLevel %d", logLevel)
	}
	clogger.Priority = priority | DEFAULT_LOG_FACILITY
	clogger.Decorations = decorations
	// https://en.wikipedia.org/wiki/Syslog
	logger, err := syslog.NewLogger(clogger.Priority, 0)
	if err != nil {
		log.Printf("[%s] Clogger profile '%s' will not log to syslog as it failed to initialize syslog.Logger(): %v", PACKAGE_NAME, clogger.Name, err)
	} else {
		clogger.Logger = logger
	}

	err = registerClogger(clogger)
	if err != nil {
		log.Panic(err)
	}
	return clogger
}

// AddDecoration (deprecated) adds the decoration to the Clogger. It probably should not be used
// hence it is being deprecated.
func (l *Clogger) AddDecoration(d Decoration) {
	l.Decorations = append(l.Decorations, d)
}

// RemoveDecoration (deprecated) removes the decorations from the Clogger. It probably should not be used
// hence it is being deprecated.
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
	msg = fmt.Sprintf("[%s] %s", l.Name, msg)
	if LogToSyslog && l.Logger != nil {
		l.Logger.Print(msg)
	}
	if LogToStdOut && LogLevel <= l.LogLevel {
		l.PrintStdOut(msg)
	}
}

// Printf formats the msg with the provided args and logs to Syslog. If LogToStdOut flag
// is set to true, it also logs the message to the standard out. Printf formats the message
// with the provided args. It logs the message in the Syslog if LogToSyslog is
// set to true. It logs to the standard out (terminal) if LogToStdOut flag is set to true.
func (l *Clogger) Printf(formatString string, args ...interface{}) {
	formatString = fmt.Sprintf("[%s] %s", l.Name, formatString)
	if LogToSyslog && l.Logger != nil {
		l.Logger.Printf(formatString, args...)
	}
	if LogToStdOut && LogLevel <= l.LogLevel {
		l.PrintfStdOut(formatString, args...)
	}
}

// StdPrintf formats msg with the provided args and prints it as a line in the standard output. If PrependTimestamp is
// set to true, it prepends timestamp to the log messages. If UseDecoration is set to true, it adds all the decorations
// associated with the l Clogger.
func (l *Clogger) PrintfStdOut(formatString string, args ...interface{}) {
	msg := fmt.Sprintf(formatString, args...)
	l.PrintStdOut(msg)
}

// StdPrint prints msg as a line in the standard output (terminal). If PrependTimestamp is set to true,
// it prepends timestamp to the log messages. If UseDecoration is set to true, it adds all the decorations
// associated with the l Clogger.
func (l *Clogger) PrintStdOut(msg string) {
	if PrependTimestamp {
		msg = prependTimestamp(msg)
	}
	if UseDecoration {
		msg = decorate(msg, l.Decorations...)
	}
	fmt.Println(msg)
}
