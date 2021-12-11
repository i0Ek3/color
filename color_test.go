package color

import (
    "fmt"
    "testing"
)

func TestColor(t *testing.T) {
    hi := "hi"
    text := "color"

    fmt.Print(White(hi, text))
    fmt.Println(Black(hi, text))
    fmt.Println(Cyan(hi, text))
    fmt.Println(Yellow(hi, text))
    fmt.Println(Blue(hi, text))
    fmt.Println(Green(hi, text))
    fmt.Println(Red(hi, text))
    fmt.Println(Magenta(hi, text))
}
