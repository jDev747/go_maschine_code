package main //TODO: Find a good name!
import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var OPTION_AUTOCLEAR_ARG = false
var STACK_ARG []any
var STACK_PERSONAL []any
var DECODER string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!?.()[]{}_&%$'\"/\\@<>|+-*~#= \n\t" // add nmore charaters
func clearScreen() {
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
func intToBin(int_ int) string {
	return strconv.FormatInt(int64(int_), 2)
}
func ReadStackArg(index int) any {
	if len(STACK_ARG) < index+1 {
		log.Fatal("PANIC: INVALID STACK [STACK ARG] <GMC> MissingParamInStack")
	}
	return STACK_ARG[index]
}
func GetInstructions(path string) [][]byte {
	bytesoriginal, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	bytesnew := bytes.Split(bytesoriginal, []byte{0xA8})
	cleaned := make([][]byte, 0, len(bytesnew))
	for _, bytearr := range bytesnew {

		if len(bytearr) > 0 {
			cleaned = append(cleaned, bytearr)
		}
	}
	return cleaned
}
func ReadInstruction(instruction []byte) {
	itype := int(instruction[0])
	if itype < 0xA9 {
		log.Fatal("PANIC: INVALID INSTRUCTION <GMC>")
	}
	switch itype {
	case 0xA9:
		fmt.Println("exit code 0")
		os.Exit(0)
	case 0xAA:
		if len(instruction) < 2 {
			log.Fatal("PANIC: INVALID INSTRUCTION [OPTION <MISSING>] <GMC>")
		}
		options := instruction[1]
		binoptions := intToBin(int(options))
		if binoptions[0] == '1' {
			OPTION_AUTOCLEAR_ARG = true
		} //todo: add more options
	case 0xAB:
		if len(instruction) < 2 {
			log.Fatal("PANIC: INVALID INSTRUCTION [PUSHSTR <MISSING> ...] <GMC> MissingStack")
		}
		if len(instruction) < 3 {
			log.Fatal("PANIC: INVALID INSTRUCTION [PUSHSTR STACK <MISSING>] <GMC> MissingString")
		}
		stackany := instruction[1]

		stack := int(stackany)
		if stack > 1 {
			log.Fatal("PANIC: INVALID INSTRUCTION [PUSHSTR <TO_BIG> ...] <GMC> InvalidStack: " + fmt.Sprint(stack))
		}
		var stringtopush strings.Builder
		for _, byteitem := range instruction[2:] {
			convint := int(byteitem)
			if convint == 0xAC {
				break
			}
			if convint > len(DECODER) - 1 {
				log.Fatal("PANIC: INVALID STRING [PUSHSTR STACK <INVALID STRING>] <GMC> InvalidChar: " + fmt.Sprint(int(byteitem)))
			}
			stringtopush.WriteString(string(DECODER[convint]))
		}
		if stack == 0 {
			STACK_ARG = append(STACK_ARG, stringtopush.String())
		} else {
			STACK_PERSONAL = append(STACK_PERSONAL, stringtopush.String())
		}
	case 0xAD:
		if len(instruction) < 2 {
			log.Fatal("PANIC: INVALID INSTRUCTION [CALL <MISSING>] <GMC> MissingFunction")
		}
		function := int(instruction[1])
		switch function {
		case 0x00:
			fmt.Println(ReadStackArg(0))
		case 0x01:
			clearScreen()
		case 0x02:
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
		default:
			log.Fatal("PANIC: INVALID INSTRUCTION [CALL <NOT_FOUND>] <GMC> InvalidFunction: "+ fmt.Sprint(function))
		}
		if OPTION_AUTOCLEAR_ARG {
			STACK_ARG = make([]any, 0, 6)
		}
	case 0xAF:
		if len(instruction) < 2 {
			log.Fatal("PANIC: INVALID INSTRUCTION [CLEARSTACK <MISSING>] <GMC> MissingStack")
		}
		stack := instruction[1]
		if stack > 1 {
			log.Fatal("PANIC: INVALID INSTRUCTION [PUSHSTR <TO_BIG> ...] <GMC> InvalidStack: " + fmt.Sprint(stack))
		}
		//Todo: Continue
	}
}

func main() {
	for _, instruction := range GetInstructions("instructions.gmc") {
		ReadInstruction(instruction)
	}
}
