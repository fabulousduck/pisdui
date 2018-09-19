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

See issues
