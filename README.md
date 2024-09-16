<div align="center">
    <h1>lumber</h1>
    <a href="https://pkg.go.dev/github.com/gleich/lumber/v3"><img alt="Godoc Reference" src="https://godoc.org/github.com/gleich/lumber?status.svg"></a>
    <img alt="lint workflow result" src="https://github.com/gleich/lumber/workflows/lint/badge.svg">
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/gleich/lumber">
    <img alt="Golang report card" src ="https://goreportcard.com/badge/github.com/gleich/lumber/v3">
    <br/>
    <br/>
    <i>Easy to use & pretty logger for golang</i>
</div>
<hr>

- [Install](#-install)
- [Logging Functions](#-logging-functions)
  - [`lumber.Done()`](#lumberDone)
  - [`lumber.Info()`](#lumberinfo)
  - [`lumber.Debug()`](#lumberdebug)
  - [`lumber.Warning()`](#lumberwarning)
  - [`lumber.Error()`](#lumbererror)
  - [`lumber.ErrorMsg()`](#lumbererrormsg)
  - [`lumber.Fatal()`](#lumberfatal)
  - [`lumber.FatalMsg()`](#lumberfatalmsg)
- [Customization](#️-customization)
- [Examples](#-examples)

## Install

Simply run the following from your project root:

```bash
go get -u github.com/gleich/lumber/v3
```

## Logging Functions

### [`lumber.Done()`](https://pkg.go.dev/github.com/gleich/lumber/v3#Done)

Output a "DONE" log.

Demo:

```go
package main

import (
    "time"

    "github.com/gleich/lumber/v3"
)

func main() {
    lumber.Done("booted up the program!")
    time.Sleep(2 * time.Second)
    lumber.Done("waited 2 seconds!")
}
```

Outputs:

![Done output](images/done.png)

### [`lumber.Info()`](https://pkg.go.dev/github.com/gleich/lumber/v3#Info)

Output an info log.

Demo:

```go
package main

import (
    "time"

    "github.com/gleich/lumber/v3"
)

func main() {
    lumber.Info("Getting the current year")
    now := time.Now()
    lumber.Info("Current year is", now.Year())
}
```

Outputs:

![info output](images/info.png)

### [`lumber.Debug()`](https://pkg.go.dev/github.com/gleich/lumber/v3#Debug)

Output a debug log.

Demo:

```go
package main

import (
    "os"

    "github.com/gleich/lumber/v3"
)

func main() {
    homeDir, _ := os.UserHomeDir()
    lumber.Debug("User's home dir is", homeDir)
}
```

Outputs:

![debug output](images/debug.png)

### [`lumber.Warning()`](https://pkg.go.dev/github.com/gleich/lumber/v3#Warning)

Output a warning log.

Demo:

```go
package main

import (
    "time"

    "github.com/gleich/lumber/v3"
)

func main() {
    now := time.Now()
    if now.Year() != 2004 {
        lumber.Warning("Current year isn't 2004")
    }
}
```

Outputs:

![warning output](images/warning.png)

### [`lumber.Error()`](https://pkg.go.dev/github.com/gleich/lumber/v3#Error)

Output an error log with a stack trace.

Demo:

```go
package main

import (
    "os"

    "github.com/gleich/lumber/v3"
)

func main() {
    fname := "invisible-file.txt"
    _, err := os.ReadFile(fName)
    if err != nil {
        lumber.Error(err, "Failed to read from", fname)
    }
}
```

Outputs:

![error output](images/error.png)

### [`lumber.ErrorMsg()`](https://pkg.go.dev/github.com/gleich/lumber/v3#ErrorMsg)

Output an error message.

Demo:

```go
package main

import "github.com/gleich/lumber/v3"

func main() {
    lumber.ErrorMsg("Ahhh stuff broke")
}
```

Outputs:

![errorMsg output](images/errorMsg.png)

### [`lumber.Fatal()`](https://pkg.go.dev/github.com/gleich/lumber/v3#Fatal)

Output a fatal log with a stack trace.

Demo:

```go
package main

import (
    "os"

    "github.com/gleich/lumber/v3"
)

func main() {
    fName := "invisible-file.txt"
    _, err := os.ReadFile(fName)
    if err != nil {
        lumber.Fatal(err, "Failed to read from", fName)
    }
}
```

Outputs:

![fatal output](images/fatal.png)

### [`lumber.FatalMsg()`](https://pkg.go.dev/github.com/gleich/lumber/v3#FatalMsg)

Output a fatal message.

Demo:

```go
package main

import "github.com/gleich/lumber/v3"

func main() {
    lumber.FatalMsg("Ahhh stuff broke")
}
```

Outputs:

![fatalMsg output](images/fatalMsg.png)

## Customization

You can customize the logger that lumber uses. Below is an example of some of this customization:

```go
package main

import (
    "time"

    "github.com/gleich/lumber/v3"
)

func main() {
    lumber.SetTimezone(time.Local)
    lumber.SetTimeFormat("Mon Jan 2 15:04:05 MST 2006")
    lumber.SetFatalExitCode(0)

    lumber.Done("Calling from custom logger")
}
```

# Examples

See some examples in the [\_examples/](_examples/) folder.
