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
