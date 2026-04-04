package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"

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
func FuncClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default: // stuff thatis not windows
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
	// Does not work in vscode terminal btw
}
