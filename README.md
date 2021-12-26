# color

> color.go modified from https://github.com/wxnacy/study/blob/master/goland/src/color/main.go, thanks to the author.

## Feature

Support multiple text output.


## Getting Started

### Install

`$ go get https://github.com/i0Ek3/color`


### Usage

```Go
package main

import (
    "fmt"

    "github.com/i0Ek3/color"
)

func main() {
    hi := "hi"
    color := "color"

    fmt.Println(color.Cyan(hi, color))
    fmt.Println(color.Black(hi, color))

    fmt.Println(color.Red(str...))
    // Or just use Colorize()
    fmt.Println(Colorize("blue", str...))
}
```


## Credit

- @wxnacy
