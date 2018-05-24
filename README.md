# Clogger

Clogger package provides a small, simple clored logging library for your Go projects. It logs to the syslog and/or your terminal's std output. 

There are two main _type_ in this package: a Decoration and a Clogger. 

__Decoration__ is an ANSI escape sequence, hence a string, that can be used to format a message that is logged to the standard out (terminal). 

__Clogger__ is the primary logger object, a logger profile in other words. It holds information neccesary for both Syslog and Std. Out logging for that particular profile. Therefore, messages logged with the same Clogger will show same behavior and use the same decorations. 

The package comes with some default Cloggers, namely Debug, Info, Warning, Error, Critical, Fatal. These cloggers have preset configuration making it very easy to use it out of the box

Click [here](https://godoc.org/github.com/teejays/clogger) for code documentation.
 
## Getting Started

Install Clogger package in your system  by running the following go get command in your terminal:

`go get -u github.com/teejays/clogger`

You can then include Clogger in your personal Go project by adding the following import statement in your Go file:
```go
import github.com/teejays/clogger
```

You can log using one of the default profiles (also called default _Cloggers_) or create your own one. There are six default Cloggers available: Debug, Info, Notice, Warning, Error, Crit. For these default Cloggers, which is what you will mostly need in your project, you can use the built in functions for quick logging. Here is a sample logging code for the Info Clogger.

```go
clogger.Info("This is a simple logging message using the info Clogger")
clogger.Infof("This is a formatted logging message using the %s Clogger", "info")
```

## Decorations
By default, decorated logging is turned off. You can turn it on by setting the _UseDecoration_ flag to true.
```go
clogger.UseDecoration = true
```
All the default Cloggers have decorations associated with them. You can change the them by using AddDecoration() and RemoveDecoration() methods on the Clogger. You can either use one of the Decorations provided as constants, or create and use your own if you have the ANSI code. 
```go
// change the color of Error Clogger to one of the provided color contsants
cl := GetCloggerByName("Error")
cl.RemoveDecoration(clogger.FG_CYAN)
cl.AddDecoration(clogger.FG_YELLOW)
cl.AddDecoration(clogger.BRIGHT)
```
```go
// change the color of Error Clogger to your own color
fgYellow := clogger.NewDecoration("\x1b[33;1m")
cl := GetCloggerByName("Error")
cl.AddDecoration(yellow)
```

## Logging Outputs (Syslog vs. Std. Out)
By default, clogger package logs messages to both the [Syslog](https://en.wikipedia.org/wiki/Syslog) and the standard out (i.e. your terminal). If you want to stop logging to either channel your terminal, you can use the below flags.
```go
clogger.LogToStdOut = false // stop logging to std. out
clogger.LogToSyslog = false // stop logging to syslog
```
While logging to the std. out (terminal), clogger would prepend all the messages with a timestamp. You can stop this behavior by setting the _UseTimestamp_ flag to false.
```go
clogger.UseTimestamp = false
```

## Create your own Clogger
Although you will rarely have to, you can create, save, and use a custom Clogger if you want. This allows you specify the logging priority and decorations of your Clogger. The following code demonstrates how this can be done.
```go
cl := clogger.NewClogger(syslog.LOG_WARNING|syslog.LOG_LOCAL1, FG_RED, BG_BLUE, BRIGHT)
clogger.RegisterClogger("myClogger", cl)
```
You can then use your saved clogger from anywhere in your project/executable by calling your saved clogger and using print functions.
```go
myClogger := clogger.GetCloggerByName("myClogger")
myClogger.Print("This is a simple logging message using myClogger")
myClogger.Printf("This is a simple logging message using %s", "myClogger")
```

 ### Contact
For any issues, please open a new issue. If you want to contribute, please feel free to submit a merge request or reach out to me at clogger@teejay.me.
