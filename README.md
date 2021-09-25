<!-- DO NOT REMOVE - contributor_list:data:start:["gleich", "ImgBotApp"]:end -->

<div align="center">
    <h1>🪵 lumber 🪵</h1>
    <a href="https://pkg.go.dev/github.com/gleich/lumber/v2"><img alt="Godoc Reference" src="https://godoc.org/github.com/gleich/lumber?status.svg"></a>
    <img alt="test workflow result" src="https://github.com/gleich/lumber/workflows/test/badge.svg">
    <img alt="lint workflow result" src="https://github.com/gleich/lumber/workflows/lint/badge.svg">
    <br/>
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/gleich/lumber">
    <img alt="Golang report card" src ="https://goreportcard.com/badge/github.com/gleich/lumber/v2">
    <br/>
    <br/>
    <i>A dead simple, pretty, and feature-rich logger for golang</i>
</div>
<hr>

- [🚀 Install](#-install)
- [🌲 Logging Functions](#-logging-functions)
  - [`lumber.Success()`](#lumbersuccess)
  - [`lumber.Info()`](#lumberinfo)
  - [`lumber.Debug()`](#lumberdebug)
  - [`lumber.Warning()`](#lumberwarning)
  - [`lumber.Error()`](#lumbererror)
  - [`lumber.ErrorMsg()`](#lumbererrormsg)
  - [`lumber.Fatal()`](#lumberfatal)
  - [`lumber.FatalMsg()`](#lumberfatalmsg)
- [⚙️ Customization](#️-customization)
- [✨ Examples](#-examples)
- [🙌 Contributing](#-contributing)
- [👥 Contributors](#-contributors)

## 🚀 Install

Simply run the following from your project root:

```bash
go get -u github.com/gleich/lumber/v2
```

## 🌲 Logging Functions

### [`lumber.Success()`](https://pkg.go.dev/github.com/gleich/lumber/v2#Success)

Output a success log.

Demo:

```go
package main

import (
    "time"

    "github.com/gleich/lumber/v2"
)

func main() {
    lumber.Success("Loaded up the program!")
    time.Sleep(2 * time.Second)
    lumber.Success("Waited 2 seconds!")
}
```

Outputs:

![success output](images/success.png)

### [`lumber.Info()`](https://pkg.go.dev/github.com/gleich/lumber/v2#Info)

Output an info log.

Demo:

```go
package main

import (
    "time"

    "github.com/gleich/lumber/v2"
)

func main() {
    lumber.Info("Getting the current year")
    now := time.Now()
    lumber.Info("Current year is", now.Year())
}
```

Outputs:

![info output](images/info.png)

### [`lumber.Debug()`](https://pkg.go.dev/github.com/gleich/lumber/v2#Debug)

Output a debug log.

Demo:

```go
package main

import (
    "os"

    "github.com/gleich/lumber/v2"
)

func main() {
    homeDir, _ := os.UserHomeDir()
    lumber.Debug("User's home dir is", homeDir)
}
```

Outputs:

![debug output](images/debug.png)

### [`lumber.Warning()`](https://pkg.go.dev/github.com/gleich/lumber/v2#Warning)

Output a warning log.

Demo:

```go
package main

import (
    "time"

    "github.com/gleich/lumber/v2"
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

### [`lumber.Error()`](https://pkg.go.dev/github.com/gleich/lumber/v2#Error)

Output an error log with a stack trace.

Demo:

```go
package main

import (
    "os"

    "github.com/gleich/lumber/v2"
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

### [`lumber.ErrorMsg()`](https://pkg.go.dev/github.com/gleich/lumber/v2#ErrorMsg)

Output an error message.

Demo:

```go
package main

import "github.com/gleich/lumber/v2"

func main() {
    lumber.ErrorMsg("Ahhh stuff broke")
}
```

Outputs:

![errorMsg output](images/errorMsg.png)

### [`lumber.Fatal()`](https://pkg.go.dev/github.com/gleich/lumber/v2#Fatal)

Output a fatal log with a stack trace.

Demo:

```go
package main

import (
    "os"

    "github.com/gleich/lumber/v2"
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

### [`lumber.FatalMsg()`](https://pkg.go.dev/github.com/gleich/lumber/v2#FatalMsg)

Output a fatal message.

Demo:

```go
package main

import "github.com/gleich/lumber/v2"

func main() {
    lumber.FatalMsg("Ahhh stuff broke")
}
```

Outputs:

![fatalMsg output](images/fatalMsg.png)

## ⚙️ Customization

You can customize lumber by creating a custom logger and changing values on it. You then call the log functions on the custom logger. Below is an example of this.

```go
package main

import "github.com/gleich/lumber/v2"

func main() {
    log := lumber.NewCustomLogger()
    log.ColoredOutput = false
    log.ExitCode = 2

    log.Success("Calling from custom logger!")
}
```

Here are all the variables that can be changed:

| **Variable Name** | **Description**                                                                      | **Default Value**      | **Type**         |
| ----------------- | ------------------------------------------------------------------------------------ | ---------------------- | ---------------- |
| `NormalOut`       | The output file for Debug, Success, Warning, and Info                                | `os.Stdout`            | `*os.File`       |
| `ErrOut`          | The output file for Fatal and Error                                                  | `os.Stderr`            | `*os.File`       |
| `ExitCode`        | Fatal exit code                                                                      | `1`                    | `int`            |
| `Padding`         | If the log should have an extra new line at the bottom                               | `false`                | `bool`           |
| `ColoredOutput`   | If the output should have color                                                      | `true`                 | `bool`           |
| `TrueColor`       | If the output colors should be true colors. Default is true if terminal supports it. | `has256ColorSupport()` | `bool`           |
| `ShowStack`       | If stack traces should be shown                                                      | `true`                 | `bool`           |
| `Multiline`       | If the should should be spread out to more than one line                             | `false`                | `bool`           |
| `Timezone`        | Timezone you want the times to be logged in                                          | `time.UTC`             | `*time.Location` |

# ✨ Examples

See some examples in the [examples/](examples/) folder! Just run them using `go run main.go`.

## 🙌 Contributing

Before contributing please read the [CONTRIBUTING.md file](https://github.com/gleich/lumber/v2/v2blob/master/CONTRIBUTING.md).

<!-- DO NOT REMOVE - contributor_list:start -->

## 👥 Contributors

- **[@gleich](https://github.com/gleich)**

- **[@ImgBotApp](https://github.com/ImgBotApp)**

<!-- DO NOT REMOVE - contributor_list:end -->
