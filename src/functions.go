package main

import (
	"log"
	"github.com/fatih/color"
)

func FuncColorPrint() {
	pcolor := ReadStackArg(0).(string)
	stringtoprint := ReadStackArg(1).(string) // what
	switch pcolor {
	case "red":
		color.Red(stringtoprint)
	case "blue":
		color.Blue(stringtoprint)
	case "yellow":
		color.Yellow(stringtoprint)
	case "white":
		color.White(stringtoprint)
	case "cyan":
		color.Cyan(stringtoprint)
	default:
		log.Fatal("PANIC: INVALID STACK [CALL COLORPRINT] [STACK ARG]<GMC> InvalidColor: " + pcolor)
	}

}
