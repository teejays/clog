# Clogger

Clogger package provides a small, simple clored logging library for your Go projects. It logs to the syslog, but allows logging to terminal's std output as well. It comes with several default logging profiles (called Cloggers) such as Debug, Info, Warning, Error, Critical, Fatal. You can also create your own Clogger using the exported methods, choosing a color of your choice.

Click [here](https://godoc.org/github.com/teejays/clogger) for code documentation.
 
## Getting Started

Install Clogger package in your system  by running the following go get command in your terminal:

`go get -u github.com/teejays/clogger`

You can then include Clogger in your personal Go project by adding the following import statement in your Go file:
```go
import github.com/teejays/copperchain
```

You can log using one of the default profiles (also called default _Cloggers_) or create your own one. There are six default Cloggers available: Debug, Info, Notice, Warning, Error, Crit. For these default Cloggers, which is what you will mostly need in your project, you can use the built in functions for quick logging. Here is a sample logging code for the Info Clogger.

```go
clogger.Info("This is a simple logging message using the info Clogger")
clogger.Infof("This is a formatted logging message using the %s Clogger", "info")
```

## Color
By default, colored logging is turned off. You can turn it on by setting the _UseColor_ flag to true.
```go
clogger.UseColor = true
```
All the default Cloggers come have colors associated with them by default. You can change the colors by using SetColor() method on the Clogger. You can either use a provided color, or create your own if you have the ANSI code. 
```go
// change the color of Error Clogger to one of the provided color contsants
cl := GetCloggerByName("Error")
cl.SetColor(clogger.COLOR_CYAN)
```
```go
// change the color of Error Clogger to your own color
yellow := clogger.NewColor("\x1b[33;1m")
cl := GetCloggerByName("Error")
cl.SetColor(yellow)
```

## Terminal (Standard Out)
By default, clogger package logs the messages to [Syslog](https://en.wikipedia.org/wiki/Syslog) + the standard out (i.e. your terminal). If you want to stop logging to your terminal, you can set the _LogToStdOut_ flag to false.
```go
clogger.LogToStdOut = false
```
While logging to the terminal, clogger package would prepend all the messages with a timestamp. You can remove the timestamps by _UseTimestamp_ flag to false.
```go
clogger.UseTimestamp = false
```

## Create your own Clogger
Although you will rarely have to, you can create, save, and use a custom Clogger if you want. This allows you specify the logging priority, color of the Clogger yourself. 
```go
cl := clogger.NewClogger(syslog.LOG_WARNING|syslog.LOG_LOCAL1, COLOR_RED)
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
