# pisdui
psd file interpreter in go

# Example

```go
    package main

    import (
        "github.com/fabulousduck/pisdui/pisdui"
    )

    func main() {

        pd := pisdui.NewPSD()
        pd.LoadFile("./test4.psd")
        pd.Parse()
    }

```
