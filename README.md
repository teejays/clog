# Clog

Clog package provides a small, simple, colored logging library for your Go projects. It can log to the Standard Output (read: your terminal) and Unix Syslog. 

See this in your terminal instead of the boring standard logs!

![Clog Terminal Screenshot](https://i.gyazo.com/f7853cfa693f2f11f94e401dbb75c514.gif)

The aim of this package is to make it brain dead easy for any Go developer to see color-rich output logs. Hence, Clog comes with ready-use functions that can simply replace the Go's standard __fmt.Print...__ and __log.Print...__ type functions.

There are a few different ways to use this package, depending on how advanced you want to go. The simplest way is to use the color functions, that automatically print in color. For example:

```go
clog.Red("This is a red statement")
clog.Redf("This is a %s statement", "red")

clog.Yellow("This is a yellow statement")
clog.Yellowf("This is a %s statement", "yellow")

clog.Green("This is a green statement")
clog.Greenf("This is a %s statement", "green")

```

The entire list of color printing functions is available in the [docs](https://godoc.org/github.com/teejays/clog).

For more sophisticated projects where you want a closer control over the type of logging statements and the logging level, you could use the log-level functions. There are six default loggers available: Debug, Info, Notice, Warning/Warn, Error, Crit. For these default Cloggers, which is what you will mostly need in your project, you can use the built in functions for quick logging. For example:

```go
clog.Debug("This is a simple debug logging message")
clog.Debugf("This is a formatted %s logging message", "debug")

clog.Info("This is a simple info logging message")
clog.Infof("This is a formatted %s logging message", "info")

clog.Notice("This is a simple notice logging message")
clog.Noticef("This is a formatted %s logging message", "notice")

clog.Warn("This is a simple warn logging message")
clog.Warnf("This is a formatted %s logging message", "warn")

```
Depending on the logging level set, you can stop some of these statements from being printed. Log level can be set by simply setting the LogLevel variable. There are six log levels (from 0-5). Higher the log level, less the logging. For example, for a log level of 4, only logs of log level Err or Crit will be printed. Different log levels can be found in the [docs](https://godoc.org/github.com/teejays/clog).

```go
clog.LogLevel = 5
```

Click [here](https://godoc.org/github.com/teejays/clog) for code documentation.
 
_**Windows Users**_: The package has not been tested for Windows command prompt. 

## Getting Started

Install Clog package in your system  by running the following go get command in your terminal:

`go get -u github.com/teejays/clog`

You can then include clog package in your Go project by adding the following import statement in your .go file:
```go
import github.com/teejays/clog
```

Then you can start calling the Clog functions in your project, just like in the examples provided above. 


## Structure

This section is primarily for advanced users and describes how the projec is set up.

### Types

There are two main _type_ in this package: a Decoration and a Clogger. 

__Decoration__ is an ANSI escape sequence, hence a string, that can be used to format a message that is logged to the standard out (terminal). 

__Clogger__ is the primary logger object, a logger profile in other words. It holds information neccesary to log with a certain style to both Syslog and Std. Out. Therefore, messages logged with the same Clogger show same styles and use the same decorations. 

The package comes with some default Cloggers, namely Debug, Info, Warning, Notice, Error, Critical, Fatal. These cloggers have preset configuration making it very easy to use it out of the box

### Decorations
By default, decorated logging i.e. logging with colors etc. is turned on. You can turn it off by setting the _UseDecoration_ flag to false.
```go
clog.UseDecoration = false
```
All the default Cloggers have pre-defined decorations associated with them. You can change the them by using AddDecoration() and RemoveDecoration() methods on the Clogger. You can either use one of the Decorations provided as constants, or create and use your own if you have the ANSI code. For example, the Error Clogger is by default set to log using a red color, which you can change if you want. 
```go
// change the color of Error Clogger to one of the provided color contsants
cl := clog.GetCloggerByName("Error")
cl.RemoveDecoration(clog.FG_RED) // remove the red color decoration
cl.AddDecoration(clog.FG_YELLOW) // add the yellow color decoration
cl.AddDecoration(clog.BRIGHT) 	// let's make it bright as well, because why not
```
```go
// change the color of Error Clogger to your own color
fgYellow := clog.NewDecoration("\x1b[33;1m") / create your own Decoration
cl := clog.GetCloggerByName("Error")
cl.AddDecoration(yellow)
```

## Logging Outputs (Syslog vs. Std. Out)
By default, clogger package logs messages only to the standard output (i.e. the terminal). It does not log to [Syslog](https://en.wikipedia.org/wiki/Syslog). If you want to enable or disable logging to either, you can change the below flags.
```go
clog.LogToStdOut = false // stop logging to standard output
clog.LogToSyslog = true // start logging to syslog
```
While logging to the standard output (terminal), clog package would prepend all the messages with a timestamp. You can stop this behavior by setting the _UseTimestamp_ flag to false.
```go
clog.UseTimestamp = false
```

## Create your own Clogger
Although you will rarely have to, you can create, save, and use a custom Clogger if you want. This allows you to specify the logging priority and your own decorations for your Clogger. The following code demonstrates how this can be done.
```go
cl := clog.NewClogger("myClogger", syslog.LOG_WARNING|syslog.LOG_LOCAL1, clog.FG_RED, clog.BG_BLUE, clog.BRIGHT)
```
You can then use your myClogger from anywhere in your project/executable by calling your saved clogger and using print functions.
```go
myClogger := clog.GetCloggerByName("myClogger")
myClogger.Print("This is a simple logging message using myClogger")
myClogger.Printf("This is a simple logging message using %s", "myClogger")
```

 ### Contact
For any issues, please open a new issue. If you want to contribute, please feel free to submit a merge request or reach out to me at clog@teejay.me.
