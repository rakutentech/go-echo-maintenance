[![CircleCI](https://circleci.com/gh/rakutentech/go-echo-maintenance.svg?style=svg)](https://circleci.com/gh/rakutentech/go-echo-maintenance)

# go-echo-maintenance

Middleware for echo framework to make all requests respond according to a handler function based on the presence/absence of a file specified. It is built with `go version go1.13.1 darwin/amd64`

## Installing

### *go get

```bash
    go get -u "github.com/rakutentech/go-echo-maintenance"
```

## Usage

```golang
package main

import (
    "net/http"

    maintenance "github.com/rakutentech/go-echo-maintenance"
    "github.com/labstack/echo"
)

func customHandlerFunc(c echo.Context) error {
    return c.HTML(http.StatusNotFound, `<h1 style="color:red;">Request is intercepted by maintenance middleware!</h1>`)
}

func main() {
    e := echo.New()
    middleware := maintenance.NewMaintMiddleware("/path/to/maint/file", customHandlerFunc)

    e.Use(middleware.CheckMaintenance)
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    // Start server
    e.Logger.Fatal(e.Start(":1323"))
}
```

## Create maintenance file

```bash
touch /path/to/maint/file
```

The middleware will intercept requests to api domain and respond according to the custom handler specified by user.

## Remove maintenance file

```bash
rm /path/to/maint/file
```

The middleware will no longer intercept requests to api domain.
