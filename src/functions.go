package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/fatih/color"
)

func CommandCall(instruction []byte) {
	if len(instruction) < 2 {
		log.Fatal("PANIC: INVALID INSTRUCTION [CALL <MISSING>] <GMC> MissingFunction")
	}
	function := int(instruction[1])
	// we can add functions without changed the compiler code.
	switch function {
	case 0x00:
		for _, i := range STACK_ARG {
			fmt.Print(i)
		}
	case 0x01:
		FuncClearScreen()
	case 0x02:
		FuncColorPrint()
	case 0x03:
		FuncOpAll("+")
	case 0x04:
		FuncOpAll("-")
	case 0x05:
		FuncOpAll("*")
	case 0x06:
		FuncOpAll("/")
	case 0x07:
		FuncReadUserInp()
	case 0x08:
		FuncReadUserInpInt()
	default:
		log.Fatal("PANIC: INVALID INSTRUCTION [CALL <NOT_FOUND>] <GMC> InvalidFunction: " + fmt.Sprint(function))
	}
	if OPTION_AUTOCLEAR_CALL {
		STACK_ARG = make([]any, 0, 6)
	}
}
func FuncOpAll(op string) {
	var total float64
	first := true

	for _, num := range STACK_ARG {
		var v float64
		switch x := num.(type) {
		case float32:
			v = float64(x)
		case float64:
			v = x
		case int:
			v = float64(x)
		default:
			log.Fatal("PANIC: INVALID STACK [CALL ADDALL] [NON-NUMERIC STACK!] <GMC> InvalidStackVal: " + fmt.Sprint(num))
		}

		if first {
			total = v
			first = false
			continue
		}

		switch op {
		case "+":
			total += v
		case "-":
			total -= v
		case "*":
			total *= v
		case "/":
			total /= v
		default:
			log.Fatal("PANIC: INVALID OPERATION [CALL ADDALL] <GMC> InvalidOp: " + op)
		}
	}

	STACK_RETURN = append(STACK_RETURN, total)
}
func FuncColorPrint() {
	pcolor := ReadStackArg(0).(int)
	stringtoprint := ReadStackArg(1).(string)
	switch pcolor {
	case 0x00:
		color.New(color.FgRed).Print(stringtoprint)
	case 0x01:
		color.New(color.FgBlue).Print(stringtoprint)
	case 0x02:
		color.New(color.FgYellow).Print(stringtoprint)
	case 0x03:
		color.New(color.FgWhite).Print(stringtoprint)
	case 0x04:
		color.New(color.FgCyan).Print(stringtoprint)
	default:
		log.Fatal("PANIC: INVALID STACK [CALL COLORPRINT] [STACK ARG] <GMC> InvalidColor: " + fmt.Sprint(pcolor))
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
func FuncReadUserInp() {
	var result string
	_, err := fmt.Scanln(&result)
	if err != nil {
		log.Fatal(err)
	}
	STACK_RETURN = append(STACK_RETURN, result)
}
func FuncReadUserInpInt() {
	var result string
	var r2 int64
	_, err := fmt.Scanln(&result)
	if v, err2 := strconv.ParseInt(result, 10, 64); err2 == nil {
		r2 = v
	} else {
		log.Fatal("PANIC: INVALID USER INPUT [CALL INPUTINT] <GMC> InvalidInt: " + result)
	}
	if err != nil {
		log.Fatal(err)
	}
	STACK_RETURN = append(STACK_RETURN, r2)
}
