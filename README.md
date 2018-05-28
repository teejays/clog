# Clog

Clog package provides a small, simple, colored logging library for your Go projects. It can log to the syslog and your terminal's std output. 

There are two main _type_ in this package: a Decoration and a Clogger. 

__Decoration__ is an ANSI escape sequence, hence a string, that can be used to format a message that is logged to the standard out (terminal). 

__Clogger__ is the primary logger object, a logger profile in other words. It holds information neccesary to log with a certain style to both Syslog and Std. Out. Therefore, messages logged with the same Clogger show same styles and use the same decorations. 

The package comes with some default Cloggers, namely Debug, Info, Warning, Notice, Error, Critical, Fatal. These cloggers have preset configuration making it very easy to use it out of the box

Click [here](https://godoc.org/github.com/teejays/clogger) for code documentation.
 
## Getting Started

Install Clog package in your system  by running the following go get command in your terminal:

`go get -u github.com/teejays/clog`

You can then include clog package in your Go project by adding the following import statement in your .go file:
```go
import github.com/teejays/clog
```

You can log using one of the default _Clogger_ or create your own. There are six default Cloggers available: Debug, Info, Notice, Warning, Error, Crit. For these default Cloggers, which is what you will mostly need in your project, you can use the built in functions for quick logging. Here is a sample logging code for the Info Clogger.

```go
clog.Info("This is a simple logging message using the info Clogger")
clog.Infof("This is a formatted logging message using the %s Clogger", "info")
```

## Decorations
By default, decorated logging i.e. logging with colors etc. is turned on. You can turn it off by setting the _UseDecoration_ flag to false.
```go
clog.UseDecoration = false
```
All the default Cloggers have pre-defined decorations associated with them. You can change the them by using AddDecoration() and RemoveDecoration() methods on the Clogger. You can either use one of the Decorations provided as constants, or create and use your own if you have the ANSI code. For example, the Error Clogger is by default set to log using a red color, which you can change if you want. 
```go
// change the color of Error Clogger to one of the provided color contsants
cl := clog.GetCloggerByName("Error")
cl.RemoveDecoration(clogger.FG_RED) // remove the red color decoration
cl.AddDecoration(clogger.FG_YELLOW) // add the yellow color decoration
cl.AddDecoration(clogger.BRIGHT) 	// let's make it bright as well, because why not
```
```go
// change the color of Error Clogger to your own color
fgYellow := clog.NewDecoration("\x1b[33;1m") / create your own Decoration
cl := GetCloggerByName("Error")
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
cl := clog.NewClogger("myClogger", syslog.LOG_WARNING|syslog.LOG_LOCAL1, FG_RED, BG_BLUE, BRIGHT)
```
You can then use your myClogger from anywhere in your project/executable by calling your saved clogger and using print functions.
```go
myClogger := clog.GetCloggerByName("myClogger")
myClogger.Print("This is a simple logging message using myClogger")
myClogger.Printf("This is a simple logging message using %s", "myClogger")
```

 ### Contact
For any issues, please open a new issue. If you want to contribute, please feel free to submit a merge request or reach out to me at clog@teejay.me.
