// package color shows the color for given message
package color

import (
    "fmt"
)

const (
	BlackX = iota + 30
	RedX
	GreenX
	YellowX
	BlueX
	MagentaX
	CyanX
	WhiteX
)

// Black shows the color black
func Black(msg ...string) string {
    return setColor(0, 0, BlackX, msg...)
}

// Red shows the color red
func Red(msg ...string) string {
    return setColor(0, 0, RedX, msg...)
}

// Green shows the color green
func Green(msg ...string) string {
    return setColor(0, 0, GreenX, msg...)
}

// Yellow shows the color yellow
func Yellow(msg ...string) string {
    return setColor(0, 0, YellowX, msg...)
}

// Blue shows the color blue
func Blue(msg ...string) string {
    return setColor(0, 0, BlueX, msg...)
}

// Magenta shows the color magenta
func Magenta(msg ...string) string {
    return setColor(0, 0, MagentaX, msg...)
}

// Cyan shows the color cyan
func Cyan(msg ...string) string {
    return setColor(0, 0, CyanX, msg...)
}

// White shows the color white
func White(msg ...string) string {
    return setColor(0, 0, WhiteX, msg...)
}

// setColor sets color for msg
func setColor(conf, bg, text int, msg ...string) string {
    res := ""
    for _, v := range msg {
        res += v + " "
    }
    return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, res, 0x1B)
}

// Colorize made for debug use
func Colorize(colorName string, str ...string) {
    switch colorName {
    case "white":
        fmt.Println(White(str...))
    case "black":
        fmt.Println(Black(str...))
    case "green":
        fmt.Println(Green(str...))
    case "yellow":
        fmt.Println(Yellow(str...))
    case "red":
        fmt.Println(Red(str...))
    case "cyan":
        fmt.Println(Cyan(str...))
    case "magenta":
        fmt.Println(Magenta(str...))
    }
}
