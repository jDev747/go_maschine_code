package main //TODO: Find a good name!
import (
	"bytes"
	"errors"
	"strconv"

	// "fmt"
	"log"
	"os"
)

var OPTION_AUTOCLEAR_ARG = false
func safeIndex(slice []any, index int) (any, error) {
	if index < 0 {
		log.Fatal("PANIC: SAFEINDEX PARAM < 0 <GMC>")
	}
	if len(slice) < index {
		return 0, errors.New("error: not safe indexing")
	} else {
		return slice[index], nil
	}
}
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
	itype := int(instruction[0])
	if itype < 0xA9 {
		log.Fatal("PANIC: INVALID INSTRUCTION <GMC>")
	}

	switch itype {
	case 0xA9:
		os.Exit(0)
	case 0xAA:
		optionsAny, err := safeIndex([]any{instruction}, 1)
		if err != nil {
			log.Fatal("PANIC: INVALID INSTRUCTION [OPTION <MISSING>] <GMC>")
		}
		binoptions := intToBin(optionsAny.(int))
		if binoptions[0] == '1' {
			OPTION_AUTOCLEAR_ARG = true
		} //todo: add more options
	}
}

func main() {
	//tests go here
}
