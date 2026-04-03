package main //TODO: Find a good name!
import (
	"bytes"
	// "fmt"
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
func ReadInstruction(intruction []byte) {
	// TODO: make this
}

func main() {
	//tests go here
}
