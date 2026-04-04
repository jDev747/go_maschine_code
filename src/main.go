package main //TODO: Find a good name!
import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)


var DECODER string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!?.()[]{}_&%$'\"/\\@<>|+-*~#= \n\t" // add nmore charaters

func intToBin(int_ int) string {
	return strconv.FormatInt(int64(int_), 2)
}
func BinToInt(bin string) int64 {
	val, err := strconv.ParseInt(bin, 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	return val
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
		CommandPushStr(instruction)
	case 0xAD:
		if len(instruction) < 2 {
			log.Fatal("PANIC: INVALID INSTRUCTION [CALL <MISSING>] <GMC> MissingFunction")
		}
		function := int(instruction[1])
		switch function {
		case 0x00:
			fmt.Println(ReadStackArg(0))
		case 0x01:
			FuncClearScreen()
		case 0x02:
			FuncColorPrint()
		
		default:
			log.Fatal("PANIC: INVALID INSTRUCTION [CALL <NOT_FOUND>] <GMC> InvalidFunction: " + fmt.Sprint(function))
			if OPTION_AUTOCLEAR_ARG {
				STACK_ARG = make([]any, 0, 6)
			}
		}
	case 0xAF:
		CommandClearStack(instruction)
	case 0xB0:
		CommandPushInt(instruction)
	}
}

func main() {
	for _, instruction := range GetInstructions("../instructions.gmc") {
		ReadInstruction(instruction)
	}
}
