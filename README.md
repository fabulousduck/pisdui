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

# About the errors
Since we do a metric butt load of individual reads to the file to get the data out, I have chosen not to error check them all since this would make the code extremely
painfull to work with and a mess. Instead. at the end of each major part we check for the size already read, and then see if its still possible to read the next part
of the file given its size. this way we can avoid doing the EOF checking.