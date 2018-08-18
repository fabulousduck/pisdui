# pisdui
psd file interpreter in go

# Example

```go
    package main

    import (
        "github.com/fabulousduck/pisdui/pisdui"
    )

    func main() {

        psd := pisdui.NewPSD("./psd/test.psd")
        psd.Parse()
    }

```

# Todo

- [ ] make reading functions for little endian and big endian
- [ ] make complete guide for the psd file format as its kinda janky right now
