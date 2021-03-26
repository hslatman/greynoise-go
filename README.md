# greynoise-go

## Usage

```bash
go get github.com/hslatman/greynoise-go
```

```go
package main

import (
    "log"
    "net"
    "github.com/hslatman/greynoise-go/client"
)

func main() {
    key := "<key>"
    c, err := client.New(key)
    if err != nil {
        log.Fatal(err)
    }

    if _, err := c.Ping(); err != nil {
        log.Fatal(err)
    }

    r, err := c.Community(net.ParseIP("127.0.0.1"))
    if err != nil {
        // NOTE: a 404 is also an error, but can actually be OK.
        log.Fatal(err)
    }

    fmt.Printf("IP: %s; Classification: %s\n", r.IP, r.Classification)
}
```