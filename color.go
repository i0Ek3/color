// package color shows the color for given message
package color

import (
    "fmt"
)

const (
	Black = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Black shows the color black
func Black(msg string) string {
    return setColor(msg, 0, 0, Black)
}

// Red shows the color red
func Red(msg string) string {
    return setColor(msg, 0, 0, Red)
}

// Green shows the color green
func Green(msg string) string {
    return setColor(msg, 0, 0, Green)
}

// Yellow shows the color yellow
func Yellow(msg string) string {
    return setColor(msg, 0, 0, Yellow)
}

// Blue shows the color blue
func Blue(msg string) string {
    return setColor(msg, 0, 0, Blue)
}

// Magenta shows the color magenta
func Magenta(msg string) string {
    return setColor(msg, 0, 0, Magenta)
}

// Cyan shows the color cyan
func Cyan(msg string) string {
    return setColor(msg, 0, 0, Cyan)
}

// White shows the color white
func White(msg string) string {
    return setColor(msg, 0, 0, White)
}

// setColor sets color for msg
func setColor(msg string, conf, bg, text int) string {
    return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}
