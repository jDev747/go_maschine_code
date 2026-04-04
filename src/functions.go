package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

func FuncColorPrint() {
	pcolor := ReadStackArg(0).(int)
	stringtoprint := ReadStackArg(1).(string) // what
	switch pcolor {
	case 0x00:
		color.Red(stringtoprint)
	case 0x01:
		color.Blue(stringtoprint)
	case 0x02:
		color.Yellow(stringtoprint)
	case 0x03:
		color.White(stringtoprint)
	case 0x04:
		color.Cyan(stringtoprint)
	default:
		log.Fatal("PANIC: INVALID STACK [CALL COLORPRINT] [STACK ARG]<GMC> InvalidColor: " + string(pcolor))
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
