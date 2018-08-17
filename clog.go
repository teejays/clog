// clog package provides a small, simple, fast logging library for your Go projects.
// It can log to the syslog and/or your terminal's std output.
//
// There are two main types in this package that you need to understand: Decoration, and Clogger.
// Decoration is an ANSI escape sequence, hence a string, that can be used to format a message
// that is logged to the standard out (terminal).
//
// Clog is the primary logger object, a profile. It holds information
// neccesary for both Syslog and Std. Out logging for that particular profile. Therefore, messages
// logged with the same Clogger will show same behavior and use the same decorations. This package
// comes with some default Cloggers, namely Debug, Info, Warning, Error, Critical, Fatal. These cloggers have
// preset configuration making it very easy to use it out of the box.
package clog

import (
	"fmt"
	"log"
	"time"
)

const PACKAGE_NAME string = `Clog`

// LogToStdOut flag determines if messages should be logged to the standard terminal output
var LogToStdOut bool = true

// LogToSyslog flag determines if messages should be logged to the syslog
var LogToSyslog bool = false

// UseDecoration flag determines whether standard output logs should use any of the decorations associated with the logger
var UseDecoration bool = true

// PrependTimestamp flag determines whether standard output logs should prepend timestamp
var PrependTimestamp bool = true

// PrependLoggerName determines whether standard output logs with the name of the logger profile prepended
var PrependLoggerName bool = true

// TimestampFormat is the format of the timestamp that is prepernded to std out logs. The default value
// is 2006/01/02 15:04:05
var TimestampFormat string = "2006/01/02 15:04:05"

// Info logs the msg using the "Info" default clogger.
func Info(msg string) {
	clogger := GetCloggerByName("Info")
	clogger.Print(msg)
}

// Infof formats the message using the provided args, and logs the message using the 'Info' default clogger.
func Infof(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Info")
	clogger.Printf(formatString, args...)
}

// Debug logs the msg using the "Debug" default clogger.
func Debug(msg string) {
	clogger := GetCloggerByName("Debug")
	clogger.Print(msg)
}

// Debugf formats the message using the provided args, and logs the message using the 'Debug' default clogger.
func Debugf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Debug")
	clogger.Printf(formatString, args...)
}

// Notice logs the msg using the "Notice" default clogger.
func Notice(msg string) {
	clogger := GetCloggerByName("Notice")
	clogger.Print(msg)
}

// Noticef formats the message using the provided args, and logs the message using the 'Notice' default clogger.
func Noticef(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Notice")
	clogger.Printf(formatString, args...)
}

// Warning logs the msg using the "Warning" default clogger.
func Warning(msg string) {
	clogger := GetCloggerByName("Warning")
	clogger.Print(msg)
}

// Warningf formats the message using the provided args, and logs the message using the 'Warning' default clogger.
func Warningf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Warning")
	clogger.Printf(formatString, args...)
}

// Warn logs the msg using the "Warning" default clogger.
func Warn(msg string) {
	Warning(msg)
}

// Warningf formats the message using the provided args, and logs the message using the 'Warning' default clogger.
func Warnf(formatString string, args ...interface{}) {
	Warningf(formatString, args...)
}

// Error logs the msg using the "Error" default clogger.
func Error(msg string) {
	clogger := GetCloggerByName("Error")
	clogger.Print(msg)
}

// Errorf formats the message using the provided args, and logs the message using the 'Error' default clogger.
func Errorf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Error")
	clogger.Printf(formatString, args...)
}

// Crit logs the msg using the "Crit" default clogger.
func Crit(msg string) {
	clogger := GetCloggerByName("Crit")
	clogger.Print(msg)
}

// Critf formats the message using the provided args, and logs the message using the 'Crit' default clogger.
func Critf(formatString string, args ...interface{}) {
	clogger := GetCloggerByName("Crit")
	clogger.Printf(formatString, args...)
}

// Fatal logs the msg using the "Fatal" default clogger. It also terminates the process by calling log.Fatal.
func Fatal(msg string) {
	Crit(msg)
	log.Fatal(msg)
}

// Fatalf formats the message using the provided args, and logs the message using the 'Fatal' default clogger.
// It also terminates the process by calling log.Fatalf.
func Fatalf(formatString string, args ...interface{}) {
	Critf(formatString, args...)
	log.Fatalf(formatString, args...)
}

func Redf(msg string, args ...interface{}) {
	Red(fmt.Sprintf(msg, args...))
}

func Red(msg string) {
	PrintWithDecorations(msg, FG_RED)
}

func Greenf(msg string, args ...interface{}) {
	Green(fmt.Sprintf(msg, args...))
}

func Green(msg string) {
	PrintWithDecorations(msg, FG_GREEN)
}

func Yellowf(msg string, args ...interface{}) {
	Yellow(fmt.Sprintf(msg, args...))
}

func Yellow(msg string) {
	PrintWithDecorations(msg, FG_YELLOW)
}

func Bluef(msg string, args ...interface{}) {
	Blue(fmt.Sprintf(msg, args...))
}
func Blue(msg string) {
	PrintWithDecorations(msg, FG_BLUE)
}

func Println(msg string) {
	fmt.Println(msg)
}

func Printf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func PrintWithDecorations(msg string, decorations ...Decoration) {
	msg = decorate(msg, decorations...)
	fmt.Println(msg)
}

// Panic takes an error as an argument and calls logs.Panic
func Panic(err error) {
	log.Panic(err)
}

func prependTimestamp(msg string) string {
	return fmt.Sprintf("%s %s", timestamp(), msg)
}

func decorate(msg string, Decorations ...Decoration) string {
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
