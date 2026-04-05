package main //TODO: Find a good name!
import (
	"bytes"
	"fmt"
	"log"
	"os"
)

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
		log.Fatal("PANIC: INVALID INSTRUCTION <GMC> InvalidInstruction: " + fmt.Sprint(itype))
	}
	switch itype {
	case 0xA9:
		fmt.Println("exit code 0")
		os.Exit(0)
	case 0xAA:
		if len(instruction) < 2 {
			log.Fatal("PANIC: INVALID INSTRUCTION [OPTION <MISSING>] <GMC> Missing Options")
		}
		options := instruction[1]
		binoptions := IntToBin(int(options))
		if binoptions[0] == '1' {
			OPTION_AUTOCLEAR_ARG = true
		} //todo: add more options
	case 0xAB:
		CommandPushStr(instruction)
	case 0xAD:
		CommandCall(instruction)
	case 0xAF:
		CommandClearStack(instruction)
	case 0xB0:
		CommandPushInt(instruction)
	case 0xB1:
		CommandStrVar(instruction)
	case 0xB2:
		CommandPushVar(instruction)
	case 0xB3:
		CommandIntVar(instruction)
	case 0xB4:
		CommandReadStack(instruction)
	default:
		log.Fatal("PANIC: INVALID INSTRUCTION <GMC> InvalidInstruction: " + fmt.Sprint(itype))
	}
}

func main() {
	for _, instruction := range GetInstructions("../test/smth.gmc") {
		ReadInstruction(instruction)
	}
}
