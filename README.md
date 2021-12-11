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
    text := "hi"

    fmt.Println(color.Cyan(text))
    fmt.Println(color.Black(text))
}
```


## Credit

- @wxnacy
