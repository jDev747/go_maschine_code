package main //TODO: Find a good name!
import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

var OPTION_AUTOCLEAR_ARG = false
/*
func safeIndex[T any | int](slice []T, index int) (T, error) {
	if index < 0 {
		log.Fatal("PANIC: SAFEINDEX PARAM < 0 <GMC>")
	}
	if len(slice) < index + 1 {
		return 0, errors.New("error: not safe indexing")
	} else {
		return slice[index], nil
	}
}
	*/
func intToBin(int_ int) string {
	return strconv.FormatInt(int64(int_), 2)
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
	fmt.Println(instruction)
	itype := int(instruction[0])
	if itype < 0xA9 {
		log.Fatal("PANIC: INVALID INSTRUCTION <GMC>")
	}
	fmt.Println(itype)
	switch itype {
	case 0xA9:
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
			log.Fatal("PANIC: INVALID INSTRUCTION [PUSHSTR <MISSING> ...] <GMC>")
		}
		stackany := instruction[1]
		
		stack := int(stackany)
		if stack > 2 {
			log.Fatal("PANIC: INVALID INSTRUCTION [PUSHSTR <TO_BIG> ...] <GMC>")
		}
		//TODO: Continue
	}
}

func main() {
	ReadInstruction(GetInstructions("instructions.gmc")[0])
}
