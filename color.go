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
func Black(msg string) string {
    return setColor(msg, 0, 0, BlackX)
}

// Red shows the color red
func Red(msg string) string {
    return setColor(msg, 0, 0, RedX)
}

// Green shows the color green
func Green(msg string) string {
    return setColor(msg, 0, 0, GreenX)
}

// Yellow shows the color yellow
func Yellow(msg string) string {
    return setColor(msg, 0, 0, YellowX)
}

// Blue shows the color blue
func Blue(msg string) string {
    return setColor(msg, 0, 0, BlueX)
}

// Magenta shows the color magenta
func Magenta(msg string) string {
    return setColor(msg, 0, 0, MagentaX)
}

// Cyan shows the color cyan
func Cyan(msg string) string {
    return setColor(msg, 0, 0, CyanX)
}

// White shows the color white
func White(msg string) string {
    return setColor(msg, 0, 0, WhiteX)
}

// setColor sets color for msg
func setColor(msg string, conf, bg, text int) string {
    return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}
