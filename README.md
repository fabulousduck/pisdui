# pisdui
psd file interpreter in go

# Example

```go
    package main

    import (
        "github.com/fabulousduck/pisdui/pisdui"
    )

    func main() {

        psd, err := pisdui.NewPSD("./psd/test.psd")
        if err != nil {
            panic(err)
        }
         psd.Parse()
    }

```

# Todo

See issues
IDK
