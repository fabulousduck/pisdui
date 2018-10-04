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

# contributing

At the moment I only have so many psd files to work with as I do not own a copy of Photoshop

If you have a .psd with resourceblocks that pisdui does not parse yet, please post an issue and attach a
link to the .psd so I can implement it.
